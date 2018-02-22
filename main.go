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
		panic(err)
	}

	defer db.Close()

	env := Env{db: db}
	router := NewRouter(env)

	log.Fatal(http.ListenAndServe(":8080", router))
}
