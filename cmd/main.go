package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/xiusl/inject/app/interface/rpc"
	"github.com/xiusl/inject/app/registry"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("faild to listen: %v.\n", err)
	}

	ctn, err := registry.NewContainer()
	if err != nil {
		log.Fatalf("failed to build container: %v.\n", err)
	}

	server := grpc.NewServer()

	rpc.Apply(server, ctn)

	go func() {
		log.Printf("start grpc server port: 8080.\n")
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Printf("stopping grpc server...\n")
	server.GracefulStop()
	ctn.Clean()
}
