package models

type Progress int

//Error Types
const (
	OPEN Progress = iota
	WORKING
	FAILED
	ACCEPTED
)