package main

import (
	"fmt"
)

type JoinFunc func(n, m int) int

func JoinValues(
	values []int,
	initValue int,
	fn JoinFunc,
) int {
	finalValue := initValue

	for _, v := range values {
		finalValue = fn(finalValue, v)
	}

	return finalValue
}

func SumValues(values []int) int {
	sumFunc := func(n, m int) int {
		return n + m
	}
	return JoinValues(values, 0, sumFunc)
}

func MultiplyValues(values []int) int {
	multiplyFunc := func(n, m int) int {
		return n * m
	}
	return JoinValues(values, 1, multiplyFunc)
}

func main() {
	values := []int{2, -4, 1, 1, -2, -6, 25}

	fmt.Println("Sum: ", SumValues(values))
	fmt.Println("Product: ", MultiplyValues(values))
}
