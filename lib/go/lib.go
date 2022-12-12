//go:build js && wasm

// NOTE THIS WILL NOT ACTUALLY WORK WITH WASMER DUE TO ITS CURRENT SUPPORT.
// A SHIM EXISTS TO GET IT WORKING BUT NOT NEEDED FOR THIS EXAMPLE
package main

import (
	"fmt"
	"syscall/js"
)

func add(x int, y int) int {
	return x + y
}

func addWrapper() js.Func {
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			return nil
		}

		return add(args[0].Int(), args[1].Int())
	})
	return jsFunc
}

func hello() string {
	return fmt.Sprintf("hello world from -> GO")
}

func main() {
	js.Global().Set("add", addWrapper())
	js.Global().Set("hello", hello())
	<-make(chan bool) // prevent closing the webassembly
}
