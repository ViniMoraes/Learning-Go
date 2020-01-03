package main

import(
	"fmt"
	"strings"
)

type checkList []string

func main(){

	list := make(checkList, 6)
	list[0] = "Task1 - Trivial"
	list[1] = "Task2 - Important"
	list[2] = "Task3 - Normal"
	list[3] = "Task4 - Normal"
	list[4] = "Task5 - Trivial"
	list[5] = "Task6 - Normal"

	trivialTasks, normalTasks, importantTasks := list.Sort()

	fmt.Println(trivialTasks, normalTasks, importantTasks)

}

func (list checkList) Sort() ([]string, []string, []string){
	var trivialTasks, normalTasks, importantTasks []string

	for _, e := range list {
		if strings.Contains(e, "Trivial"){
			trivialTasks = append(trivialTasks, e)
		} else if strings.Contains(e, "Normal"){
			normalTasks = append(normalTasks, e)
		} else if strings.Contains(e, "Important"){
			importantTasks = append(importantTasks, e)
		}
	}
	return trivialTasks, normalTasks, importantTasks
}