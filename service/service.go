package service

import (
	"errors"

	"github.com/afthaab/job-portal-graphql/graph/model"
	"github.com/afthaab/job-portal-graphql/repository"
)

type UserService interface {
	UserSignup(model.NewUser) (*model.User, error)
	CreateCompany(model.NewComapany) (*model.Company, error)
	ViewAllCompanies() ([]*model.Company, error)
}

type Service struct {
	userRepo repository.UserRepo
}

func NewService(userRepo repository.UserRepo) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Service{
		userRepo: userRepo,
	}, nil
}
