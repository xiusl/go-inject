package v1

import (
	"context"

	"github.com/xiusl/inject/app/interface/rpc/v1.0/protocol"
	"github.com/xiusl/inject/app/usecase"
)

// userService
// 对 gRpc 处理程序的封装，userService 只依赖于 usecase.UserUseCase
type userService struct {
	userUsecase usecase.UserUseCase
}

// NewUserService
// 工厂方法
func NewUserService(userUsecase usecase.UserUseCase) *userService {
	return &userService{
		userUsecase: userUsecase,
	}
}

// ListUser
// rpc 获取全部用户
func (us *userService) ListUser(context context.Context, in *protocol.ListUserRequestType) (*protocol.ListUserResponseType, error) {
	users, err := us.userUsecase.ListUser()
	if err != nil {
		return nil, err
	}

	res := &protocol.ListUserResponseType{
		Users: convertToUser(users),
	}
	return res, nil
}

// ListUser
// rpc 注册用户
func (us *userService) Register(context context.Context, in *protocol.RegisterRequestType) (*protocol.RegisterResponseType, error) {
	if err := us.userUsecase.RegisterUser(in.GetName()); err != nil {
		return &protocol.RegisterResponseType{}, err
	}
	return &protocol.RegisterResponseType{}, nil
}

func convertToUser(users []*usecase.User) []*protocol.User {
	res := make([]*protocol.User, len(users))
	for i, user := range users {
		res[i] = &protocol.User{
			Id:   user.ID,
			Name: user.Name,
		}
	}
	return res
}
