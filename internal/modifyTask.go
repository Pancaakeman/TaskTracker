package internal

import (
	"fmt"
	"slices"
)

func AddTask(newTask Task) {

	TaskList = append(TaskList, newTask)
	TaskIndex[newTask.Name] = len(TaskIndex) - 1
	fmt.Printf("Task: %v Added to list\n", newTask.Name)
}
func DeleteTask(Oldtask Task) {
	i, exists := TaskIndex[Oldtask.Name]
	if !exists {
		fmt.Printf("Given Task does not exist!\n")
		return
	}
	TaskList = slices.Delete(TaskList, i, i)
	delete(TaskIndex, Oldtask.Name)

}
func EditTask(oldTask Task, newTask Task) {
	i, exists := TaskIndex[oldTask.Name]
	if !exists {
		fmt.Printf("Task to Edit doesn't exist")
		return
	}
	TaskList[i] = newTask
	delete(TaskIndex, oldTask.Name)
	TaskIndex[newTask.Name] = len(TaskIndex) - 1

}
