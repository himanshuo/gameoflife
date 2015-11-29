package main

//_ is in order to import a package solely for its side-effects at initialization.
//In this case, go-sqlite3's side effects are allowing sqlite3 to be usable as a
//database for  sql.Open
import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/himanshuo/gameoflife/models"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var db *sql.DB
var Router *mux.Router

const (
	DB_TYPE = "sqlite3"
	DB_DIR  = "./data/data.db"
)
const (
	BASE_TEMPLATE = "views/static/templates/base.html"
)

//start database and create a url router
func init() {
	if err := startDB(); err!=nil{
		log.Fatal(err)
	} else {
		Router = mux.NewRouter()
	}

}

// creates database if not already created. Creates DB tables if not already created.
//set global db variable
func startDB() error {
	var err error
	db, err = sql.Open(DB_TYPE, DB_DIR)
	if err != nil {
		return err
	}
	//make sure we can actually query the database.
	if err := db.Ping(); err != nil {
		return err
	}
	//make sure table actually exists in db
	tableName := "Task"
	checkTableQuery := "SELECT name FROM sqlite_master WHERE type='table' AND name=?;"
	var output string
	err = db.QueryRow(checkTableQuery, tableName).Scan(&output)
    	switch {
    		case err == sql.ErrNoRows:
	            		log.Printf("Table %s does not exist. Creating it.", tableName)
	            		sqlStmt := "CREATE TABLE Task (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, name TEXT);"
			//run statement to create table. return object of Result type is not needed.
			_, err = db.Exec(sqlStmt)
			if err != nil {
				return err
			}
    		case err != nil:
            			return err
    	}
	log.Printf("Database Started")
	return nil
}

//Views

// Home page view. For now, simply lists all the tasks one by one.
func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("Home Screen Opened")
	//get all tasks
	rows, err := db.Query("SELECT id, name FROM Task")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	tasks := []models.Task{}
	//iterate through response and put into models.Task slice
	for rows.Next() {
		var id int
		var name string
		if err = rows.Scan(&id, &name); err != nil{
			log.Fatal(err)
		}
		curTask := models.Task{Id: id, Name: name}
		tasks = append(tasks, curTask)
	}
	t, err := template.ParseFiles(BASE_TEMPLATE)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, tasks)
}

//API

//create a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("Create Task")
	//note: r.FormValue searches for key in POST data fields, then PUT data fields
	//2 types of POST submissions: application/x-www-form-urlencoded AND multipart/form-data.
	// need to understand both. Generally speaking, urlencoded takes up extra space so is for normal post requests. multipart form-data does not increase space usage by a lot so is for uploading files
	//http://stackoverflow.com/a/4073451/4710047

	//expecting to come from PUT data field
	//name as part of a x-www-form-urlencoded PUT request
	taskName := r.PostFormValue("name")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO Task(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	resp, err := stmt.Exec(taskName)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	newTaskId, err := resp.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	newTask := models.Task{Id: int(newTaskId), Name: taskName}
	if err := json.NewEncoder(w).Encode(newTask); err != nil {
		panic(err)
	}
}

//update a task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("Update Task")
	//create a variable that has the parameters sent to the api in the url
	// /task/<id>/<x>/  will lead to variables id and x in vars
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	newName := r.PostFormValue("name")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("UPDATE Task SET name=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	// output is unneccessary
	_, err = stmt.Exec(newName, taskId)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	updatedTask := models.Task{Id: int(taskId), Name: newName}
	if err := json.NewEncoder(w).Encode(updatedTask); err != nil {
		panic(err)
	}
}

//view a single task
func ViewTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("View Task")
	//create vars variable to access the url parameters
	vars := mux.Vars(r)
	//id of task from url parameter
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	row := db.QueryRow("SELECT id, name FROM Task WHERE id=?", taskId)
	var id int
	var name string
	if err = row.Scan(&id, &name); err!=nil{
		log.Fatal(err)
	}
	curTask := models.Task{Id: id, Name: name}
	if err := json.NewEncoder(w).Encode(curTask); err != nil {
		panic(err)
	}
}

//list all tasks
func ViewAllTasks(w http.ResponseWriter, r *http.Request) {
	log.Printf("View All Tasks")
	rows, err := db.Query("SELECT id, name FROM Task")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	tasks := []models.Task{}
	for rows.Next() {
		var id int
		var name string
		if err = rows.Scan(&id, &name); err!=nil{
			log.Fatal(err)
		}
		curTask := models.Task{Id: id, Name: name}
		tasks = append(tasks, curTask)
	}
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		panic(err)
	}
}

///delete a task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("Delete Task")
	vars := mux.Vars(r)
	//task id in url
	taskId, err := strconv.Atoi(vars["id"])
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("DELETE FROM Task WHERE Task.id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(taskId)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func main() {
	//Views
	Router.HandleFunc("/", Home)

	//API
	//task
	s := Router.PathPrefix("/task").Subrouter()
	s.HandleFunc("/", CreateTask).Methods("PUT").Name("CreateTaskUrl")
	s.HandleFunc("/{id:[0-9]+}/", UpdateTask).Methods("POST").Name("UpdateTaskUrl")
	s.HandleFunc("/", ViewAllTasks).Methods("GET").Name("ViewAllTasksUrl")
	s.HandleFunc("/{id:[0-9]+}/", ViewTask).Methods("GET").Name("ViewTaskUrl")
	s.HandleFunc("/{id:[0-9]+}/", DeleteTask).Methods("DELETE").Name("DeleteTaskUrl")
	//static files
	fs := http.FileServer(http.Dir("views/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", Router)
	log.Printf("Server Started")
	log.Fatal(http.ListenAndServe(":8080", nil))
	db.Close()
}
