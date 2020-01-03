package main

import ("fmt")

type CheckList []string

func printSlice(slice []string){
	fmt.Println("Slice: ", slice)
}

func printCheckList(checkList CheckList){
	fmt.Println("CheckList: ", checkList)
}

func main() {

	checkList := CheckList{"Task1", "Task1", "Task2"}
	normalSlice := []string{"Task1", "Task1", "Task2"}

	printSlice([]string(checkList))
	printCheckList(CheckList(normalSlice))
}