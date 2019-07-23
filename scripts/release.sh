#!/bin/bash
# this script is intended to be run inside the dockercore/golang-cross docker image, eg:
#
# docker run \
# 	--rm \
#   -e "GITHUB_TOKEN=$GITHUB_TOKEN" \
#   -e "GPG_KEY=$GPG_KEY" \
#   -v `pwd`:/src \
#   -w /src \
#   dockercore/golang-cross \
#     /src/release.sh
#
# (optional) arguments will be passed to goreleaser, eg:
#
#    /src/release.sh --snapshot --rm-dist
#
# (optional) sign releases with $GPG_KEY. The key should be base64 encoded.

set -eou pipefail

GORELEASER_ARGS=("$@")

if [[ -n "${GPG_KEY:-}" ]]; then
  GNUPGHOME="$HOME/releaser-gpg"
  export GNUPGHOME
  mkdir -p "$GNUPGHOME"
  chmod 0700 "$GNUPGHOME"

  echo "$GPG_KEY" \
    | base64 --decode --ignore-garbage \
    | gpg --batch --allow-secret-key-import --import

  gpg --keyid-format LONG --list-secret-keys

  trap 'rm -rf -- "$GNUPGHOME"' EXIT
else
  echo "==> WARNING: Missing GPG_KEY env var, skipping GPG signing of the release"
  GORELEASER_ARGS+=("--skip-sign")
fi

apt-get -qy update
apt-get -qy install rpm

curl -sL https://git.io/goreleaser | bash -s -- "${GORELEASER_ARGS[@]}"
