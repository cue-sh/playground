// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"
)

// TODO: for some reason recompiling the main.wasm file does not trigger
// webpack to hot reload. We get "nothing hot updated"

func main() {
	api := js.Global().Get("WasmAPI")
	api.Set("CUECompile", js.FuncOf(cueCompile))
	api.Call("FireOnChange")
	select {}
}

func cueCompile(this js.Value, args []js.Value) interface{} {
	// args[0] is the input type
	// args[1] is the function
	// args[2] is the output type
	// args[3] is the actual input value
	const expArgs = 4
	if len(args) != expArgs {
		panic(fmt.Errorf("cueCompile: expected %v args, got %v", expArgs, len(args)))
	}
	for i := 0; i < expArgs; i++ {
		if t := args[i].Type(); t != js.TypeString {
			panic(fmt.Errorf("cueCompile: expected arg %v to be of type syscall/js.TypeString, got %v", i, t))
		}
	}
	in := input(args[0].String())
	fn := function(args[1].String())
	out := output(args[2].String())
	inVal := args[3].String()

	val, err := handleCUECompile(in, fn, out, inVal)
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	return map[string]interface{}{
		"value": val,
		"error": errStr,
	}
}
