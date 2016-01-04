package tests

import (
	"encoding/json"
	"github.com/himanshuo/gameoflife/models"
	"strings"
	"testing"
	"time"
	"fmt"
)



//things to test:
//subtask <->
//empty values for each each

type TestCase struct {
	taskObj models.Task
	taskJSON string // json
}

var testData []TestCase
var curTime time.Time
var curTimeJSON []byte
var emptyTask *models.Task
var emptyGoal *models.Goal


func init(){
	curTime = time.Now()
	curTimeJSON, _ = json.Marshal(curTime)
	emptyTask = &models.Task{}
	emptyGoal = &models.Goal{}
	emptyGoalJSON,_ := json.Marshal(emptyGoal)

	testData = []TestCase{
		//basic
		{
			models.Task{
				Id: 1,
				Name: "task1",
				Description: "test",
				FailureCriteria: "failT",
				AcceptanceCriteria: "acceptT",
				AchievementPoints: 0,
				Goals: []*models.Goal{},
				Progress: models.OPEN,
				Deadline: curTime,
				FailTime: curTime,
				AcceptTime: curTime,
				SubTasks: []*models.Task{},
				Recurring: false,
				RecurStart: curTime,
				RecurEnd: curTime,
				RecurInterval: 0,
			},
			fmt.Sprintf(`{"id":1,"name":"task1","description":"test","failureCriteria":"failT","acceptanceCriteria":"acceptT","achievementPoints":0,"goals":[],"progress":0,"deadline":%[1]s,"failTime":%[1]s,"acceptTime":%[1]s,"recurring":false,"recurStart":%[1]s,"recurEnd":%[1]s,"recurInterval":0}`, curTimeJSON),
		},
		//subtasks
		{
		models.Task{
			Id: 1,
			Name: "task1",
			Description: "test",
			FailureCriteria: "failT",
			AcceptanceCriteria: "acceptT",
			AchievementPoints: 0,
			Goals: []*models.Goal{},
			Progress: models.OPEN,
			Deadline: curTime,
			FailTime: curTime,
			AcceptTime: curTime,
			SubTasks: []*models.Task{emptyTask, emptyTask},
			Recurring: false,
			RecurStart: curTime,
			RecurEnd: curTime,
			RecurInterval: 0,
		},
			fmt.Sprintf(`{"id":1,"name":"task1","description":"test","failureCriteria":"failT","acceptanceCriteria":"acceptT","achievementPoints":0,"goals":[],"progress":0,"deadline":%[1]s,"failTime":%[1]s,"acceptTime":%[1]s,"recurring":false,"recurStart":%[1]s,"recurEnd":%[1]s,"recurInterval":0}`, curTimeJSON),
		},
		//goals
		{
		models.Task{
			Id: 1,
			Name: "task1",
			Description: "test",
			FailureCriteria: "failT",
			AcceptanceCriteria: "acceptT",
			AchievementPoints: 0,
			Goals: []*models.Goal{emptyGoal, emptyGoal},
			Progress: models.OPEN,
			Deadline: curTime,
			FailTime: curTime,
			AcceptTime: curTime,
			SubTasks: []*models.Task{emptyTask, emptyTask},
			Recurring: false,
			RecurStart: curTime,
			RecurEnd: curTime,
			RecurInterval: 0,
		},
			fmt.Sprintf(`{"id":1,"name":"task1","description":"test","failureCriteria":"failT","acceptanceCriteria":"acceptT","achievementPoints":0,"goals":[%[2]s,%[2]s],"progress":0,"deadline":%[1]s,"failTime":%[1]s,"acceptTime":%[1]s,"recurring":false,"recurStart":%[1]s,"recurEnd":%[1]s,"recurInterval":0}`, curTimeJSON, emptyGoalJSON),
		},
	}
}


func TestTaskConversions(t *testing.T) {
	for i, test := range testData {

		//json to Task
		dec := json.NewDecoder(strings.NewReader(test.taskJSON))
		var resTask models.Task
		if err := dec.Decode(&resTask); err != nil {
			t.Errorf("TestTaskConversions %d: %s", i, err)
		}


		if !resTask.Equals(&test.taskObj){
			t.Errorf("TestTaskConversions %d:  \nEXPECTED: %s\nACTUAL:   %s", i, test.taskObj, resTask)
		}

		//Task to JSON
		actualJSON, err := json.Marshal(test.taskObj)
		if err!= nil{
			t.Errorf("TestTaskConversions %d: %s", i, err)
		}
		actualJSONString := string(actualJSON)
		if actualJSONString!= test.taskJSON {
			t.Errorf("TestTaskConversions %d: \nEXPECTED: %s\nACTUAL:   %s", i, test.taskJSON, actualJSON)
		}

	}
}


//todo: test negative cases
func TestTaskConstructor(t * testing.T){
	goal := &models.Goal{}
	myTask := models.NewTask(1, "n", "d", "f", "a", 2, []*models.Goal{goal})
	shouldBe := &models.Task{
		Id: 1,
		Name: "n",
		Description: "d",
		FailureCriteria: "f",
		AcceptanceCriteria: "a",
		AchievementPoints: 2,
		Goals: []*models.Goal{goal},
		Progress: models.OPEN,
		Deadline: time.Time{},
		FailTime: time.Time{},
		AcceptTime: time.Time{},
		SubTasks: []*models.Task{},
		Recurring: false,
		RecurStart: time.Time{},
		RecurEnd: time.Time{},
		RecurInterval: 0,
	}
	// shouldBe := &models.Task{1, "n", "d", "f", "a", time.Time{}, time.Time{}, time.Time{}, 2, []*models.Task{}, []*models.Goal{goal}, models.OPEN, false, time.Time{}, time.Time{}, 0}
	if myTask.EqualsRecursive(shouldBe){
		t.Errorf("TestTaskConstructor: EXPECT: %s, ACTUAL: %s", shouldBe, myTask)
	}
}
