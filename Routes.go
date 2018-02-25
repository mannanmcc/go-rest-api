package main

import (
	"github.com/gorilla/mux"
	"github.com/mannanmcc/rest-api/handlers"
)

func NewRouter(env handlers.Env) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/add-company", env.AddNewCompany).Methods("POST")
	r.HandleFunc("/update-company", env.UpdateCompany).Methods("POST")
	r.HandleFunc("/search", env.Search).Methods("GET")
	r.HandleFunc("/company/{id:[0-9]+}", env.GetCompany).Methods("GET")

	return r
}
