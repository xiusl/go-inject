package v1

import (
	"github.com/xiusl/inject/app/interface/rpc/v1.0/protocol"
	"github.com/xiusl/inject/app/registry"
	"github.com/xiusl/inject/app/usecase"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	protocol.RegisterUserServiceServer(server, NewUserService(ctn.Resolve("user-usecase").(usecase.UserUseCase)))
}
