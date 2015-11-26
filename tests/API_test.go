package tests

import (
	"testing"
	"net/http"
	"strings"
	"github.com/himanshuo/gameoflife/models"
	"encoding/json"
	"fmt"
)



func TestGetTasks(t *testing.T){
	_,err := http.Get("http://localhost:8080/")
	if err != nil{
		t.Errorf("TestGetTasks Failed with err %s", err)
	}
}


func TestCRUDTasks(t *testing.T){
	//create task
	// req, err := http.NewRequest("PUT","http://localhost:8080/tasks/", 
	// 	url.Values{"name": {"temp"}}
	// 	)
	// if err != nil{
	// 	t.FailNow()
	// }
	// resp, err := http.DefaultClient.Do(req)

	//create task
	name := "Test Task 1"
	body := strings.NewReader(fmt.Sprintf("name=%s", name))
	req, err:= http.NewRequest("PUT","http://localhost:8080/task/", 
		body)
	if err != nil{
		t.Errorf("request not built properly")
	}
	req.Header.Add("Content-Type","application/x-www-form-urlencoded")
	resp, err:= http.DefaultClient.Do(req)
	if err != nil{
		t.Errorf("could not create task: %$", err)
	}

	task := models.Task{}
	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil{
		t.Errorf("TestCRUDTasks got invalid response %s",  err)
	}
	if task.Name != name {
		t.Errorf("TestCRUDTasks returned invalid name: %s, %s", task.Name, resp.Body)
	}

	//read task
	//update task
	//read task
	//delete task
	//read task fails
	
}

// func TestCreateTasksWithSameName(t *testing.T){
// 	resp,err := http.Get("http://localhost:8080/")
// 	if err != nil{
// 		t.FailNow()
// 	}
// }

// func TestReadNonexistantTask(t *testing.T){
// 	resp,err := http.Get("http://localhost:8080/")
// 	if err != nil{
// 		t.FailNow()
// 	}
// }

// func TestUpdateNonexistantTask(t *testing.T){
// 	resp,err := http.Get("http://localhost:8080/")
// 	if err != nil{
// 		t.FailNow()
// 	}
// }


// func TestDeleteNonexistantTask(t *testing.T){
// 	resp,err := http.Get("http://localhost:8080/")
// 	if err != nil{
// 		t.FailNow()
// 	}
// }

