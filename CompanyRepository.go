package main

import (
"fmt"
"database/sql"
_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "test"
	dbPass := "test"
	dbName := "test"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func getAllCompanies() []Company {
	db := dbConn()

	rows, err := db.Query("SELECT `id`, `name`, `status` FROM `test`.`company`")
	if err != nil {
		panic(err.Error())
	}

	var companies []Company
	var id int
	var name, status string
	for rows.Next() {
		err := rows.Scan(&id, &name, &status)
		if err != nil { /* error handling */}
		companies = append(companies, Company{ID: id, Status: status, Name: name})
	}

	return companies
}

func AddCompany(company Company){
	db := dbConn()

	statement, err :=db.Prepare("INSERT INTO company (name) VALUES (?)")
	if err != nil {
		fmt.Print(err)
	}

	if _, err := statement.Exec(company.Name); err != nil {
		fmt.Print(err)
	}
}

func DeleteCompanyFromDatabase(companyId int64){
	db := dbConn()

	statement, err := db.Prepare("DELETE FROM company WHERE id = ?")

	if err != nil {
		fmt.Println ("Oops, unable to delete the company")
	}

	statement.Exec(companyId)
}
