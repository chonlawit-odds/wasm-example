//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"
)

var counter int

func updateDisplay() {
	document := js.Global().Get("document")
	display := document.Call("getElementById", "counter-display")
	display.Set("innerText", fmt.Sprintf("Counter: %d", counter))
}

func increment(_ js.Value, _ []js.Value) any {
	counter++
	updateDisplay()
	return nil
}

func decrement(_ js.Value, _ []js.Value) any {
	counter--
	updateDisplay()
	return nil
}

func main() {
	// Get document
	document := js.Global().Get("document")

	// Create increment button
	incrementButton := document.Call("createElement", "button")
	incrementButton.Set("innerText", "Increment")
	incrementButton.Call("addEventListener", "click", js.FuncOf(increment))

	// Create decrement button
	decrementButton := document.Call("createElement", "button")
	decrementButton.Set("innerText", "Decrement")
	decrementButton.Call("addEventListener", "click", js.FuncOf(decrement))

	// Create display element
	display := document.Call("createElement", "div")
	display.Set("id", "counter-display")
	display.Set("innerText", "Counter: 0")

	// Append body
	body := document.Call("querySelector", "body")
	body.Call("appendChild", incrementButton)
	body.Call("appendChild", decrementButton)
	body.Call("appendChild", display)

	select {}
}
