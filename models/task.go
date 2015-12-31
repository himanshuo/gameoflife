package models

import (
	"time"
)


type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	FailureCriteria string `json:"failureCriteria"`
	AcceptanceCriteria string `json:"acceptanceCriteria"`
	Deadline JSONTime `json:"deadline"` 
	FailTime JSONTime `json:"failTime"` 
	AcceptTime JSONTime `json:"acceptTime"` 
	AchievementPoints int `json:"achievementPoints"`
	SubTasks []Task `json:"subtasks"`
	Goals []Goal `json:"goals"`
	Progress Progress `json:"progress"`
	//todo: handle recurring in a smart way
	Recurring bool `json:"recurring"`
	RecurStart JSONTime `json:"recurStart"`
	RecurEnd JSONTime `json:"recurEnd"`
	//time.Duration is just int64
	//todo: unclear whether json conversion handles Duration->int64->string
	//todo: if not, will have to make marshal and unmarshal functions as for JSONtime
	RecurInterval  time.Duration `json:"recurInterval"`
}
