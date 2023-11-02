package service

import (
	"strconv"

	"github.com/afthaab/job-portal-graphql/graph/model"
	"github.com/afthaab/job-portal-graphql/models"
)

func (s *Service) ViewJobById(jid string) (*model.Job, error) {
	var jobData models.Jobs
	jobData, err := s.userRepo.ViewJobByJobId(jid)
	if err != nil {
		return nil, err
	}
	return &model.Job{
		ID:   int(jobData.ID),
		Name: jobData.Name,
		Type: jobData.Type,
		Cid:  strconv.FormatUint(uint64(jobData.Cid), 10),
		Company: &model.Company{
			ID:       strconv.FormatUint(uint64(jobData.Company.ID), 10),
			Name:     jobData.Company.Name,
			Location: jobData.Company.Location,
			Salary:   jobData.Company.Salary,
		},
	}, nil
}

func (s *Service) CreateJob(input model.NewJob) (*model.Job, error) {
	cid, err := strconv.ParseUint(input.Cid, 10, 64)
	if err != nil {
		return nil, err
	}
	jobData := models.Jobs{
		Name: input.Name,
		Cid:  uint(cid),
		Type: input.Type,
	}
	jobData, err = s.userRepo.CreateJob(jobData)
	if err != nil {
		return nil, err
	}
	return &model.Job{
		ID:   int(jobData.ID),
		Name: jobData.Name,
		Type: jobData.Type,
		Cid:  strconv.FormatUint(uint64(jobData.Cid), 10),
		Company: &model.Company{
			ID:       strconv.FormatUint(uint64(jobData.Company.ID), 10),
			Name:     jobData.Company.Name,
			Location: jobData.Company.Location,
			Salary:   jobData.Company.Salary,
		},
	}, nil
}
