package repository

import (
	"errors"

	"github.com/afthaab/job-portal-graphql/models"
)

func (r *Repo) CreateUser(userDetails models.User) (models.User, error) {
	result := r.DB.Create(&userDetails)
	if result.Error != nil {
		return models.User{}, errors.New("could not create the record in the table")
	}
	return userDetails, nil
}
