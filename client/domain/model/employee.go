package model

import (
	"time"
)

// EmployeeRequestParam employee request param
type EmployeeRequestParam struct {
	Name        string
	Address     string
	Phone       string
	BirthOfDate time.Time
}

// EmployeeResponseData employee response param
type EmployeeResponseData struct {
	UUID    string
	Message string
}
