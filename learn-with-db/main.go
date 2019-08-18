package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"type:varchar(100);unique_index"`
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gobyexample sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})
	defer db.Close()
}
