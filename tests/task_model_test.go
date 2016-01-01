package tests

import (
	"encoding/json"
	"github.com/himanshuo/gameoflife/models"
	"strings"
	"testing"
	"time"
	"fmt"
	"reflect"
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
var emptyTaskJSON []byte

func init(){
	curTime = time.Now()
	curTimeJSON, _ = json.Marshal(curTime)
	emptyTaskJSON, _ = json.Marshal(models.Task{})

	testData = []TestCase{
		//basic
		{ 
			models.Task{1, "task1", "test", "failT", "acceptT", curTime,curTime,curTime,0, []models.Task{}, []models.Goal{}, models.OPEN, false, curTime, curTime, 0},
			fmt.Sprintf(`{"id":1,"name":"task1","description":"test","failureCriteria":"failT","acceptanceCriteria":"acceptT","deadline":%[1]s,"failTime":%[1]s,"acceptTime":%[1]s,"achievementPoints":0,"subtasks":[],"goals":[],"progress":0,"recurring":false,"recurStart":%[1]s,"recurEnd":%[1]s,"recurInterval":0}`, curTimeJSON),
		},
		//subtasks
		{ 
			models.Task{1, "task1", "test", "failT", "acceptT", curTime,curTime,curTime,0, []models.Task{models.Task{},models.Task{}}, []models.Goal{}, models.OPEN, false, curTime, curTime, 0},
			fmt.Sprintf(`{"id":1,"name":"task1","description":"test","failureCriteria":"failT","acceptanceCriteria":"acceptT","deadline":%[1]s,"failTime":%[1]s,"acceptTime":%[1]s,"achievementPoints":0,"subtasks":[%[2]s,%[2]s],"goals":[],"progress":0,"recurring":false,"recurStart":%[1]s,"recurEnd":%[1]s,"recurInterval":0}`, curTimeJSON, emptyTaskJSON),
		},
		
	}
}



//test all crud options
func TestTaskConversions(t *testing.T) {
	for i, test := range testData {
		//json to Task
		dec := json.NewDecoder(strings.NewReader(test.taskJSON))
		var resTask models.Task
		if err := dec.Decode(&resTask); err != nil {
			t.Errorf("TestTaskConversions %d: %s", i, err)
		}


		if !reflect.DeepEqual(resTask, test.taskObj){
			t.Errorf("TestTaskConversions %d:  \nEXPECTED: %s \nACTUAL:     %s", i, test.taskObj, resTask)
		}

		//Task to JSON
		actualJSON, err := json.Marshal(test.taskObj)
		if err!= nil{
			t.Errorf("TestTaskConversions %d: %s", i, err)
		}
		if string(actualJSON) != test.taskJSON {
			t.Errorf("TestTaskConversions %d: \nEXPECTED: %s \nACTUAL:     %s", i, test.taskJSON, actualJSON)
		}

	}
}