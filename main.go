package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Some Title", Desc: "Some Description", Content: "Some Content"},
	}
	fmt.Println("Endpoint: Articles")
	json.NewEncoder(w).Encode(articles)
}

func PostArticles(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Article has been posted")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/post", PostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main(){
	handleRequests()
}