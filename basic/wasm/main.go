//go:build js && wasm
// +build js,wasm

package main

import (
	"math"
	"syscall/js"
)

func hi(_ js.Value, args []js.Value) any {
	name := args[0].String()
	return "Hello, " + name
}

func getCurrentTime(_ js.Value, _ []js.Value) any {
	return js.Global().Get("Date").New().Call("toISOString")
}

func pythagorean(_ js.Value, args []js.Value) any {
	a, b := args[0].Float(), args[1].Float()
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	js.Global().Set("greet", js.FuncOf(hi))
	js.Global().Set("getCurrentTime", js.FuncOf(getCurrentTime))
	js.Global().Set("pythagorean", js.FuncOf(pythagorean))
	select {}
}
