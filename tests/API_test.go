package tests
import (
	"encoding/json"
	"fmt"
	"github.com/himanshuo/gameoflife/models"
	"net/http"
	"strings"
	"testing"
	"runtime"
)

//helper function in order to check errors
func checkError(t *testing.T, err error){
    	if err != nil {
    		//get caller function
		pc := make([]uintptr, 10)  // at least 1 entry needed
	    	runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[0])
	    	_, line := f.FileLine(pc[0])
	    	funcName := strings.Split(f.Name(),".")[1]
		t.Errorf("[%s][line %d]:%s:%s", funcName, line, err)    		
    	}
}

//test home view shows something
func TestHomeView(t *testing.T) {
	_,err := http.Get("http://localhost:8080/")
	checkError(t, err)
}

//helper to view a given task
func readTask(t *testing.T, task models.Task) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/task/%d/", task.Id))
	checkError(t, err)
	readTask := models.Task{}
	err = json.NewDecoder(resp.Body).Decode(&readTask)
	checkError(t, err)
	if readTask.Id != task.Id {
		t.Errorf("TestCRUDTasks read got invalid id %s", readTask.Id)
	}
	if readTask.Name != task.Name {
		t.Errorf("TestCRUDTasks read got invalid name %s", readTask.Name)
	}
}

//helper to create a task given a name
func createTask(t *testing.T, name string) models.Task {
	task := models.Task{}
	body := strings.NewReader(fmt.Sprintf("name=%s", name))
	req, err := http.NewRequest("PUT", "http://localhost:8080/task/",
		body)
	checkError(t, err)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	checkError(t, err)
	err = json.NewDecoder(resp.Body).Decode(&task)
	checkError(t, err)
	if task.Name != name {
		t.Errorf("TestCRUDTasks returned invalid name: %s, %s", task.Name, resp.Body)
	}
	return task

}

//helper to update a task given its id and a new name
func updateTask(t *testing.T, task models.Task) models.Task {
	updateUrl := fmt.Sprintf("http://localhost:8080/task/%d/", task.Id)
	body := strings.NewReader(fmt.Sprintf("name=%s", task.Name))
	req, err := http.NewRequest("POST", updateUrl, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	checkError(t, err)
	resp, err := http.DefaultClient.Do(req)
	checkError(t, err)
	updatedTask := models.Task{}
	err = json.NewDecoder(resp.Body).Decode(&updatedTask)
	checkError(t, err)
	if updatedTask.Id != task.Id {
		t.Errorf("TestCRUDTasks update task did not have correct id %d", updatedTask.Id)
	}
	if updatedTask.Name != task.Name {
		t.Errorf("TestCRUDTasks update task did not have correct name %d", updatedTask.Name)
	}
	return updatedTask
}

//helper to delete a task
func deleteTask(t *testing.T, task models.Task) {
	deleteUrl := fmt.Sprintf("http://localhost:8080/task/%d/", task.Id)
	req, err := http.NewRequest("DELETE", deleteUrl, nil)
	checkError(t, err)
	_, err = http.DefaultClient.Do(req)
	checkError(t, err)
}

//test all crud options
func TestCRUDTasks(t *testing.T) {
	task := createTask(t, "Test Task 1")
	readTask(t, task)
	task.Name = "Test Task 2"
	task = updateTask(t, task)
	readTask(t, task)
	deleteTask(t, task)
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
