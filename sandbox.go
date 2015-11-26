package main
import (
	"fmt"
	"net/http"
	"strings"
)
func main(){
	body := strings.NewReader("name=afadf")
	req, err := http.NewRequest("PUT","http://localhost:8080/task/",  
		body)
	if err != nil{
		fmt.Println("req preperation incorrect")
		fmt.Println(err)

		return
	}
	req.Header.Add("Content-Type","application/x-www-form-urlencoded")
	resp, err:= http.DefaultClient.Do(req)
	if err != nil{
		fmt.Println("could not make request")
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}