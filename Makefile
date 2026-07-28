js/main.wasm: wasm/wasm.go internal/*.go heapfast/*.go
	@tinygo build --target=wasm -o js/main.wasm ./wasm 

.PHONY: testheap
testheap: go.* heapfast/*.go
	@go test github.com/raulpy271/heapfast/heapfast

.PHONY: testinternal
testinternal: testheap internal/*.go
	@go test github.com/raulpy271/heapfast/internal

.PHONY: testjs
testjs: js/main.wasm js/*.js js/test/*.js js/package*.json
	@cd js && npm test

.PHONY: test
test: testheap testinternal testjs

.PHONY: fmt
fmt: wasm/wasm.go internal/*.go heapfast/*.go
	@go fmt github.com/raulpy271/heapfast/heapfast
	@go fmt github.com/raulpy271/heapfast/internal
	@go fmt github.com/raulpy271/heapfast/wasm

.PHONY: build
build: js/main.wasm

