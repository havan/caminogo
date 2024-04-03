#!/usr/bin/env bash

set -euo pipefail

# Camino-Node root folder
CAMINOGO_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )

echo "Downloading dependencies..."
(cd "$CAMINOGO_PATH" && go mod download)

# Build caminogo
"$CAMINOGO_PATH"/scripts/build_camino.sh