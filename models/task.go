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
	Deadline time.Time `json:"deadline"` //todo: use JSONTime type
	FailTime time.Time `json:"failTime"` //todo: use JSONTime type
	AcceptTime time.Time `json:"acceptTime"` //todo: use JSONTime type
	AchievementPoints int `json:"achievementPoints"`
	SubTasks []Task `json:"subtasks"`
	Goals []Goal `json:"goals"`
	Progress Progress `json:"progress"`
	//todo: handle recurring in a smart way
	Recurring bool `json:"recurring"`
	RecurStart time.Time `json:"recurStart"`
	RecurEnd time.Time `json:"recurEnd"`
	RecurInterval time.Time `json:"recurInterval"`
}
