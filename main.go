package main

import (
	"log"
	"net/http"
	"github.com/mannanmcc/rest-api/models"
	"github.com/jinzhu/gorm"
)

type Env struct {
	db *gorm.DB
}

func main() {
	db, err := models.NewDB("test:test@tcp(127.0.0.1:3306)/test")

	if err !=nil{
		log.Println("")
	}

	defer db.Close()

	env := Env{db: db}

	router := NewRouter(env)

	if err != nil {
		log.Printf("could not create new company: %d", err.Error())
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}
