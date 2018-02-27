package models

import (
    "fmt"
	"strings"
	"errors"
)

type Company struct {
    ID int `gorm:"primary_key";"AUTO_INCREMENT"`
    RemoteId int `gorm:"column:remoteId"`
    Name string
    Ticker string
    LinkedInId int `gorm:"column:linkedInId"`
    Status string
    ApprovalStatus string `gorm:"column:approvalStatus"`
}

func (company *Company) Validate() (bool, error) {
	if strings.TrimSpace(company.Name) == "" {
		return false, errors.New("name field missing")
	}

	if strings.TrimSpace(company.Ticker) == "" {
		return false, errors.New("ticker field missing")
	}

	if strings.TrimSpace(company.Status) == "" {
		return false, errors.New("status field missing")
	}

	if company.RemoteId == 0 {
		return false, errors.New("remoteId field missing")
	}

	return true, nil
}

func (company *Company) TableName() string {
    return "company"
}

type CompanyDuplicateError struct {
    name string
}

func (c *CompanyDuplicateError) Error() string {
    return fmt.Sprintf("another company with name: '%s' already exists", c.name)
}

type DuplicateCompanyRemoteIdError struct {
    remoteId int
}

func (c *DuplicateCompanyRemoteIdError) Error() string {
    return fmt.Sprintf("another company with remote id: '%d' already exists", c.remoteId)
}
