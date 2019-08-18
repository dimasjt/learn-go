package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	StaticDir = "/static/"
	Port      = "3000"
)

type Book struct {
	Title     string
	Content   string
	Published bool
}

type BooksPageData struct {
	Books []Book
}

type BookCreateData struct {
	Book
	Success bool
}

func BookPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	// page := vars["page"]

	book := Book{
		Title:     title,
		Content:   "Man without legs",
		Published: true,
	}

	Render(w, "book", book)
}

func BooksPage(w http.ResponseWriter, r *http.Request) {
	books := BooksPageData{
		Books: []Book{
			{Title: "Book 1", Content: "one paragraph", Published: false},
			{Title: "Book 2", Content: "two paragraph", Published: true},
			{Title: "Book 3", Content: "three paragraph", Published: true},
		},
	}

	Render(w, "books", books)
}

func BookCreatePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Render(w, "book_create", nil)
		return
	}

	published, _ := strconv.ParseBool(r.FormValue("published"))

	newBook := BookCreateData{
		Book: Book{
			Title:     r.FormValue("title"),
			Content:   r.FormValue("content"),
			Published: published,
		},
		Success: true,
	}

	Render(w, "book_create", newBook)
}

func Render(w io.Writer, templatetName string, data interface{}) {
	tmpl := template.Must(template.ParseFiles("views/" + templatetName + ".html"))
	tmpl.Execute(w, data)
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/gobyexample?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	// db := ConnectDB()

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/books/{title}/page/{page}", BookPage)
	r.HandleFunc("/books", BooksPage)
	r.HandleFunc("/books/create", BookCreatePage).Methods("GET", "POST")

	r.PathPrefix(StaticDir).
		Handler(http.StripPrefix(StaticDir, http.FileServer(http.Dir("."+StaticDir))))

	fmt.Println("server started")
	http.ListenAndServe(":"+Port, r)
}
