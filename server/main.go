package main

import (
	"fmt"
	"net/http"
)

type View struct {
	MyName string
}

func FuncView(w http.ResponseWriter, req *http.Request) {
	fmt.Println("inside funcview")
	//view:=View{MyName:"Partha"}
	//response,_:= json.Marshal(view)
	// w.Header().Set("Content-Type", "text/plain")
	// Set CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	fmt.Fprintf(w, "Hi")
}

func main() {
	fmt.Println("inside main")
	http.HandleFunc("/view", FuncView)
	http.ListenAndServe(":8084", nil)
}
