#!/usr/bin/env bash

set -euo pipefail

if ! [[ "$0" =~ scripts/mock.gen.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

if ! command -v mockgen &> /dev/null
then
  echo "mockgen not found, installing..."
  # https://github.com/uber-go/mock
  go install -v go.uber.org/mock/mockgen@v0.2.0
fi

# tuples of (source interface import path, comma-separated interface names, output file path)
input="scripts/mocks.mockgen.txt"
while IFS= read -r line
do
  IFS='=' read src_import_path interface_name output_path <<< "${line}"
  package_name=$(basename $(dirname $output_path))
  echo "Generating ${output_path}..."
  mockgen -copyright_file=./LICENSE.header -package=${package_name} -destination=${output_path} ${src_import_path} ${interface_name}
done < "$input"

echo "SUCCESS"
