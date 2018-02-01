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

func NewRouter() *mux.Router {
	var routes = Routes{
		Route{"Index", "GET", "/", Index},
		Route{"CompanyList", "GET", "/companies", CompanyList},
		Route{"AddNewCompany", "POST", "/add-company", AddNewCompany},
		Route{"DeleteCompany", "POST", "/delete-company", DeleteCompany},
	}

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}

	return router
}
