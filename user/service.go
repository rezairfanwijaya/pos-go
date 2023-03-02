package user

import (
	"fmt"
	"pos/helper"
)

type IService interface {
	Login(input InputUserLogin) (User, error)
	GetUserByID(userID int) (User, error)
}

type Service struct {
	userRepo IRepository
}

func NewService(userRepo IRepository) *Service {
	return &Service{userRepo}
}

func (s *Service) Login(input InputUserLogin) (User, error) {
	userByUsername, err := s.userRepo.FindByUsername(input.Username)
	if err != nil {
		return userByUsername, err
	}

	if userByUsername.ID == 0 {
		return userByUsername, fmt.Errorf(
			"username %v not found",
			input.Username,
		)
	}

	if err := helper.VerifyPassword(
		[]byte(input.Password),
		[]byte(userByUsername.Password),
	); err != nil {
		return userByUsername, fmt.Errorf("invalid password")
	}

	return userByUsername, nil
}

func (s *Service) GetUserByID(userID int) (User, error) {
	userByUserID, err := s.userRepo.FindByUserID(userID)
	if err != nil {
		return userByUserID, err
	}

	if userByUserID.ID == 0 {
		return userByUserID, fmt.Errorf(
			"id %v not found",
			userID,
		)
	}

	return userByUserID, nil
}
