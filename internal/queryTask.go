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
