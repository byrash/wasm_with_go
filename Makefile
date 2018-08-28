.PHONY: build wasm
.DEFAULT_GOAL := serve

GOROOT = $(shell go env GOROOT)

serve: build
	go run server.go

build: build_wasm

clear:
	rm -rf ./build/*
	# rm -rf ./build/wasm_exec.{html,js}

copy_wasm_files:
	cp $(GOROOT)/misc/wasm/wasm_exec.{html,js} ./build

copy-static-stuff:
	cp index.html ./build
	cp wasm_exec.js ./build

build_wasm: clear copy-static-stuff
	GOARCH=wasm GOOS=js go build -o ./build/calculator.wasm calculator.go