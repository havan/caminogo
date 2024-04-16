#!/usr/bin/env bash

set -euo pipefail

# e.g.,
# ./scripts/tests.upgrade.sh                                                # Use default version
# ./scripts/tests.upgrade.sh 1.1.18                                         # Specify a version
# CAMINOGO_BIN_PATH=./path/to/caminogo ./scripts/tests.upgrade.sh 1.1.18    # Customization of caminogo path
if ! [[ "$0" =~ scripts/tests.upgrade.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

# The CaminoGo local network does not support long-lived
# backwards-compatible networks. When a breaking change is made to the
# local network, this flag must be updated to the last compatible
# version with the latest code.
#
# v1.1.18 includes restrictions on ports sent over the p2p network along with
# proposervm and P-chain rule changes on the local network.
DEFAULT_VERSION="1.1.18"

VERSION="${1:-${DEFAULT_VERSION}}"
if [[ -z "${VERSION}" ]]; then
  echo "Missing version argument!"
  echo "Usage: ${0} [VERSION]" >>/dev/stderr
  exit 255
fi

CAMINOGO_BIN_PATH="$(realpath "${CAMINOGO_BIN_PATH:-./build/caminogo}")"

#################################
# download caminogo
# https://github.com/ava-labs/caminogo/releases
GOARCH=$(go env GOARCH)
GOOS=$(go env GOOS)
DOWNLOAD_URL=https://github.com/ava-labs/caminogo/releases/download/v${VERSION}/caminogo-linux-${GOARCH}-v${VERSION}.tar.gz
DOWNLOAD_PATH=/tmp/caminogo.tar.gz
if [[ ${GOOS} == "darwin" ]]; then
  DOWNLOAD_URL=https://github.com/ava-labs/caminogo/releases/download/v${VERSION}/caminogo-macos-v${VERSION}.zip
  DOWNLOAD_PATH=/tmp/caminogo.zip
fi

rm -f ${DOWNLOAD_PATH}
rm -rf "/tmp/caminogo-v${VERSION}"
rm -rf /tmp/caminogo-build

echo "downloading caminogo ${VERSION} at ${DOWNLOAD_URL}"
curl -L "${DOWNLOAD_URL}" -o "${DOWNLOAD_PATH}"

echo "extracting downloaded caminogo"
if [[ ${GOOS} == "linux" ]]; then
  tar xzvf ${DOWNLOAD_PATH} -C /tmp
elif [[ ${GOOS} == "darwin" ]]; then
  unzip ${DOWNLOAD_PATH} -d /tmp/caminogo-build
  mv /tmp/caminogo-build/build "/tmp/caminogo-v${VERSION}"
fi
find "/tmp/caminogo-v${VERSION}"

# Sourcing constants.sh ensures that the necessary CGO flags are set to
# build the portable version of BLST. Without this, ginkgo may fail to
# build the test binary if run on a host (e.g. github worker) that lacks
# the instructions to build non-portable BLST.
source ./scripts/constants.sh

#################################
echo "building upgrade.test"
# to install the ginkgo binary (required for test build and run)
go install -v github.com/onsi/ginkgo/v2/ginkgo@v2.13.1
ACK_GINKGO_RC=true ginkgo build ./tests/upgrade
./tests/upgrade/upgrade.test --help

#################################
# By default, it runs all upgrade test cases!
echo "running upgrade tests against the local cluster with ${CAMINOGO_BIN_PATH}"
./tests/upgrade/upgrade.test \
  --ginkgo.v \
  --avalanchego-path="/tmp/caminogo-v${VERSION}/caminogo" \
  --avalanchego-path-to-upgrade-to="${CAMINOGO_BIN_PATH}"
