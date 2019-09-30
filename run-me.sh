#!/usr/bin/env bash
set +e
set -x
if [ -f assets_vfsdata.go ]; then
  rm assets_vfsdata.go
fi
# make assets_vfsdata.go
make bin/lazy-wasm-server
./bin/lazy-wasm-server
