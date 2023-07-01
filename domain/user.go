package domain

import (
	"context"
	"twitter"
)

type UserService struct {
	UserRepo twitter.UserRepo
}

func NewUserService(ur twitter.UserRepo) *UserService {
	return &UserService{
		UserRepo: ur,
	}
}

func (ur *UserService) GetById(c context.Context, id string) (twitter.User, error) {
	return ur.UserRepo.GetById(c, id)
}
