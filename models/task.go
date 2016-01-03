package models

import (
	"time"
//	"encoding/json"
//	"strconv"
)


type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	FailureCriteria string `json:"failureCriteria"`
	AcceptanceCriteria string `json:"acceptanceCriteria"`
	Deadline time.Time `json:"deadline"`
	FailTime time.Time `json:"failTime"`
	AcceptTime time.Time `json:"acceptTime"`
	AchievementPoints int `json:"achievementPoints"`
	//SubTasks is a list NOT a set because the order of the subtasks can matter
	//ignoring subtasks in json conversions
	SubTasks []*Task `json:"-"`
	Goals []*Goal `json:"goals"`
	Progress Progress `json:"progress"`
	//todo: handle recurring in a smart way
	Recurring bool `json:"recurring"`
	RecurStart time.Time `json:"recurStart"`
	RecurEnd time.Time `json:"recurEnd"`
	//time.Duration is just int64
	//todo: unclear whether json conversion handles Duration->int64->string
	//todo: if not, will have to make marshal and unmarshal functions as for JSONtime
	RecurInterval  time.Duration `json:"recurInterval"`
}

func (t * Task) Equals(other * Task) bool{
	if t.Id != other.Id {
		return false
	}
	if t.Name != other.Name {
		return false
	}
	if t.Description != other.Description {
		return false
	}
	if t.FailureCriteria != other.FailureCriteria {
		return false
	}
	if t.AcceptanceCriteria != other.AcceptanceCriteria {
		return false
	}
	if t.Deadline != other.Deadline {
		return false
	}
	if t.FailTime != other.FailTime {
		return false
	}
	if t.AcceptTime != other.AcceptTime {
		return false
	}
	if t.AchievementPoints != other.AchievementPoints {
		return false
	}
	if len(t.Goals) != len(other.Goals){
		return false
	}
	for i, _ := range t.Goals{
		if !t.Goals[i].Equals(other.Goals[i]){
			return false
		}
	}
	if t.Progress != other.Progress {
		return false
	}
	if t.Recurring != other.Recurring {
		return false
	}
	if t.RecurStart != other.RecurStart {
		return false
	}
	if t.RecurEnd != other.RecurEnd {
		return false
	}
	if t.RecurInterval != other.RecurInterval {
		return false
	}
	return true
}

func (t * Task) EqualsRecursive(other * Task) bool{
	if !t.Equals(other){
		return false
	}
	if len(t.SubTasks) != len(other.SubTasks){
		return false
	}
	for i, _ := range t.SubTasks{
		if !t.SubTasks[i].EqualsRecursive(other.SubTasks[i]){
			return false
		}
	}
	return true
}
