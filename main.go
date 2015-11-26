package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/himanshuo/gameoflife/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	//_ is in order to import a package solely for its side-effects at initialization.
	//In this case, go-sqlite3's side effects are allowing sqlite3 to be usable as a
	//database for  sql.Open
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {

	var err error
	db, err = sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
		create table Task (id integer not null primary key autoincrement, name text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q", err)
		return
	}

}

//views
func Home(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select id, name from Task")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	Tasks := []models.Task{}

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		cur_task := models.Task{Id: id, Name: name}
		Tasks = append(Tasks, cur_task)
	}

	t, _ := template.ParseFiles("views/static/templates/base.html")
	t.Execute(w, Tasks)

}

//API
func CreateTask(w http.ResponseWriter, r *http.Request) {
	//note: r.FormValue searches for key in POST data fields, then PUT data fields
	//2 types of POST submissions: application/x-www-form-urlencoded AND multipart/form-data.
	// need to understand both. Generally speaking, urlencoded takes up extra space so is for normal post requests. multipart form-data does not increase space usage by a lot so is for uploading files
	//http://stackoverflow.com/a/4073451/4710047

	taskName := r.PostFormValue("name")

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into Task(name) values(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	resp, err := stmt.Exec(taskName)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()

	new_task_id, err := resp.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	newTask := models.Task{Id: int(new_task_id), Name: taskName}

	if err := json.NewEncoder(w).Encode(newTask); err != nil {
		panic(err)
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	//note: r.FormValue searches for key in POST data fields, then PUT data fields
	//2 types of POST submissions: application/x-www-form-urlencoded AND multipart/form-data.
	// need to understand both. Generally speaking, urlencoded takes up extra space so is for normal post requests. multipart form-data does not increase space usage by a lot so is for uploading files
	//http://stackoverflow.com/a/4073451/4710047
	log.Printf("called update task")
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
	stmt, err := tx.Prepare("UPDATE Task SET  name=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
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

func ViewTask(w http.ResponseWriter, r *http.Request) {
	//note: r.FormValue searches for key in GET queries, then POST data fields, then PUT data fields
	log.Printf("view task called")
	vars := mux.Vars(r)

	taskId, err := strconv.Atoi(vars["id"])
	log.Printf(vars["id"])
	log.Printf("%s", taskId)

	if err != nil {
		log.Fatal(err)
	}

	q := fmt.Sprintf("select id, name from Task where id=%d", taskId)

	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	rows.Next()
	var id int
	var name string
	rows.Scan(&id, &name)
	cur_task := models.Task{Id: id, Name: name}
	log.Printf("%v", cur_task)
	if err := json.NewEncoder(w).Encode(cur_task); err != nil {
		panic(err)
	}
}

func ViewAllTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select id, name from Task")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	Tasks := []models.Task{}

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		cur_task := models.Task{Id: id, Name: name}
		Tasks = append(Tasks, cur_task)
	}

	if err := json.NewEncoder(w).Encode(Tasks); err != nil {
		panic(err)
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := r.PostFormValue("id")

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("delete from Task where Task.id = ?")
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

	r := mux.NewRouter()

	//views
	r.HandleFunc("/", Home)

	//API

	//task

	//todo: can name each route in order to reverse them.
	s := r.PathPrefix("/task").Subrouter()
	s.HandleFunc("/", CreateTask).Methods("PUT")
	s.HandleFunc("/{id:[0-9]+}/", UpdateTask).Methods("POST")
	s.HandleFunc("/", ViewAllTasks).Methods("GET")
	s.HandleFunc("/{id:[0-9]+}/", ViewTask).Methods("GET")
	s.HandleFunc("/{id:[0-9]+}/", DeleteTask).Methods("DELETE")

	//static files
	fs := http.FileServer(http.Dir("views/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))

	db.Close()

}
