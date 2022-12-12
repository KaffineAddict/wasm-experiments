//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
)

func helloUser(n string) string {
	return fmt.Sprintf("Hello %s, this is the JS->GO interop", n)
}

func helloUserWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return fmt.Sprintf("Hello User expects exactly 1 argument. %d arguments recieved", len(args))
		}

		return helloUser(args[0].String())
	})
	return jsonFunc
}

func main() {
	fmt.Println("GO hello user assembly loaded")
	js.Global().Set("helloUserInterop", helloUserWrapper())
	<-make(chan bool) // prevent closing the webassembly
}
