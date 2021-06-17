package service

import (
	"github.com/xiusl/inject/app/domain/repository"
	apperror "github.com/xiusl/inject/app/error"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s UserService) Duplicated(name string) error {
	user, err := s.repo.FindByName(name)
	if user != nil {
		return apperror.NewAppNameDuplicatedErr(name)
	}
	if err != nil {
		return err
	}
	return nil
}
