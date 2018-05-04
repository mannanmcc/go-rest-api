package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mannanmcc/rest-api/models"
)

//Response represent response type
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// AddNewCompany - handle add new company request
func (env Env) AddNewCompany(w http.ResponseWriter, r *http.Request) {
	var err error
	remoteID, _ := strconv.Atoi(r.FormValue("remoteId"))
	linkedInID, _ := strconv.Atoi(r.FormValue("linkedInId"))

	company := models.Company{
		Name:       r.FormValue("name"),
		RemoteId:   remoteID,
		LinkedInId: linkedInID,
		Status:     r.FormValue("status"),
		Ticker:     r.FormValue("ticker"),
	}

	if _, err := company.Validate(); err != nil {
		JSONResponse("FAILED", err.Error(), w)
		return
	}

	companyRepo := models.CompanyRepository{Db: env.Db}

	var companyID int
	if companyID, err = companyRepo.Create(company); err != nil {
		JSONResponse("FAILED", err.Error(), w)
		return
	}

	company.ID = companyID
	indexCompany(company)

	JSONResponse("SUCCESS", "company added successfully", w)
}

//JSONResponse builds up the response object and encode
func JSONResponse(status string, msg string, w http.ResponseWriter) {
	response := Response{
		Status:  status,
		Message: msg,
	}

	json.NewEncoder(w).Encode(response)
}

//UpdateCompany handles request of updating company details
func (env Env) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	var err error

	companyRepo := models.CompanyRepository{Db: env.Db}
	remoteID, _ := strconv.Atoi(r.FormValue("remoteId"))
	linkedInID, _ := strconv.Atoi(r.FormValue("linkedInId"))
	companyID, _ := strconv.Atoi(r.FormValue("id"))

	company := models.Company{
		ID:             companyID,
		RemoteId:       remoteID,
		Name:           r.FormValue("name"),
		Ticker:         r.FormValue("ticker"),
		LinkedInId:     linkedInID,
		Status:         r.FormValue("status"),
		ApprovalStatus: r.FormValue("approvalStatus"),
	}

	if _, err := company.Validate(); err != nil {
		JSONResponse("FAILED", err.Error(), w)
		return
	}

	if _, err = companyRepo.Update(company); err != nil {
		JSONResponse("FAILED", err.Error(), w)
		return
	}

	JSONResponse("SUCCESS", "Company update successful", w)
}

//Search handles company search request
func (env Env) Search(w http.ResponseWriter, r *http.Request) {
	searchKeywords := r.FormValue("q")

	companies, err := searchCompanyByName(searchKeywords)
	//if company not found in the elasticsearch index, then lookup in the database
	if err != nil {
		companyRepo := models.CompanyRepository{Db: env.Db}
		companies, err = companyRepo.SearchAllCompaniesByName(searchKeywords)

		if err != nil {
			JSONResponse("FAILED", err.Error(), w)
			return
		}
	}

	json.NewEncoder(w).Encode(companies)
}

//GetCompany handles finding company by company id
func (env Env) GetCompany(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	companyID, err := strconv.Atoi(params["id"])

	if err != nil {
		JSONResponse("FAILED", "company id not provided", w)
		return
	}

	companyRepo := models.CompanyRepository{Db: env.Db}
	companyFound, err := companyRepo.FindByID(companyID)

	if err != nil {
		JSONResponse("FAILED", "No company found with id provided", w)
		return

	}

	json.NewEncoder(w).Encode(&companyFound)
}
