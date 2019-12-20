package main

import(
	"errors"
	"fmt"
)

func main()  {
	stack := Stack{}
	fmt.Println("Stack created. Length: ", stack.Length())
	fmt.Println("Is Empty? ", stack.IsEmpty())

	stack.Push("Go")
	stack.Push(1234)
	stack.Push(3.14)
	stack.Push("End")
	fmt.Println("Stack Length: ", stack.Length())
	fmt.Println("Is Empty? ", stack.IsEmpty())

	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		fmt.Println("Value: ", value)
		fmt.Println("Length: ", stack.Length())
		fmt.Println("Is Empty? ", stack.IsEmpty())
	}

	_, err := stack.Pop()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

type Stack struct {
	values []interface {}
}

func (stack Stack) Length() int {
	return len(stack.values)
}

func (stack Stack) IsEmpty() bool {
	return stack.Length() == 0
}

func (stack *Stack) Push(value interface{}) {
	stack.values = append(stack.values, value)
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Empty stack")
	}
	value := stack.values[stack.Length() - 1]
	stack.values = stack.values[:stack.Length() - 1]
	return value, nil
}