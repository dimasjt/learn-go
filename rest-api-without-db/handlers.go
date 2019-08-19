package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ArticlesRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/articles")
	json.NewEncoder(w).Encode(Articles)
}

func ArticleRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/articles/:id")
	vars := mux.Vars(r)
	id := vars["id"]

	for _, article := range Articles {
		if idParam, err := strconv.ParseInt(id, 10, 64); err == nil {
			if idParam == article.ID {
				json.NewEncoder(w).Encode(article)
				break
			}
		}
	}
}

func CreateArticleRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/articles/create")
	var article Article
	json.NewDecoder(r.Body).Decode(&article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(&article)
}

func DeleteArticleRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/articles/{id} DELETE")

	vars := mux.Vars(r)
	idParam := vars["id"]
	id, _ := strconv.ParseInt(idParam, 10, 64)

	for index, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			json.NewEncoder(w).Encode(article)
		}
	}
}

func UpdateArticleRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/articles/{id} DELETE")

	vars := mux.Vars(r)
	idParam := vars["id"]
	id, _ := strconv.ParseInt(idParam, 10, 64)

	for index, article := range Articles {
		if article.ID == id {
			var article Article
			if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
				log.Fatal(err)
			}

			Articles[index] = article
			json.NewEncoder(w).Encode(&article)
		}
	}
}
