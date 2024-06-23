package models

type EmployeeType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Doesn't need to update or obtain anything, but still here is the model
