package models

type Time struct {
	Time string
}

type Project struct {
	ProjectName string `json:"projectname"`
	ProjectId   int    `json:"projectid"`
}

type ProjectWorkflow struct {
	Tasks           *Task
	TaskId          []int
	Finish          *Time
	ProjectDeadline *Time
}
