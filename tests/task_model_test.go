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

func init(){
	curTime = time.Now()
	curTimeJSON, _ = json.Marshal(curTime)
	emptyTask = &models.Task{}

	testData = []TestCase{
		//basic
		{
			models.Task{1, "task1", "test", "failT", "acceptT", curTime,curTime,curTime,0, []*models.Task{}, []*models.Goal{}, models.OPEN, false, curTime, curTime, 0},
			fmt.Sprintf(`{"id":1,"name":"task1","description":"test","failureCriteria":"failT","acceptanceCriteria":"acceptT","deadline":%[1]s,"failTime":%[1]s,"acceptTime":%[1]s,"achievementPoints":0,"goals":[],"progress":0,"recurring":false,"recurStart":%[1]s,"recurEnd":%[1]s,"recurInterval":0}`, curTimeJSON),
		},
		//subtasks
		{
			models.Task{1, "task1", "test", "failT", "acceptT", curTime,curTime,curTime,0, []*models.Task{emptyTask,emptyTask}, []*models.Goal{}, models.OPEN, false, curTime, curTime, 0},
			fmt.Sprintf(`{"id":1,"name":"task1","description":"test","failureCriteria":"failT","acceptanceCriteria":"acceptT","deadline":%[1]s,"failTime":%[1]s,"acceptTime":%[1]s,"achievementPoints":0,"goals":[],"progress":0,"recurring":false,"recurStart":%[1]s,"recurEnd":%[1]s,"recurInterval":0}`, curTimeJSON, emptyTask),
		},
		//goals
		{
			models.Task{1, "task1", "test", "failT", "acceptT", curTime,curTime,curTime,0, []*models.Task{emptyTask,emptyTask}, []*models.Goal{}, models.OPEN, false, curTime, curTime, 0},
			fmt.Sprintf(`{"id":1,"name":"task1","description":"test","failureCriteria":"failT","acceptanceCriteria":"acceptT","deadline":%[1]s,"failTime":%[1]s,"acceptTime":%[1]s,"achievementPoints":0,"goals":[],"progress":0,"recurring":false,"recurStart":%[1]s,"recurEnd":%[1]s,"recurInterval":0}`, curTimeJSON, emptyTask),
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
			t.Errorf("TestTaskConversions %d:  \nEXPECTED: %s\nACTUAL:     %s", i, test.taskObj, resTask)
		}

		//Task to JSON
		actualJSON, err := json.Marshal(test.taskObj)
		if err!= nil{
			t.Errorf("TestTaskConversions %d: %s", i, err)
		}
		actualJSONString := string(actualJSON)
		if actualJSONString!= test.taskJSON {
			t.Errorf("TestTaskConversions %d: \nEXPECTED: %s\nACTUAL:     %s", i, test.taskJSON, actualJSON)
		}

	}
}


//todo: test negative cases
func TestTaskConstructor(t * testing.T){
	goal := &models.Goal{}
	myTask := models.NewTask(1, "n", "d", "f", "a", 2, []*models.Goal{goal})
	shouldBe := &models.Task{1, "n", "d", "f", "a", time.Time{}, time.Time{}, time.Time{}, 2, []*models.Task{}, []*models.Goal{goal}, models.OPEN, false, time.Time{}, time.Time{}, 0}
	if myTask.EqualsRecursive(shouldBe){
		t.Errorf("TestTaskConstructor: EXPECT: %s, ACTUAL: %s", shouldBe, myTask)
	}
}
