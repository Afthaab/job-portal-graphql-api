package service

import (
	"strconv"

	"github.com/afthaab/job-portal-graphql/graph/model"
	newModel "github.com/afthaab/job-portal-graphql/models"
)

func (s *Service) ViewAllCompanies() ([]*model.Company, error) {
	companies, err := s.userRepo.ViewAllCompanies()
	if err != nil {
		return nil, err
	}
	cpDatas := []*model.Company{}

	for _, v := range companies {
		cpData := model.Company{
			ID:       strconv.FormatUint(uint64(v.ID), 10),
			Name:     v.Name,
			Location: v.Location,
			Salary:   v.Salary,
		}
		cpDatas = append(cpDatas, &cpData)
	}
	return cpDatas, nil
}

func (s *Service) CreateCompany(companyData model.NewComapany) (*model.Company, error) {
	companyDetails := newModel.Company{
		Name:     companyData.Name,
		Location: companyData.Location,
		Salary:   companyData.Salary,
	}

	companyDetails, err := s.userRepo.CreateCompany(companyDetails)
	if err != nil {
		return nil, err
	}

	return &model.Company{
		ID:       strconv.FormatUint(uint64(companyDetails.ID), 10),
		Name:     companyData.Name,
		Location: companyData.Location,
		Salary:   companyData.Salary,
	}, nil

}
