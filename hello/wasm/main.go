//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

func hi(_ js.Value, args []js.Value) any {
	name := args[0].String()
	return "Hello, " + name
}

func currentTime(_ js.Value, _ []js.Value) any {
	return js.Global().Get("Date").New().Call("toISOString")
}

func main() {
	js.Global().Set("greet", js.FuncOf(hi))
	js.Global().Set("now", js.FuncOf(currentTime))

	select {}
}
