package main

import (
	"fmt"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, World")
}

func hey(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"hey!")
}

func main(){

	http.HandleFunc("/", handler)
	http.HandleFunc("/hey",hey)
	http.ListenAndServe(":8080",nil)
}
