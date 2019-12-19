package main

import(
	"fmt"
	"os"
	"strconv"
)

func main()  {
	
	inputValues := os.Args[1:]
	numbers := make([]int, len(inputValues))

	for index, value := range inputValues {
		number, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("%s is not a valid number!\n", value)
			os.Exit(1)
		}
		numbers[index] = number
	}
	
	fmt.Println(quicksort(numbers))

}

func quicksort(numbers []int) []int  {

	if len(numbers) <= 1{
		return numbers
	}

	n := make([]int, len(numbers))
	copy(n, numbers)

	pivotIndex := len(n) / 2
	pivotValue := n[pivotIndex]

	n = append(n[:pivotIndex], n[pivotIndex+1:]...)

	lower, higher := slice(n, pivotValue)

	return append (
		append(quicksort(lower), pivotValue),
		quicksort(higher)...
	)
}

func slice(numbers []int, pivotValue int) (lower []int, higher []int)  {
	for _, value := range numbers {
		if value <= pivotValue {
			lower = append(lower, value)
		} else {
			higher = append(higher, value)
		}
	}
	
	return lower, higher
}