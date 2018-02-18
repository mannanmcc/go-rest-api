package main

import (
	"github.com/mannanmcc/rest-api/models"
	"net/http"
	"strconv"
	"fmt"
)

type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

func (env Env) AddNewCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	var linkedInId int
	if r.FormValue("remoteId") == "" {
		fmt.Fprintln(w, "remoteId is missing!")
		return
	}

	remoteId, err:= strconv.Atoi(r.FormValue("remoteId"))
	if err != nil {
		fmt.Fprintln(w, "invalid remote id provided!")
		return
	}

	linkedInId, err= strconv.Atoi(r.FormValue("linkedInId"))
	if err != nil {
		fmt.Fprintln(w, "invalid linkedInId provided!")
		return
	}

	company = models.Company{
		Name: r.FormValue("name"),
		RemoteId: remoteId,
		Ticker: r.FormValue("ticker"),
		LinkedInId: linkedInId,
		Status: "NEW",
		ApprovalStatus:string("PENDING"),
	}

	companyRepo := models.CompanyRepository{Db: env.db}
	_, err = companyRepo.Create(company)

	if err != nil {
		fmt.Fprintln(w, StatusError{500, err})
	}
}

func (env Env) updateCompany(w http.ResponseWriter, r *http.Request) {
	var linkedInId, remoteId, companyId int
	var err error

	companyRepo := models.CompanyRepository{Db: env.db}

	if r.FormValue("id") == "" {
		fmt.Fprintln(w, "id is missing!")
		return
	}


	if companyId, err = strconv.Atoi(r.FormValue("id")); err != nil {
		fmt.Fprintln(w, "id is missing")
		return
	}

	if remoteId, err = strconv.Atoi(r.FormValue("remoteId")); err != nil {
		fmt.Fprintln(w, "remoteId is missing")
		return
	}

	if linkedInId, err = strconv.Atoi(r.FormValue("linkedInId")); err != nil {
		fmt.Fprintln(w, "linkedInId is missing")
	}

	company := models.Company{
		ID: companyId,
		RemoteId: remoteId,
		Name: r.FormValue("name"),
		Ticker: r.FormValue("ticker"),
		LinkedInId: linkedInId,
		Status: r.FormValue("Status"),
		ApprovalStatus: r.FormValue("approvalStatus"),
	}

	if _, err = companyRepo.Update(company); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	fmt.Fprintln(w, "company update successful")
}
