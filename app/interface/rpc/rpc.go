package rpc

import (
	"github.com/xiusl/inject/app/interface/rpc/v1.0"
	"github.com/xiusl/inject/app/register"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *register.Container) {
	v1.Apply(server, ctn)
}
