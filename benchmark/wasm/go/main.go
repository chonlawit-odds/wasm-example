//go:build js && wasm
// +build js,wasm

package main

import (
	"sync"
	"syscall/js"
)

// Fibonacci resursive algorithm
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

// Fibonacci tail recursive algorithm
func tailFibonacci(n, left, right uint) uint {
	if n == 0 {
		return left
	}

	return tailFibonacci(n-1, right, left+right)
}

func goTailFibonacci(_ js.Value, args []js.Value) any {
	n := uint(args[0].Int())

	return tailFibonacci(n, 0, 1)
}

// Fibonacci recursive goroutine algorithm
var memo sync.Map

func routineFibonacci(n uint, wg *sync.WaitGroup, resultChan chan<- uint) {
	defer wg.Done()
	if n <= 1 {
		resultChan <- n
		return
	}

	if val, ok := memo.Load(n); ok {
		resultChan <- val.(uint)
		return
	}

	var (
		leftChan  = make(chan uint, 1)
		rightChan = make(chan uint, 1)
		nestedWg  sync.WaitGroup
	)

	nestedWg.Add(2)
	go routineFibonacci(n-1, &nestedWg, leftChan)
	go routineFibonacci(n-2, &nestedWg, rightChan)
	nestedWg.Wait()

	close(leftChan)
	close(rightChan)

	left, right := <-leftChan, <-rightChan
	result := left + right

	memo.Store(n, result)
	resultChan <- result
}

func goRoutineFibonacci(_ js.Value, args []js.Value) any {
	var (
		n          = uint(args[0].Int())
		wg         sync.WaitGroup
		resultChan = make(chan uint, 1)
	)

	memo = sync.Map{}

	wg.Add(1)
	go routineFibonacci(n, &wg, resultChan)
	wg.Wait()

	close(resultChan)

	return <-resultChan
}

// Fibonacci tail recursive goroutine algorithm
func routineTailFibonacci(n uint, left, right uint, resultChan chan<- uint) {
	if n == 0 {
		resultChan <- left
		return
	}

	go routineTailFibonacci(n-1, right, left+right, resultChan)
}

func goRoutineTailFibonacci(_ js.Value, args []js.Value) any {
	var (
		n          = uint(args[0].Int())
		resultChan = make(chan uint)
	)

	go routineTailFibonacci(n, 0, 1, resultChan)

	return <-resultChan
}

func main() {
	js.Global().Set("wsGoFibonacci", js.FuncOf(goFibonacci))
	js.Global().Set("wsGoRoutineFibonacci", js.FuncOf(goRoutineFibonacci))
	js.Global().Set("wsGoTailFibonacci", js.FuncOf(goTailFibonacci))
	js.Global().Set("wsGoRoutineTailFibonacci", js.FuncOf(goRoutineTailFibonacci))

	select {}
}
