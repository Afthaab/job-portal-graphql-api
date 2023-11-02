package repository

import (
	"errors"

	"github.com/afthaab/job-portal-graphql/models"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

type UserRepo interface {
	CreateUser(userData models.User) (models.User, error)
	CreateCompany(companyDetails models.Company) (models.Company, error)
	ViewAllCompanies() ([]models.Company, error)
	ViewCompanyID(cid string) (models.Company, error)
	CreateJob(input models.Jobs) (models.Jobs, error)
	ViewJobByJobId(jid string) (models.Jobs, error)
}

func NewRepository(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be null")
	}
	return &Repo{
		DB: db,
	}, nil
}
