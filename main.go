package main
import ( 
	"github.com/himanshuo/gameoflife/models"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"html/template"
	"strconv"
	//_ is in order to import a package solely for its side-effects at initialization. 
	//In this case, go-sqlite3's side effects are allowing sqlite3 to be usable as a
	//database for  sql.Open
	_ "github.com/mattn/go-sqlite3"  
	"database/sql"

)

var db *sql.DB

func init(){
	

	var err error
	db, err = sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		log.Fatal(err)
	}
	
	sqlStmt := `
		create table Task (id integer not null primary key, name text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	

	
}

//views
func Home(w http.ResponseWriter, r *http.Request){
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
		cur_task := models.Task{Id:id, Name:name}
		Tasks = append(Tasks, cur_task)
	}
	

	t, _ := template.ParseFiles("views/static/templates/base.html")
    t.Execute(w, Tasks)



	

}


//API
func CreateTask(w http.ResponseWriter, r *http.Request){
	//note: r.FormValue searches for key in POST data fields, then PUT data fields
	taskName := r.PostFormValue("name")
	newTask := models.Task{Id: len(Tasks), Name: taskName}

	Tasks = append(Tasks, newTask)

    if err := json.NewEncoder(w).Encode(newTask); err != nil {
        panic(err)
    }
}

func ViewTask(w http.ResponseWriter, r *http.Request){
	//note: r.FormValue searches for key in GET queries, then POST data fields, then PUT data fields
    taskId,_ := strconv.Atoi(r.FormValue("id"))
    if err := json.NewEncoder(w).Encode(Tasks[taskId]); err != nil {
        panic(err)
    }
}
func ViewAllTasks(w http.ResponseWriter, r *http.Request){
    if err := json.NewEncoder(w).Encode(Tasks); err != nil {
        panic(err)
    }
}



func main(){

	
	



	r := mux.NewRouter()



	//views
    r.HandleFunc("/", Home)
    

    //API

    //task

    //todo: can name each route in order to reverse them. 
    s := r.PathPrefix("/task").Subrouter()
    s.HandleFunc("/", CreateTask).Methods("PUT")
    // s.HandleFunc("/{id:[0-9]+}/", UpdateTask).Methods("POST")
    s.HandleFunc("/", ViewAllTasks).Methods("GET")
    s.HandleFunc("/{id:[0-9]+}/", ViewTask).Methods("GET")
    // s.HandleFunc("/{id:[0-9]+}/", DeleteTask).Methods("DELETE")
    

   
    //static files
    fs := http.FileServer(http.Dir("views/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.Handle("/", r)
    
	log.Fatal(http.ListenAndServe(":8080", nil))


	db.close()
}