package models

type ErrorType int

//Error Types
const (
	SERVER_ERROR ErrorType = iota
	INVALID_REQUEST
	TASK_NOT_FOUND
	ACCESS_VIOLATION
	TASK_ALREADY_EXISTS
)

type ApiInvalidResponse struct {
	Code      int       `json:"code"`
	ErrorType ErrorType `json:"error_type"`
	ErrorMsg  string    `json:"error_msg"`
}
