package main 

import("fmt")

type Operation interface {
	Calc() int
}

type Sum struct {
	num1, num2 int
}

func (s Sum) Calc() int {
	return s.num1 + s.num2
}

func (s Sum) String() string {
	return fmt.Sprintf("%d + %d = %d", s.num1, s.num2, s.Calc())
}

type Subtraction struct {
	num1, num2 int
}

func (s Subtraction) Calc() int {
	return s.num1 - s.num2
}

func (s Subtraction) String() string {
	return fmt.Sprintf("%d - %d = %d", s.num1, s.num2, s.Calc())
}

func main(){
	opList := make([]Operation, 4)
	opList[0] = Sum{1,3}
	opList[1] = Subtraction{1,8}
	opList[2] = Sum{10,17}
	opList[3] = Subtraction{28,4}

	for _, operation := range opList {
		fmt.Println(operation)
	}
}