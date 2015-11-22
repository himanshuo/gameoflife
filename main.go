package main
import ( 
	"github.com/himanshuo/gameoflife/models"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"html/template"
	"strconv"

)
var Tasks []models.Task
func init(){
	Tasks = []models.Task{}
}

//views
func Home(w http.ResponseWriter, r *http.Request){
	
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

}