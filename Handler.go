package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func CompanyList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getAllCompanies())
}

func AddNewCompany(w http.ResponseWriter, r *http.Request) {
	var company Company
	company = Company{Name: r.FormValue("name")}
	AddCompany(company)
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	companyId, err := strconv.ParseInt(r.FormValue("id"), 10, 64)

	if err != nil {
		fmt.Println("Oops, error while parsing posted value to int")
	}

	DeleteCompanyFromDatabase(companyId)
}
