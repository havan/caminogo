#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Avalanchego root folder
CAMINOGO_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the constants
source "$CAMINOGO_PATH"/scripts/constants.sh

echo "Building testnetctl..."
go build -ldflags\
   "-X github.com/ava-labs/avalanchego/version.GitCommit=$git_commit $static_ld_flags"\
   -o "$CAMINOGO_PATH/build/testnetctl"\
   "$CAMINOGO_PATH/tests/fixture/testnet/cmd/"*.go
