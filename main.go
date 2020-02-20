package main

import (
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article


func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint")
}

func handleRequests() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main(){
	handleRequests()
}