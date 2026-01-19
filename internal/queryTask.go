package internal

import (
	"fmt"
)

func ListTasks() {
	fmt.Println("-------------Listing Task Names---------------")
	if len(TaskList) == 0 {
		fmt.Printf("No Tasks to List! \n")
		return
	}

	for i, task := range TaskList {
		fmt.Printf("%v Task Name: %v \nTask Description: %v \nTask Urgency: %v \nTask Group: %v \n", i, task.Name, task.Desc, task.Urgency, task.Group)
	}
}
func GetTaskWithName(name string) int {
	i, exists := TaskIndex[name]
	if !exists {
		fmt.Printf("Given Task %v does not exist")
	}
	return i
}
func ListRangeOfTasks(lowerBound int, upperBound int) []Task {
	return TaskList[lowerBound:upperBound]

}
