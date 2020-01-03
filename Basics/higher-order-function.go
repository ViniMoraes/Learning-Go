package main

import (
	"fmt"
	"time"
)

func GenerateFib(n int) func() {
	return func() {
		a, b := 0, 1

		fib := func() int {
			a, b = b, a+b

			return a
		}

		for i := 0; i < n; i++ {
			fmt.Printf("%d ", fib())
		}
		fmt.Println()
	}
}

func CalcTime(function func()) {
	startTime := time.Now()

	function()

	fmt.Printf("Execution Time: %s\n\n", time.Since(startTime))
}

func main() {
	CalcTime(GenerateFib(10))
	CalcTime(GenerateFib(30))
	CalcTime(GenerateFib(80))
}
