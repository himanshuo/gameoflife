package models

import (
	"time"
	"encoding/json"
	"strconv"
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
	SubTasks []*Task `json:"subtasks"`
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
	if len(t.SubTasks) != len(other.SubTasks){
		return false
	}
	for i, _ := range t.SubTasks{
		if !t.SubTasks[i].Equals(other.SubTasks[i]){
			return false
		}
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



func (t *Task) MarshalJSON() ([]byte, error) {
	data := map[string]string{}
	data["id"] = strconv.Itoa(t.Id)
	data["name"] = t.Name
	data["description"] = t.Description
	data["failureCriteria"] = t.FailureCriteria
	data["acceptanceCriteria"] = t.AcceptanceCriteria
	if deadline, err := json.Marshal(t.Deadline); err != nil {
		return nil, err
	} else {
		data["deadline"] = string(deadline)
	}
	if failTime, err := json.Marshal(t.FailTime); err != nil {
		return nil, err
	} else {
		data["failTime"] = string(failTime)
	}
	if acceptTime, err := json.Marshal(t.AcceptTime); err != nil {
		return nil, err
	} else {
		data["acceptTime"] = string(acceptTime)
	}
	data["achievementPoints"] = strconv.Itoa(t.AchievementPoints)
	//todo: figure out how you're going to serialize subtasks
	data["subtasks"] = "[]"
	if goals, err := json.Marshal(t.Goals); err != nil {
		return nil, err
	} else {
		data["goals"] = string(goals)
	}
	data["progress"] = progressValues[t.Progress]
	data["recurring"] = strconv.FormatBool(t.Recurring)
	if recurStart, err := json.Marshal(t.RecurStart); err != nil {
		return nil, err
	} else {
		data["recurStart"] = string(recurStart)
	}
	if recurEnd, err := json.Marshal(t.RecurEnd); err != nil {
		return nil, err
	} else {
		data["recurEnd"] = string(recurEnd)
	}
	if recurInterval, err := json.Marshal(t.RecurInterval); err != nil {
		return nil, err
	} else {
		data["recurInterval"] = string(recurInterval)
	} 

	return json.Marshal(data) 
}

// func (t *JSONTime) UnmarshalJSON(b []byte) error {
// 	//use the []byte as a string
// 	//convert string to integer Unix time
// 	ts, err := strconv.Atoi(string(b))
// 	if err != nil {
// 		return err
// 	}
// 	//convert the unix time from int to int64
// 	//convert int64 unix time into local time.
// 	//0 sec offset from epoch
// 	*t = JSONTime(time.Unix(int64(ts), 0))

// 	return nil
// }



