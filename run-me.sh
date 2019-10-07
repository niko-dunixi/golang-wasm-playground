#!/usr/bin/env bash
set +e
set -x

for file_to_rm in "${@}"; do
  if [ -f "${file_to_rm}" ]; then
    rm -v "${file_to_rm}"
  fi
done

if [ -f assets_vfsdata.go ]; then
  rm assets_vfsdata.go
fi
# make assets_vfsdata.go
make bin/lazy-wasm-server
./bin/lazy-wasm-server
