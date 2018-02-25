package main

import (
	"github.com/gorilla/mux"
	"github.com/mannanmcc/rest-api/handlers"
	"github.com/codegangsta/negroni"
	"net/http"
)

func NewRouter(env handlers.Env) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/get-token", env.GetToken).Methods("POST")
	//r.HandleFunc("/company/add", env.AddNewCompany).Methods("POST")
	r.HandleFunc("/company/update", env.UpdateCompany).Methods("POST")
	r.HandleFunc("/company/search", env.Search).Methods("GET")
	r.HandleFunc("/company/{id:[0-9]+}", env.GetCompany).Methods("GET")

	//wrap a route with negroni to make it secure
	r.Handle("/company/add", negroni.New(
		negroni.HandlerFunc(env.ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(env.AddNewCompany)),
	))

	return r
}
