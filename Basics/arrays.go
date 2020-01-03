package main

import(
	"fmt"
)

func  main()  {
	var emptyArray [3]int
	numberArray := [5]int{1,2,3,4,5}
	primeNumberArray := [...]int{2,3,5,7,11,13}
	namesArray := [2]string{}

	fmt.Println(emptyArray, numberArray, primeNumberArray, namesArray)
	//Output:
	//[0 0 0] [1 2 3 4 5] [2 3 5 7 11 13] [ ]

	var multiA [2][2]int
	multiA[0][0], multiA[0][1] = 1, 4
	multiA[1][0], multiA[1][1] = 10, -2

	multiB := [2][2]int{{5,22}, {-9,-2}}

	fmt.Println("MultiA: ", multiA)
	fmt.Println("MultiB: ", multiB)
	//Output:
	//MultiA:  [[1 4] [10 -2]]
	//MultiB:  [[5 22] [-9 -2]]
}