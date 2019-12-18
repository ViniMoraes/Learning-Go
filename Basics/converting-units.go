package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {
	
	if (len(os.Args) < 3){
		fmt.Println("How to use: convert <values> <unit>")
		os.Exit(1)
	}

	sourceUnit := os.Args[len(os.Args) - 1]
	sourceValues := os.Args[1 : len(os.Args) -1]

	var targetUnit string

	if sourceUnit == "celsius" {
		targetUnit = "fahrenheit"
	} else if sourceUnit == "quilometers" {
		targetUnit = "miles"
	} else {
		fmt.Printf("%s don't match a valid unit.\n", sourceUnit)
		os.Exit(1)
	}

	for index, value := range sourceValues {
		sourceValue, err := strconv.ParseFloat(value, 64)
		var targetValue float64
		if err != nil {
			fmt.Printf("The value %s in position %d is not a valid number!\n", value, index)
			os.Exit(1)
		}

		if sourceUnit == "celsius" {
			targetValue = sourceValue * 1.8 + 32
		} else {
			targetValue = sourceValue / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n", sourceValue, sourceUnit, targetValue, targetUnit)
	}
}