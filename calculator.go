package main

import (
	"fmt"
	"strconv"
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
	clearErrors()
	defer handleFailures()
	lhs, rhs, outputID, _ := getLHSAndRHS(values)
	sum, err := operation(lhs, rhs, "+")
	handleError(err)
	setValueOfJsField(outputID, sum)
}

func sub(values []js.Value) {
	clearErrors()
	defer handleFailures()
	lhs, rhs, outputID, _ := getLHSAndRHS(values)
	sub, err := operation(lhs, rhs, "-")
	handleError(err)
	setValueOfJsField(outputID, sub)
}

func mul(values []js.Value) {
	clearErrors()
	defer handleFailures()
	lhs, rhs, outputID, _ := getLHSAndRHS(values)
	mul, err := operation(lhs, rhs, "*")
	handleError(err)
	setValueOfJsField(outputID, mul)
}

func div(values []js.Value) {
	clearErrors()
	defer handleFailures()
	lhs, rhs, outputID, _ := getLHSAndRHS(values)
	div, err := operation(lhs, rhs, "/")
	handleError(err)
	setValueOfJsField(outputID, div)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func getLHSAndRHS(values []js.Value) (float64, float64, string, error) {
	if len(values) != 3 {
		return 0, 0, "", errors.New("Not required no of values received")
	}
	lhsID := values[0].String()
	rhsID := values[1].String()
	outputID := values[2].String()
	lhs, rhs := getValueOfJsField(lhsID), getValueOfJsField(rhsID)
	return lhs, rhs, outputID, nil
}

func setErrorValue(msg string) {
	js.Global().Get("document").Call("getElementById", "errors").Set("innerHTML", msg)
}

func clearErrors() {
	setErrorValue("")
}

func handleFailures() {
	if err := recover(); err != nil {
		setErrorValue(fmt.Sprintf("%v", err))
	}
}

func getValueOfJsField(fieldID string) float64 {
	valStr := js.Global().Get("document").Call("getElementById", fieldID).Get("value").String()
	valFloat, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		panic(errors.New("Only float value accepted"))
	}
	return valFloat
}

func setValueOfJsField(fieldID string, value float64) {
	js.Global().Get("document").Call("getElementById", fieldID).Set("value", value)
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
