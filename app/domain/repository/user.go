package repository

import (
	"sync"

	"github.com/xiusl/inject/app/domain/model"
	apperror "github.com/xiusl/inject/app/error"
)

// UserRepository
// 用户仓库层
type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindById(id string) (*model.User, error)
	FindByName(name string) (*model.User, error)
	Save(u *model.User) error
}

// User
// 为了解决不同层之间解封业务逻辑的问题
type User struct {
	Id   string
	Name string
}

// userMemoryRepository
// UserRepository 基于内存的具体实现，
type userMemoryRepository struct {
	mu    *sync.Mutex
	users map[string]*User
}

// NewUserMemoryRepository
// 工厂方法，初始化 Mutex 和 map
func NewUserMemoryRepository() *userMemoryRepository {
	return &userMemoryRepository{
		mu:    &sync.Mutex{},
		users: make(map[string]*User),
	}
}

// FindAll
// 返回所有用户
func (mr *userMemoryRepository) FindAll() ([]*model.User, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	var users []*model.User
	for _, user := range mr.users {
		users = append(users, model.NewUser(user.Id, user.Name))
	}
	return users, nil
}

// FindById
// 根据 id 查询用户
func (mr *userMemoryRepository) FindById(id string) (*model.User, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	for _, user := range mr.users {
		if user.Id == id {
			return model.NewUser(user.Id, user.Name), nil
		}
	}
	return nil, apperror.NewAppNotFoundErr(id)
}

// FindByName
// 根据 name 查询用户
func (mr *userMemoryRepository) FindByName(name string) (*model.User, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	for _, user := range mr.users {
		if user.Name == name {
			return model.NewUser(user.Id, user.Name), nil
		}
	}
	return nil, apperror.NewAppNotFoundErr(name)
}

// Save
// 保存用户
func (mr *userMemoryRepository) Save(u *model.User) error {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	mr.users[u.Id] = &User{
		Id:   u.Id,
		Name: u.Name,
	}
	return nil
}
