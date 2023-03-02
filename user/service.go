package user

import (
	"fmt"
	"pos/helper"
)

type IService interface {
	Login(input InputUserLogin) (User, error)
}

type service struct {
	userRepo IRepository
}

func NewService(userRepo IRepository) *service {
	return &service{userRepo}
}

func (s *service) Login(input InputUserLogin) (User, error) {
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
