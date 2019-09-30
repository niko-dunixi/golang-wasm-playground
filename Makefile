.PHONY: package-assets assets/wasm/all
.DEFAULT: bin/lazy-wasm-server

assets/javascript/wasm_exec.js:
	mkdir -p ./assets/javascript
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./assets/javascript/wasm_exec.js

assets/wasm/simple-cat-example.wasm:
	mkdir -p ./assets/wasm/
	GOOS=js GOARCH=wasm go build -o assets/wasm/simple-cat-example.wasm ./simple-cat-example

assets/wasm/all:assets/wasm/simple-cat-example.wasm
	echo "Built all WASM files"

assets_vfsdata.go:assets/javascript/wasm_exec.js assets/wasm/all
	go run assets-generator.go

bin:
	mkdir -p ./bin

bin/lazy-wasm-server:assets_vfsdata.go bin
	go build -o bin ./lazy-wasm-server

clean:
	git clean -xdf
