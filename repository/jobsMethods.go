package repository

import (
	"errors"

	"github.com/afthaab/job-portal-graphql/models"
)

func (r *Repo) CreateJob(jobData models.Jobs) (models.Jobs, error) {
	result := r.DB.Create(&jobData)
	if result.Error != nil {
		return models.Jobs{}, errors.New("could not create the recode")
	}
	if err := r.DB.Preload("Company").Where("id = ?", jobData.ID).First(&jobData).Error; err != nil {
		// Handle the error
		return models.Jobs{}, err
	}
	return jobData, nil
}

func (r *Repo) ViewJobByJobId(jid string) (models.Jobs, error) {
	var jobData models.Jobs
	if err := r.DB.Preload("Company").Where("id = ?", jid).First(&jobData).Error; err != nil {
		// Handle the error
		return models.Jobs{}, err
	}
	return jobData, nil
}
