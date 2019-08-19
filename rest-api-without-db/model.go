package main

type Article struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var Articles []Article = []Article{
	Article{ID: 1, Title: "Learn Ruby on Rails", Description: "Ruby with Rails"},
	Article{ID: 2, Title: "Learn Go language", Description: "Creating web server using GO"},
}
