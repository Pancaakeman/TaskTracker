package main

import "taskTracker/internal"

func main() {
	myTask := internal.Task{
		Name:    "Task1",
		Desc:    "Task but its 1",
		Urgency: "Now or never",
		Group:   "idk",
	}
	internal.AddTask(myTask)
	internal.ListTasks()
}
