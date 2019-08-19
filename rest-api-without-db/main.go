package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/articles", ArticlesRoute).Methods("GET")
	r.HandleFunc("/articles", CreateArticleRoute).Methods("POST")
	r.HandleFunc("/articles/{id}", ArticleRoute).Methods("GET")
	r.HandleFunc("/articles/{id}", DeleteArticleRoute).Methods("DELETE")
	r.HandleFunc("/articles/{id}", UpdateArticleRoute).Methods("PATCH", "PUT")

	fmt.Println("server started")
	http.ListenAndServe(":3000", r)
}
