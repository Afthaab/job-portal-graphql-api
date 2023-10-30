package service

import (
	"strconv"

	"github.com/afthaab/job-portal-graphql/graph/model"
	newModel "github.com/afthaab/job-portal-graphql/models"
	"github.com/afthaab/job-portal-graphql/pkg"
)

func (s *Service) UserSignup(userData model.NewUser) (*model.User, error) {
	hashedPassword, err := pkg.HashPassword(userData.Password)
	if err != nil {
		return nil, err
	}

	userDetails := newModel.User{
		Username:     userData.Username,
		Email:        userData.Email,
		HashPassword: hashedPassword,
	}

	userDetails, err = s.userRepo.CreateUser(userDetails)
	if err != nil {
		return nil, err
	}

	uid := strconv.FormatUint(uint64(userDetails.ID), 10)

	return &model.User{
		ID:        uid,
		Username:  userDetails.Username,
		Email:     userDetails.Email,
		CreatedAt: userDetails.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: userDetails.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
