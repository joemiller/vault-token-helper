#!/bin/bash

set -eou pipefail

CODESIGN_CERT="Developer ID Application: JOSEPH MILLER (P3MF48HUD7)"

path="$1"

# sign
codesign -s "$CODESIGN_CERT" -i "vault-token-helper" "$path"

# display signature
codesign -v -d "$path"
