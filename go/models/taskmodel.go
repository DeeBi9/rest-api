package models

type Task struct {
	AssignTo     *Employee
	TaskName     string
	TaskId       int
	ToDo         []string `json:"todo"`
	InProgress   bool
	Completed    bool
	Finish       *Time
	TaskDeadline *Time
}
