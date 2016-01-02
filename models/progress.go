package models

type Progress int

const (
	OPEN Progress = iota
	WORKING
	FAILED
	ACCEPTED
)

var progressValues = [...]string{
	"OPEN",
	"WORKING",
	"FAILED",
	"ACCEPTED",
}