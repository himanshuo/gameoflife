package main
import ( 
	"github.com/himanshuo/gameoflife/models"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"html/template"
	"github.com/gorilla/websocket"

)

//The Conn type represents a WebSocket connection. 
	//todo: is this correct?
	//Conn is an interface and websocket implements it ???
//websocket.Upgrader object uses upgrade function for a given http request handler 
//to get a pointer to a conn.
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func BaseHandler(w http.ResponseWriter, r *http.Request){
	
	conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }

    for {
	    messageType, p, err := conn.ReadMessage()
	    if err != nil {
	        return
	    }
	    if err = conn.WriteMessage(messageType, p); err != nil {
	        return err
	    }
	}

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