package main

import (
	_ "embed"
	"github.com/wasmerio/wasmer-go/wasmer"

	"fmt"
)

//go:embed assets/hello_rust.wasm
var rustLib []byte

func main() {
	// Create an Engine
	engine := wasmer.NewEngine()
	// Create a Store
	store := wasmer.NewStore(engine)

	libraryCalls(store, rustLib)
}

func libraryCalls(store *wasmer.Store, wasm []byte) {
	module, err := wasmer.NewModule(store, wasm)
	if err != nil {
		fmt.Println("Failed to compile module:", err)
	}

	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		panic(fmt.Sprintln("Failed to instantiate the module:", err))
	}

	hello, err := instance.Exports.GetFunction("hello")
	if err != nil {
		panic(fmt.Sprintln("Failed to get the `hello` function:", err))
	}
	fmt.Println(hello())

	add, err := instance.Exports.GetFunction("add")
	if err != nil {
		panic(fmt.Sprintln("Failed to get the `add` function:", err))
	}

	result, err := add(25, 12)
	if err != nil {
		panic(fmt.Sprintln("Failed to call the `add` function:", err))
	}
	fmt.Println("Results of `add`:", result)
}
