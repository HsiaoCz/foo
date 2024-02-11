package server

import (
	"net"

	"github.com/HsiaoCz/foo/User/pb/pv1"
	"github.com/HsiaoCz/foo/User/service"
	"google.golang.org/grpc"
)

func ResGrpcServer(network string, addr string) error {
	listen, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pv1.RegisterUserServer(server, service.NewUserCase())
	return server.Serve(listen)
}
