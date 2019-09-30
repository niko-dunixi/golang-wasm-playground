.PHONY: package-assets

# This is only needed if new assets are
# created or existing ones modified
package-assets:
	go run assets-generator.go

bin:
	mkdir -p ./bin

bin/lazy-wasm-server:bin
	go build -o bin ./lazy-wasm-server

clean:
	git clean -xdf
