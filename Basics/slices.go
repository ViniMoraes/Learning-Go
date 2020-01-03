package main

import (
	"fmt"
)

func main() {

	//Article:
	//https://medium.com/rungo/the-anatomy-of-slices-in-go-6450e3bb2b94

	var emptySlice []int
	numberSlice := []int{1, 2, 3, 4, 5}
	namesSlice := []string{}

	fmt.Println(emptySlice, numberSlice, namesSlice)
	//Output:
	//[] [1 2 3 4 5] []

	//make signature:
	//func make ([]T, len, cap) []T

	arrayA := make([]int, 10)
	fmt.Println("arrayA: ", arrayA, len(arrayA), cap(arrayA))

	arrayB := make([]int, 10, 20)
	fmt.Println("arrayB: ", arrayB, len(arrayB), cap(arrayB))
}
