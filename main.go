package main

//todo: make sure using camelCase everywhere
import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/himanshuo/gameoflife/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/himanshuo/gameoflife/controller/apiclient"
)
//_ is in order to import a package solely for its side-effects at initialization.
	//In this case, go-sqlite3's side effects are allowing sqlite3 to be usable as a
	//database for  sql.Open

//the database resource
var db *sql.DB
var Router *mux.Router

//run at start of program. 
func init() {

	startDB()
	Router = mux.NewRouter()


}

func startDB() error{
	var err error
	//sqlite 3 database is stored in /data/data.db file
	db, err = sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}

	//this makes sure we can actually query the database.
	if err := db.Ping(); err != nil {
  		return err
	}
	//ASSUMPTION: table exists in database
	
	//statement creates Task table which only has a PK id and name textfields
	sqlStmt := "create table Task (id integer not null primary key autoincrement, name text);"
	//run statement to create table. return object of Result type is not needed.
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}

//Views
// Home page view. For now, simply lists all the tasks one by one.
func Home(w http.ResponseWriter, r *http.Request) {
	//get all tasks
	rows, err := db.Query("select id, name from Task")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//list of tasks
	tasks := []models.Task{}

	//iterate through response
	for rows.Next() {
		var id int
		var name string
		//fill id and name variables with db output
		rows.Scan(&id, &name)
		//create new controller model for task
		cur_task := models.Task{Id: id, Name: name}
		//add model to tasks list
		tasks = append(tasks, cur_task)
	}

	//choose which template to show user when this endpoint is called
	t, err := template.ParseFiles("views/static/templates/base.html")
	if err != nil{
		log.Fatal(err)
	}
	//add tasks to template
	//ASSUME that template is using the tasks properly
	t.Execute(w, tasks)

}

//API
//create a new task
//input: name as part of a x-www-form-urlencoded PUT request
//output: json encoded Task. Contains both id and name of newly created Task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	//note: r.FormValue searches for key in POST data fields, then PUT data fields
	//2 types of POST submissions: application/x-www-form-urlencoded AND multipart/form-data.
	// need to understand both. Generally speaking, urlencoded takes up extra space so is for normal post requests. multipart form-data does not increase space usage by a lot so is for uploading files
	//http://stackoverflow.com/a/4073451/4710047

	//expecting to come from PUT data field
	taskName := r.PostFormValue("name")

	//start transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// insert query
	stmt, err := tx.Prepare("insert into Task(name) values(?)")
	if err != nil {
		log.Fatal(err)
	}
	//close statement
	defer stmt.Close()
	//execute statement
	resp, err := stmt.Exec(taskName)
	if err != nil {
		log.Fatal(err)
	}
	//commit statement in db
	tx.Commit()

	//get last inserted task id in order to show it to the user
	//todo:  perhaps should do a query into the db to get the new Task
	newTaskId, err := resp.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	//create new task model
	newTask := models.Task{Id: int(newTaskId), Name: taskName}

	//encode newTask as json, and return it to user.
	if err := json.NewEncoder(w).Encode(newTask); err != nil {
		panic(err)
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	//note: r.FormValue searches for key in POST data fields, then PUT data fields
	//2 types of POST submissions: application/x-www-form-urlencoded AND multipart/form-data.
	// need to understand both. Generally speaking, urlencoded takes up extra space so is for normal post requests. multipart form-data does not increase space usage by a lot so is for uploading files
	//http://stackoverflow.com/a/4073451/4710047
	
	//todo: log for whenever any of the api methods is called

	log.Printf("called update task")
	//create a variable that has the parameters sent to the api in the url 
	// /task/<id>/<x>/  will lead to variables id and x in vars
	vars := mux.Vars(r)

	//todo: there should NOT be any newline between a command and its error validation

	//get task id
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	//get new name from POST data form
	newName := r.PostFormValue("name")

	//start the transaction (do not have to close transaction, but do have to commit them)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	//update query statement. 
	stmt, err := tx.Prepare("UPDATE Task SET  name=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	//close statement
	defer stmt.Close()
	//check for errors. response does not give you the row that had the update so _
	_, err := stmt.Exec(newName, taskId)
	if err != nil {
		log.Fatal(err)
	}
	//commit the transaction
	tx.Commit()

	//todo: perhaps should be querying db for task. There could be db triggers and
	//things that run internally

	//create model for updated task
	updatedTask := models.Task{Id: int(taskId), Name: newName}

	//encode task as json and return to user
	if err := json.NewEncoder(w).Encode(updatedTask); err != nil {
		panic(err)
	}
}

//view a single task
//input: id of task as url parameter
//output: json serialization of task with input id
func ViewTask(w http.ResponseWriter, r *http.Request) {
	//note: r.FormValue searches for key in GET queries, then POST data fields, then PUT data fields
	log.Printf("view task called")
	//create vars variable to access the url parameters
	vars := mux.Vars(r)

	//get id from url params
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	//query
	rows, err := db.Query(("select id, name from Task where id=%d", taskId)
	if err != nil {
		log.Fatal(err)
	}
	//close rows
	defer rows.Close()

	//todo: make sure only one row is returned

	var id int
	var name string
	//make row cursor point to first and only row
	rows.Next()
	//fill in id and name variables
	rows.Scan(&id, &name)
	//create task model with given id and name
	cur_task := models.Task{Id: id, Name: name}
	//serialize task and then output it to user as json
	if err := json.NewEncoder(w).Encode(cur_task); err != nil {
		panic(err)
	}
}

//return a list of all tasks
func ViewAllTasks(w http.ResponseWriter, r *http.Request) {
	//query for tasks
	rows, err := db.Query("select id, name from Task")
	if err != nil {
		log.Fatal(err)
	}
	//close query
	defer rows.Close()

	//empty task list model
	tasks := []models.Task{}

	//file up task list model with db output
	for rows.Next() {
		var id int
		var name string
		//fill id and name
		rows.Scan(&id, &name)
		//create model Task
		cur_task := models.Task{Id: id, Name: name}
		//add model task to list
		tasks = append(Tasks, cur_task)
	}
	//encode entire list as json and return it to user
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		panic(err)
	}
}

///delete a task
//input: task id in url 
//output: nothing
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	
	log.Printf("view task called")
	//create vars variable to access the url parameters
	vars := mux.Vars(r)

	//get id from url params
	taskId, err := strconv.Atoi(vars["id"])

	//begin transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	//begin statement
	stmt, err := tx.Prepare("delete from Task where Task.id = ?")
	if err != nil {
		log.Fatal(err)
	}
	//close statement
	defer stmt.Close()

	//execute statement with taskid
	_, err = stmt.Exec(taskId)
	if err != nil {
		log.Fatal(err)
	}

	//commit transaction
	tx.Commit()

}

func main() {

	

	//views
	r.HandleFunc("/", Home)

	//API

	//task

	//todo: can name each route in order to reverse them.
	s := r.PathPrefix("/task").Subrouter()
	s.HandleFunc("/", CreateTask).Methods("PUT").Name("CreateTaskUrl")
	s.HandleFunc("/{id:[0-9]+}/", UpdateTask).Methods("POST").Name("UpdateTaskUrl")
	s.HandleFunc("/", ViewAllTasks).Methods("GET").Name("ViewAllTasksUrl")
	s.HandleFunc("/{id:[0-9]+}/", ViewTask).Methods("GET").Name("ViewTaskUrl")
	s.HandleFunc("/{id:[0-9]+}/", DeleteTask).Methods("DELETE").Name("DeleteTaskUrl")

	//static files
	fs := http.FileServer(http.Dir("views/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))

	db.Close()
	apiclient.test()



}
