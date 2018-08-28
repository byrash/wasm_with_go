package main

import (
	"fmt"
	"syscall/js"

	"github.com/pkg/errors"
)

func main() {
	dummyChann := make(chan int, 0)
	fmt.Println("WASM initialised")
	mapJsMethodsToGoFuncs()
	<-dummyChann // Wait for ever;
	/*
		Go program should be running for JS registred actions/operations like add to be available for user actions lik button clicks.

				Uncaught Error: bad callback: Go program has already exited
			    at global.Go._resolveCallbackPromise (wasm_exec.js:378)
			    at wasm_exec.js:394
			    at HTMLButtonElement.onclick ((index):39)
	*/
}

func mapJsMethodsToGoFuncs() {
	js.Global().Set("add", js.NewCallback(add))
	js.Global().Set("sub", js.NewCallback(sub))
	js.Global().Set("mul", js.NewCallback(mul))
	js.Global().Set("div", js.NewCallback(div))
}

func add(values []js.Value) {
	lhs, rhs, _ := getLhsAndRhs(values)
	sum, _ := operation(lhs, rhs, "+")
	js.Global().Set("output", js.ValueOf(sum))
}

func sub(values []js.Value) {
	lhs, rhs, _ := getLhsAndRhs(values)
	sub, _ := operation(lhs, rhs, "-")
	js.Global().Set("output", js.ValueOf(sub))
}

func mul(values []js.Value) {
	lhs, rhs, _ := getLhsAndRhs(values)
	mul, _ := operation(lhs, rhs, "*")
	js.Global().Set("output", js.ValueOf(mul))
}

func div(values []js.Value) {
	lhs, rhs, _ := getLhsAndRhs(values)
	div, _ := operation(lhs, rhs, "/")
	js.Global().Set("output", js.ValueOf(div))
}

func getLhsAndRhs(values []js.Value) (float64, float64, error) {
	if len(values) != 2 {
		return 0, 0, errors.New("Not required no of values received")
	}
	lhs, rhs := values[0].Float(), values[1].Float()
	return lhs, rhs, nil
}

func operation(lhs, rhs float64, op string) (float64, error) {
	switch op {
	case "+":
		return float64(lhs + rhs), nil
	case "-":
		return float64(lhs - rhs), nil
	case "*":
		return float64(lhs * rhs), nil
	case "/":
		return float64(lhs / rhs), nil
	default:
		return 0, errors.New(fmt.Sprintf("Invalid Operation %v", op))
	}
}
