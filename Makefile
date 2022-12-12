build-all: build-hello-wasm build-user-wasm build-go-lib-wasm build-rust-lib-wasm build-server-linux build-server-macos build-server-windows

build-server-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/server-linux wasm_experiments/cmd/server

build-server-macos:
	GOOS=darwin GOARCH=arm64 go build -o bin/server-macos wasm_experiments/cmd/server

build-server-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/server-windows wasm_experiments/cmd/server

build-hello-wasm:
	GOOS=js GOARCH=wasm go build -o bin/hello_world.wasm wasm_experiments/cmd/hello_world

build-user-wasm:
	GOOS=js GOARCH=wasm go build -o bin/hello_user.wasm wasm_experiments/cmd/hello_user

build-go-lib-wasm:
	GOOS=js GOARCH=wasm go build -o bin/hello_go.wasm wasm_experiments/lib/go

build-rust-lib-wasm:
	wasm-pack build lib/rust --target web
	cp lib/rust/pkg/rust_hello_wasm.js bin/hello_rust_wasm.js
	cp lib/rust/pkg/rust_hello_wasm_bg.wasm bin/hello_rust.wasm

