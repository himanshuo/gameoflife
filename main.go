package main
import ( 
	"github.com/himanshuo/gameoflife/models"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"html/template"

)

func BaseHandler(w http.ResponseWriter, r *http.Request){
	task := models.Task{Id:1, Name: "Himanshu Ojha"}
    t, _ := template.ParseFiles("templates/base.html")
    t.Execute(w, task)

}

func main(){
	r := mux.NewRouter()
    r.HandleFunc("/", BaseHandler)
    http.Handle("/", r)
    
	log.Fatal(http.ListenAndServe(":8080", nil))

}