package taskTracker

type Task struct {
	Name    string
	Desc    string
	Urgency string
	Group   string
}

var TaskIndex = make(map[string]int)
var TaskList = make([]Task, 0)
