package internal

import (
	"fmt"
	"slices"
	task "taskTracker"
)

func AddTask(newTask task.Task) {

	task.TaskList = append(task.TaskList, newTask)
	task.TaskIndex[newTask.Name] = len(task.TaskIndex) - 1
	fmt.Printf("Task: %v Added to list", newTask.Name)
}
func DeleteTask(Oldtask task.Task) {
	i, exists := task.TaskIndex[Oldtask.Name]
	if !exists {
		fmt.Printf("Given Task does not exist!")
		return
	}
	task.TaskList = slices.Delete(task.TaskList, i, i)

}
func EditTask(oldTask task.Task, newTask task.Task) {

}
