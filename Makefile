.PHONY: package-assets

assets:
	mkdir -p ./assets

assets/wasm_exec.js:assets
	mkdir -p ./assets/javascript
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./assets/javascript/wasm_exec.js

assets_vfsdata.go:assets/javascript/wasm_exec.js
	go run assets-generator.go

bin:
	mkdir -p ./bin

bin/lazy-wasm-server:assets_vfsdata.go bin
	go build -o bin ./lazy-wasm-server

clean:
	git clean -xdf
