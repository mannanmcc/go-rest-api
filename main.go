package main

import (
	"net/http"
	"github.com/mannanmcc/rest-api/models"
	"github.com/mannanmcc/rest-api/handlers"
	"log"
)

func main() {
	db, err := models.NewDB("test:test@tcp(127.0.0.1:3306)/test")

	if err !=nil{
		panic(err)
	}

	defer db.Close()

	env := handlers.Env{Db: db}
	router := NewRouter(env)

	log.Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem",  router))
}
