package server

import (
	"net"

	"github.com/HsiaoCz/foo/User/pb/upv1"
	"github.com/HsiaoCz/foo/User/service"
	"google.golang.org/grpc"
)

func ResGrpcServer(network string, addr string) error {
	listen, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	upv1.RegisterUserPartServer(server, service.NewUserCase())
	return server.Serve(listen)
}
