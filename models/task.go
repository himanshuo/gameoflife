package models

import (
	"time"
//	"encoding/json"
//	"strconv"
	//"fmt"
)


type Task struct {
	//Required
	Id   int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	FailureCriteria string `json:"failureCriteria"`
	AcceptanceCriteria string `json:"acceptanceCriteria"`
	AchievementPoints int `json:"achievementPoints"`
	Goals []*Goal `json:"goals"`
	Progress Progress `json:"progress"`

	//Required if Active
	Deadline time.Time `json:"deadline"`
	FailTime time.Time `json:"failTime"`
	AcceptTime time.Time `json:"acceptTime"`

	//optional

	//SubTasks is a list NOT a set because the order of the subtasks can matter
	//ignoring subtasks in json conversions
	SubTasks []*Task `json:"-"`

	//todo: handle recurring in a smart way
	Recurring bool `json:"recurring"`
	RecurStart time.Time `json:"recurStart"`
	RecurEnd time.Time `json:"recurEnd"`
	//time.Duration is just int64
	//todo: unclear whether json conversion handles Duration->int64->string
	//todo: if not, will have to make marshal and unmarshal functions as for JSONtime
	RecurInterval  time.Duration `json:"recurInterval"`
}


func NewTask(id int, name string, desc string, failureCriteria string, acceptanceCriteria string, achievementPoints int, goals []*Goal) *Task {
    if id < 0 {
        return nil
    }
		if name == "" || desc == ""{
			return nil
		}
		if failureCriteria == "" || acceptanceCriteria == "" {
			return nil
		}
		if achievementPoints < 0 {
			return nil
		}
		if len(goals) == 0 {
			return nil
		}

		timeZeroValue := time.Time{}
		return &Task{
			Id: id,
			Name: name,
			Description: desc,
			FailureCriteria: failureCriteria,
			Deadline: timeZeroValue,
			FailTime: timeZeroValue,
			AcceptTime: timeZeroValue,
			AchievementPoints: achievementPoints,
			SubTasks: make([]*Task, 0),
			Goals:goals,
			Progress: OPEN,
			Recurring:false,
			RecurStart:timeZeroValue,
			RecurEnd:timeZeroValue,
			RecurInterval: 0,
		}
}


func (t * Task) Equals(other * Task) bool{
	if other == nil {
		return false
	}
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
