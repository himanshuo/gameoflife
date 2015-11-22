package main
import ( 
	"github.com/himanshuo/gameoflife/models"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"html/template"


)
var Tasks []models.Task

//views
func Home(w http.ResponseWriter, r *http.Request){
	
	t, _ := template.ParseFiles("templates/base.html")
    t.Execute(w, Tasks)
	

}


//API
func CreateTask(w http.ResponseWriter, r *http.Request){
	taskName := r.PostFormValue("name")
	newTask := models.Task{Id: len(Tasks), Name: taskName}


    if err := json.NewEncoder(w).Encode(newTask); err != nil {
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
    // s.HandleFunc("/", ViewAllTasks).Methods("GET")
    // s.HandleFunc("/{id:[0-9]+}/", ViewTask).Methods("GET")
    // s.HandleFunc("/{id:[0-9]+}/", DeleteTask).Methods("DELETE")
    

   
    

    http.Handle("/", r)
    
	log.Fatal(http.ListenAndServe(":8080", nil))

}