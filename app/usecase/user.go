package usecase

import (
	"github.com/google/uuid"
	"github.com/xiusl/inject/app/domain/model"
	"github.com/xiusl/inject/app/domain/repository"
	"github.com/xiusl/inject/app/domain/service"
)

// User
// 为了解决不同层之间解封业务逻辑的问题
type User struct {
	ID   string
	Name string
}

// UserUseCase
// 用户用例层
type UserUseCase interface {
	ListUser() ([]*User, error)
	RegisterUser(name string) error
}

// userUsecase
// UserUseCase 的具体实现
type userUsecase struct {
	repo    repository.UserRepository
	service *service.UserService
}

// NewUserUsecase
// 工厂方法，依赖 repository.UserRepository 和 service *service.UserService
// 初始化时这两个依赖必须被注入
func NewUserUsecase(repo repository.UserRepository, service *service.UserService) *userUsecase {
	return &userUsecase{
		repo:    repo,
		service: service,
	}
}

// ListUser
// 获取所有用户
func (uc *userUsecase) ListUser() ([]*User, error) {
	users, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return converToUsers(users), nil
}

// Register
// 使用 name 注册用户
func (uc *userUsecase) RegisterUser(name string) error {
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	if err = uc.service.Duplicated(name); err != nil {
		return err
	}

	user := model.NewUser(uid.String(), name)

	if err = uc.repo.Save(user); err != nil {
		return err
	}
	return nil
}

func converToUsers(users []*model.User) []*User {
	res := make([]*User, len(users))
	for i, user := range users {
		res[i] = &User{
			ID:   user.Id,
			Name: user.Name,
		}
	}
	return res
}
