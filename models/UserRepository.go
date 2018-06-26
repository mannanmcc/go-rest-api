package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//UserRepositoryInterface - user repository interface
type UserRepositoryInterface interface {
	isUserExists(username, password string) (*User, error)
}

//UserRepository - UserRepository type
type UserRepository struct {
	Db *gorm.DB
}

//FindByUserNameAndPassword - return user match with username and password passed in
func (repo *UserRepository) FindByUserNameAndPassword(username, password string) (*User, error) {
	var user User
	res := repo.Db.Find(&user, &User{Username: username, Password: password})

	if res.RecordNotFound() {
		return nil, fmt.Errorf("user not found with username: %s and password: %s", username, password)
	}

	return &user, nil
}
