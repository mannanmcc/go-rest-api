package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(env Env) *mux.Router {
	var routes = Routes{
		Route{"AddNewCompany", "POST", "/add-company", env.AddNewCompany},
		Route{"UpdateCompany", "POST", "/update-company", env.updateCompany},
		Route{"SearchCompany", "GET", "/search", env.search},
		Route{"GetCompany", "GET", "/company/{id:[0-9]+}", env.getCompany},
	}

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}

	return router
}
