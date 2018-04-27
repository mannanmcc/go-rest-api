package main

import (
	"log"
	"net/http"

	"github.com/mannanmcc/rest-api/handlers"
	"github.com/mannanmcc/rest-api/models"
)

func main() {
	db, err := models.NewDB("test:test@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	env := handlers.Env{Db: db}
	router := NewRouter(env)

	//log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", router))
	log.Fatal(http.ListenAndServe(":8080", router))
}
