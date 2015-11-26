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
	//todo: a url_for(NAME) module would be helpful. I think the gorilla thing might already have it.
	name := "Test Task 1"
	var id int
	task := models.Task{}

	//create task
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

	
	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil{
		t.Errorf("TestCRUDTasks got invalid response %s",  err)
	}
	if task.Name != name {
		t.Errorf("TestCRUDTasks returned invalid name: %s, %s", task.Name, resp.Body)
	}
	id = task.Id

	//read task
	resp,err = http.Get(fmt.Sprintf("http://localhost:8080/task/%d/", id))
	if err != nil{
		t.Errorf("TestCRUDTasks got invalid read response %s: %s", err, resp)
	}
	//nullify task
	task = models.Task{}

	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil{
		t.Errorf("TestCRUDTasks got invalid response %s",  err)
	}
	if task.Id != id{
		t.Errorf("TestCRUDTasks read got invalid id %s", task.Id)
	} 
	if task.Name != name{
		t.Errorf("TestCRUDTasks read got invalid name %s", task.Name)
	} 

	//update task
	name = "New Test Task Name 1"
	updateUrl := fmt.Sprintf("http://localhost:8080/task/%d/", id)
	body = strings.NewReader(fmt.Sprintf("name=%s", name))
    	req, err = http.NewRequest("POST", updateUrl, body)
    	req.Header.Add("Content-Type","application/x-www-form-urlencoded")
    	if err != nil{
    		t.Errorf("TestCRUDTasks invalid update request %s", err)
    	}
    	
    	resp, err = http.DefaultClient.Do(req)
    	if err != nil{
    		t.Errorf("TestCRUDTasks invalid update response %s", err)
    	}
    	//nullify task
	task = models.Task{}

	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil{
		t.Errorf("TestCRUDTasks could not decode updated task %s", err)
	}
	if task.Id != id {
		t.Errorf("TestCRUDTasks update task did not have correct id %d", task.Id)
	}
	if task.Name != name {
		t.Errorf("TestCRUDTasks update task did not have correct name %d", task.Name)
	}


	//read task
	resp,err = http.Get(fmt.Sprintf("http://localhost:8080/task/%d/", id))
	if err != nil{
		t.Errorf("TestCRUDTasks got invalid read response %s: %s", err, resp)
	}
	//nullify task
	task = models.Task{}

	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil{
		t.Errorf("TestCRUDTasks got invalid response %s",  err)
	}
	if task.Id != id{
		t.Errorf("TestCRUDTasks read got invalid id %s", task.Id)
	} 
	if task.Name != name{
		t.Errorf("TestCRUDTasks read got invalid name %s", task.Name)
	} 


	//delete task
	deleteUrl := fmt.Sprintf("http://localhost:8080/task/%d/", id)
	req, err = http.NewRequest("DELETE", deleteUrl, nil)
    	if err != nil{
    		t.Errorf("TestCRUDTasks invalid delete request %s", err)
    	}
    	
    	resp, err = http.DefaultClient.Do(req)
    	if err != nil{
    		t.Errorf("TestCRUDTasks invalid delete response %s", err)
    	}
    	
	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil{
		t.Errorf("TestCRUDTasks could not decode updated task %s", err)
	}
	if task.Id != id {
		t.Errorf("TestCRUDTasks update task did not have correct id %d", task.Id)
	}
	if task.Name != name {
		t.Errorf("TestCRUDTasks update task did not have correct name %d", task.Name)
	}
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

