#!/bin/bash
# inputs:
# - $TAG, eg: v0.1.1
#
# outcome:
# - download all existing artifacts from the release
# - codesign the macos binary
# - create a shasum file of all assets
# - sign the shasum file with GPG (created a detached sig file)
# - re-upload the codesign'd macos binary
# - upload the shasum file
# - upload the shasum signature file
# - promote release from draft to published
#
# requires: gothub - https://github.com/itchio/gothub

set -eou pipefail
shopt -s nullglob

TAG="${TAG:-}"
ORG="joemiller"
REPO="vault-token-helper"
BINARY="vault-token-helper"
CODESIGN_CERT="Developer ID Application: JOSEPH MILLER (P3MF48HUD7)"
GPG_KEY="6720A9FD78AC13F5"

if [[ -z "$TAG" ]]; then
  echo "Missing env var 'TAG'"
  exit 1
fi

release_info_json=''
assets=()
modified_assets=()
checksum_file=''
sig_file=''
description=''
tempdir="$(mktemp -d)"

echo "==> Created tempdir: $tempdir"
trap 'echo "Cleaning up."; rm -rf -- "$tempdir"' EXIT

echo
echo "==> Fetching existing release info for $TAG"
release_info_json=$(gothub info -t "$TAG" -u "$ORG" -r "$REPO" -j)

echo
echo "==> Generating a list of assets"
for i in $(jq -r '.Releases[0].assets[] | .name' <<<"$release_info_json"); do
  assets+=("$i")
  echo "$i"
done
echo "==> Found: ${#assets[@]} assets"

echo
echo "==> Downloading assets to: $tempdir"
pushd "$tempdir" >/dev/null
for i in "${assets[@]}"; do
  echo "==> Downloading: $i"
  gothub download -t "$TAG" -u "$ORG" -r "$REPO" -n "$i"
done
ls -l "$tempdir"

echo
echo "==> Apple codesigning the macOS binaries"
for i in ./*_darwin_amd64*; do
  modified_assets+=("$i")

  if [[ "$i" =~ (.tar|.zip) ]]; then
    echo "==> untarring and codesigning archived macOS binary: $i"
    tartmp="./tar-tmp"
    mkdir "$tartmp"
    tar -xzf "$i" -C "$tartmp"
    codesign -s "$CODESIGN_CERT" -i "$BINARY" "$tartmp/$BINARY"
    tar -czf "$i" -C "$tartmp" $(ls "$tartmp")
    rm -rf -- "$tartmp"
  else
    echo "==> codesigning binary: $i"
    codesign -s "$CODESIGN_CERT" -i "$BINARY" "$i"
  fi
done

echo
echo "==> Generating new checksum file"
# delete existing checksum file before gathering new checksums
checksum_file="${BINARY}_$(sed -e 's/^v//' <<<"$TAG")_checksums.txt"
rm -f -- "$checksum_file"
shasum -a 256 -- * >"$checksum_file"
cat "$checksum_file"
modified_assets+=("$checksum_file")

echo
echo "==> GPG-singing checksum file"
sig_file="${checksum_file}.sig"
rm -f -- "$sig_file"
gpg --batch -u "$GPG_KEY" --output "$sig_file" --detach-sign "$checksum_file"
modified_assets+=("$sig_file")

echo
echo "==> Re-uploading modified assets"
#for i in ./*; do
for i in "${modified_assets[@]}"; do
  echo "==> Uploading: $i"
  gothub upload -t "$TAG" -u "$ORG" -r "$REPO" -n "$(basename "$i")" -f "$i" --replace
done

echo
echo "==> Promoting release from draft to published"
# in order to preserve the current description we must provide it to the edit command:
description="$(jq -r '.Releases[0].body' <<<"$release_info_json")"
gothub edit -t "$TAG" -u "$ORG" -r "$REPO" -d "$description"

echo
echo "DONE!"
echo "Next steps:"
echo "- Download the macos tarball from: https://github.com/joemiller/vault-token-helper/releases/latest"
echo "- update sha256 sum in the homebrew formula: https://github.com/joemiller/homebrew-taps/blob/master/Formula/vault-token-helper.rb"
