package handlers

import (
	"github.com/mannanmcc/rest-api/models"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
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

	companyRepo := models.CompanyRepository{Db: env.Db}

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

func (env Env) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	var err error

	companyRepo := models.CompanyRepository{Db: env.Db}
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

func (env Env) Search(w http.ResponseWriter, r *http.Request) {
	companyRepo := models.CompanyRepository{Db: env.Db}
	companies, err := companyRepo.SearchAllCompaniesByName(r.FormValue("q"))

	if err != nil {
		JsonResponse("FAILED", err.Error(), w)
		return
	}

	json.NewEncoder(w).Encode(companies)
}

func (env Env) GetCompany(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	companyId, err := strconv.Atoi(params["id"])

	if err != nil {
		JsonResponse("FAILED", "company id not provided", w)
		return
	}

	companyRepo := models.CompanyRepository{Db: env.Db}
	companyFound, err := companyRepo.FindByID(companyId)

	if err != nil {
		JsonResponse("FAILED", "No company found with id provided", w)
		return

	}

	json.NewEncoder(w).Encode(&companyFound)
}