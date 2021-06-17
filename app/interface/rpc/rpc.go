package rpc

import (
	"github.com/xiusl/inject/app/interface/rpc/v1.0"
	"github.com/xiusl/inject/app/registry"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	v1.Apply(server, ctn)
}
