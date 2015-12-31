package tests

import (
	"encoding/json"
	"fmt"
	"github.com/himanshuo/gameoflife/models"
	"net/http"
	"runtime"
	"strings"
	"testing"
)


//helper function in order to check errors
func checkError(t *testing.T, err error) {
	if err != nil {
		//get caller function
		pc := make([]uintptr, 10) // at least 1 entry needed
		runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[0])
		_, line := f.FileLine(pc[0])
		funcName := strings.Split(f.Name(), ".")[1]
		t.Errorf("[%s][line %d]:%s:%s", funcName, line, err)
	}
}

//things to test:
//basic json <-> Task
//JSONTime json <-> task
//Duration task <-> json
//subtask <->
//empty values for each each 

//HOW TO: json to map
//var dat map[string]interface{}
//    if err := json.Unmarshal(byt, &dat); err != nil {
 //       panic(err)
  //  }
  //  fmt.Println(dat)
//    num := dat["num"].(float64)
  //  fmt.Println(num)

var testData = []struct {
	taskObj Task
	taskJSON string // json
}{
	{ 
		Task{1, "task1", "test task", "fail criteria no exist", "accept criteria no exist", "0","0","0",0, []Task{}, []Goal{}, Progress.OPEN, false, "0", "0", 0},
		`{id: 1, name: "task1", description:"test task", failureCriteria:"fail criteria no exist", acceptanceCriteria:"accept criteria no exist", deadline:"0",failTime:"0",acceptTime:"0",achievementPoints:0, subtasks: [], goals:[], progress:0, recurring:false, recurStart: "0", recurEnd:"0", recurInterval:0}`
	}
	// {"a_name","a_author","", false, true},            // simple
	// {"a_name","a_author","0", false, true},            // 0 version
	// {"a_name","a_author","1", false, true},            // version
	// {"a_name","a_author","0.9.1", false, true},        // complex version
	
	// {"s", "s", "0", false, true},					   // single char
	// {" 1234567890/.,;'[]\"@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char
		
	// {"a_name","a_author","0", true, true},            // simple multipath
	// {"a_name","a_author","1", true, true},            // version multipath
	// {"a_name","a_author","0.9.1", true, true},        // complex version multipath
	// {"s", "s", "0", false, true},					   // single char multipath
	// {"1234567890/.,;'[]\" ~!@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char multipath
	
	// //bad
	// {".", ".", "0", false, false},					   // dot input
	// {"..", "..", "0", false, false},				   // double dot input
	// {"//", "//", "0", false, false},				   // double slash
	// {"a_name","a_author","", false, true},            // simple
	// {"a_name","a_author","0", false, true},            // simple 0 version
	// {"a_name","a_author","1", false, true},            // version
	// {"a_name","a_author","0.9.1", false, true},        // complex version
	
	// {"s", "s", "0", false, true},					   // single char
	// {" 1234567890/.,;'[]\"@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char
		
	// {"a_name","a_author","0", true, true},            // simple multipath
	// {"a_name","a_author","1", true, true},            // version multipath
	// {"a_name","a_author","0.9.1", true, true},        // complex version multipath
	// {"s", "s", "0", false, true},					   // single char multipath
	// {"1234567890/.,;'[]\" ~!@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char multipath
	
	// //bad multipath
	// {".", ".", "0", true, false},					   // dot input
	// {"..", "..", "0", true, false},				   // double dot input
	// {"//", "//", "0", true, false},				   // double slash
	
	// //good single path
	// {"a_name","a_author","0", false, true},            // simple
	// {"a_name","a_author","1", false, true},            // version
	// {"a_name","a_author","0.9.1", false, true},        // complex version
	
	// {"s", "s", "0", false, true},					   // single char
	// {" 1234567890/.,;'[]\"@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char
		
	// {"a_name","a_author","0", true, true},            // simple multipath
	// {"a_name","a_author","1", true, true},            // version multipath
	// {"a_name","a_author","0.9.1", true, true},        // complex version multipath
	// {"s", "s", "0", false, true},					   // single char multipath
	// {"1234567890/.,;'[]\" ~!@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char multipath
	
	// //bad single path
	// {".", ".", "0", false, false},					   // dot input
	// {"..", "..", "0", false, false},				   // double dot input
	// {"//", "//", "0", false, false},				   // double slash
	
}

//test all crud options
func TestTaskConversions(t *testing.T) {
	for i, test := range testData{
		//json to Task
		dec := json.NewDecoder(strings.NewReader(jsonStream))
		var resTask Task
		if err := dec.Decode(&resTask); err != nil {
			t.Errorf("TestTaskConversions %d:  EXPECTED: %s->%s ACTUAL: %s->%s ERROR:", i, test.taskObj, test.taskJSON, test.taskObj, resTask, err)
		}

		//Task to JSON
		
}