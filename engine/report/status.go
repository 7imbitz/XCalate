package report

type TaskStatus int

const (
	Unknown TaskStatus = iota
	Pass
	Fail
	Manual
)

// taskResults maps task name → status
var taskResults = make(map[string]TaskStatus)

// markTaskStatus sets the task to a specific TaskStatus
func MarkTaskStatus(task string, status TaskStatus) {
	taskResults[task] = status
}

// convenience: mark pass/fail using bool
func MarkTask(task string, ok bool) {
	if ok {
		MarkTaskStatus(task, Pass)
	} else {
		MarkTaskStatus(task, Fail)
	}
}
