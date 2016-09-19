package main //Must heve that name!!


import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello GE\n")
}

func main() {
	//fmt.Println("Hello, World!")
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)



}
