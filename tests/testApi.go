package tests

import (
	"testing"
	"net/http"
	"strings"
	"fmt"

)



func TestGetTasks(t *testing.T){
	_,err := http.Get("http://localhost:8080/")
	if err != nil{
		t.FailNow()
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
	body := strings.NewReader("name=afadf")
	req, err:= http.NewRequest("PUT","http://localhost:8080/task/", 
		body)
	if err != nil{
		t.Errorf("request not built properly")
	}
	req.Header.Add("Content-Type","application/x-www-form-urlencoded")
	resp, _:= http.DefaultClient.Do(req)
	if err != nil{
		t.Errorf("could not create task: %$", err)
	}
	fmt.Println(resp)


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

