package repository

import (
	"errors"

	"github.com/afthaab/job-portal-graphql/models"
	"github.com/rs/zerolog/log"
)

func (r *Repo) ViewCompanyID(cid string) (models.Company, error) {
	var companyDetails models.Company
	result := r.DB.Where("id = ?", cid).First(&companyDetails)
	if result.Error != nil {
		return models.Company{}, errors.New("could not find the company in the records")
	}
	return companyDetails, nil
}

func (r *Repo) ViewAllCompanies() ([]models.Company, error) {
	var companyDetails []models.Company
	result := r.DB.Find(&companyDetails)
	if result.Error != nil {
		return nil, errors.New("could not fetch the data")
	}
	return companyDetails, nil
}

func (r *Repo) CreateCompany(companyDetails models.Company) (models.Company, error) {
	result := r.DB.Create(&companyDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("could not create the record in the table")
	}
	return companyDetails, nil
}
