//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"syscall/js"
)

func sum(_ js.Value, args []js.Value) any {
	data := args[0]
	slice := make([]int, data.Length())

	for index := 0; index < data.Length(); index++ {
		slice[index] = data.Index(index).Int()
	}

	var total int
	for _, value := range slice {
		total += value
	}

	return total
}

func processMap(_ js.Value, args []js.Value) any {
	jsonData := args[0].String()
	var data map[string]any
	json.Unmarshal([]byte(jsonData), &data)

	data["success"] = true

	processed, _ := json.Marshal(data)
	return string(processed)
}

func main() {
	js.Global().Set("sum", js.FuncOf(sum))
	js.Global().Set("processMap", js.FuncOf(processMap))
	select {}
}
