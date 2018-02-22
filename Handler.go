package main

import (
	"github.com/mannanmcc/rest-api/models"
	"net/http"
	"encoding/json"
	"strconv"
	"fmt"
)

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

func (env Env) AddNewCompany(w http.ResponseWriter, r *http.Request) {
	var err error

	remoteId, _ := strconv.Atoi(r.FormValue("remoteId"))
	linkedInId, _ := strconv.Atoi(r.FormValue("linkedInId"))

	company := models.Company{
		Name: r.FormValue("name"),
		RemoteId: remoteId,
		LinkedInId: linkedInId,
		Status: r.FormValue("status"),
		Ticker: r.FormValue("ticker"),
	}

	if _, err := company.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	companyRepo := models.CompanyRepository{Db: env.db}

	if _, err = companyRepo.Create(company); err != nil {
		fmt.Fprintln(w, StatusError{500, err})
	}
}

func (env Env) updateCompany(w http.ResponseWriter, r *http.Request) {
	var err error

	companyRepo := models.CompanyRepository{Db: env.db}
	remoteId, _ := strconv.Atoi(r.FormValue("remoteId"))
	linkedInId, _ := strconv.Atoi(r.FormValue("linkedInId"))
	companyId, _ := strconv.Atoi(r.FormValue("id"))

	company := models.Company{
		ID: companyId,
		RemoteId: remoteId,
		Name: r.FormValue("name"),
		Ticker: r.FormValue("ticker"),
		LinkedInId: linkedInId,
		Status: r.FormValue("status"),
		ApprovalStatus: r.FormValue("approvalStatus"),
	}

	if _, err := company.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if _, err = companyRepo.Update(company); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
}

func (env Env) search(w http.ResponseWriter, r *http.Request) {
	companyRepo := models.CompanyRepository{Db: env.db}
	companies, err := companyRepo.SearchAllCompaniesByName(r.FormValue("q"))

	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(companies)
}