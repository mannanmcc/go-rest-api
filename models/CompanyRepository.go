package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type CompanyRepositoryInterface interface {
	Create(company Company) (int, error)
	Update(company Company) (Company, error)
	SearchAllCompaniesByName(companyName string) ([]Company, error)
	FindByCompanyName(companyName string) (*Company, error)
	FindByRemoteID(remoteId int) (*Company, error)
	FindByID(int) (*Company, error)
}

type CompanyRepository struct {
	Db *gorm.DB
}

func (repo *CompanyRepository) Create(company Company) (int, error) {
	var existingCompany *Company
	existingCompany, _ = repo.FindByCompanyName(company.Name)

	if existingCompany != nil {
		log.Printf("exitsdsds 1")
		return 0, &CompanyDuplicateError{name: company.Name}
	}

	existingCompany, _ = repo.FindByRemoteID(company.RemoteId)
	if existingCompany != nil {
		log.Printf("exitsdsds 3")
		return 0, &DuplicateCompanyRemoteIdError{remoteId: company.RemoteId}
	}
	log.Printf("saving company")
	repo.Db.Save(&company)

	return company.ID, nil
}

func (repo *CompanyRepository) Update(company Company) (Company, error) {
	var existingCompany *Company
	var err error

	if existingCompany, err = repo.FindByID(company.ID); err != nil {
		return company, errors.New(fmt.Sprintf("Company not found with %d", company.ID))
	}

	//check company name is not belongs to another company
	companyFound, _ := repo.FindByCompanyName(company.Name)
	if companyFound != nil && existingCompany.ID != companyFound.ID {
		return company, &CompanyDuplicateError{name: company.Name}
	}

	id := repo.Db.Save(&company)

	if id == nil {
		return company, errors.New("company saving failed")
	}

	return company, nil
}

func (repo *CompanyRepository) FindByRemoteID(remoteID int) (*Company, error) {
	var company Company
	res := repo.Db.Find(&company, &Company{RemoteId: remoteID})

	if res.RecordNotFound() {
		return nil, errors.New(fmt.Sprintf("company not found with remote id: %d", remoteID))
	}

	return &company, nil
}

func (repo *CompanyRepository) FindByID(id int) (*Company, error) {
	var company Company
	res := repo.Db.Find(&company, &Company{ID: id})

	if res.RecordNotFound() {
		return nil, errors.New(fmt.Sprintf("company not found with remote id: %d", id))
	}

	return &company, nil
}

func (repo *CompanyRepository) FindByCompanyName(name string) (*Company, error) {
	var company Company
	res := repo.Db.Find(&company, &Company{Name: name})

	if res.RecordNotFound() {
		return nil, errors.New(fmt.Sprintf("company not found with name: %s", name))
	}

	return &company, nil
}

func (repo *CompanyRepository) SearchAllCompaniesByName(companyName string) ([]Company, error) {
	var companies []Company
	query := "name LIKE '%" + companyName + "%'"

	repo.Db.Where(query).Find(&companies)

	if len(companies) < 1 {
		return nil, errors.New(fmt.Sprintf("No company name matched with keywords: %s", companyName))
	}

	return companies, nil
}
