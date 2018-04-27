package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mannanmcc/rest-api/handlers"
)

//NewRouter points a request to a handler function
func NewRouter(env handlers.Env) *negroni.Negroni {
	r := mux.NewRouter()

	r.HandleFunc("/get-token", env.GetToken).Methods("POST")
	r.HandleFunc("/company/add", env.AddNewCompany).Methods("POST")
	r.HandleFunc("/company/update", env.UpdateCompany).Methods("POST")
	r.HandleFunc("/company/search", env.Search).Methods("GET")
	r.HandleFunc("/company/{id:[0-9]+}", env.GetCompany).Methods("GET")

	//wrap with a second muxer to apply middleware
	routerMux := http.NewServeMux()
	routerMux.Handle("/get-token", r)
	routerMux.Handle("/company/", negroni.New(
		negroni.HandlerFunc(env.ValidateTokenMiddleware),
		negroni.Wrap(r),
	))

	n := negroni.Classic()
	n.UseHandler(routerMux)

	return n
}
