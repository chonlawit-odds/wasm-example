//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

func fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func goFibonacci(_ js.Value, args []js.Value) any {
	round := args[0].Int()

	return fibonacci(uint(round))
}

// Fibonacci algorithm
func routineFibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}
}

func goRoutineFibonacci(_ js.Value, args []js.Value) any {
	c := make(chan int, 1)
	quit := make(chan int)

	var (
		result int
		round  = args[0].Int()
	)

	go func() {
		for i := 0; i < round; i++ {
			result = <-c
		}
		quit <- 0
	}()

	routineFibonacci(c, quit)

	return result
}

func main() {
	js.Global().Set("wsGoFibonacci", js.FuncOf(goFibonacci))
	js.Global().Set("wsGoRoutineFibonacci", js.FuncOf(goRoutineFibonacci))

	select {}
}
