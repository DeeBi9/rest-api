package models

// Admin will have all the permission to add remove a user
// Admin will also have the permission to assign task to users
type Information struct {
	Post string `json:"post"`
	Name string `json:"name"`
	Id   int    `json:"id"`
}

// Manager will have the permission to assign task to users
type Manager struct {
	Designation    string `default:"Manager"`
	Name           string `json:"manager"`
	ManagerId      int
	ManagerProject *Project
}

// Doesn't have special permission but to edit its name
type Employee struct {
	Designation     string `default:"Employee"`
	Name            string `json:"name"`
	EmployeeID      int    `json:"id"`
	EmployeeProject *Project
}

type Assignees struct {
	Information Information `json:"information"`
	Project     Project     `json:"project"`
}
