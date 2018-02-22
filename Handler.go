package main

import (
	"github.com/mannanmcc/rest-api/models"
	"net/http"
	"encoding/json"
	"strconv"
)

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
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
		JsonResponse("FAILED", err.Error(), w)
		return
	}

	companyRepo := models.CompanyRepository{Db: env.db}

	if _, err = companyRepo.Create(company); err != nil {
		JsonResponse("FAILED", err.Error(), w)
		return
	}

	JsonResponse("SUCCESS", "New company added", w)
}

func JsonResponse(status string, msg string, w http.ResponseWriter)  {
	response := Response{
		Status: status,
		Message: msg,
	}

	json.NewEncoder(w).Encode(response)
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
		JsonResponse("FAILED", err.Error(), w)
		return
	}

	if _, err = companyRepo.Update(company); err != nil {
		JsonResponse("FAILED", err.Error(), w)
		return
	}

	JsonResponse("SUCCESS", "Company update successful", w)
}

func (env Env) search(w http.ResponseWriter, r *http.Request) {
	companyRepo := models.CompanyRepository{Db: env.db}
	companies, err := companyRepo.SearchAllCompaniesByName(r.FormValue("q"))

	if err != nil {
		JsonResponse("FAILED", err.Error(), w)
		return
	}

	json.NewEncoder(w).Encode(companies)
}