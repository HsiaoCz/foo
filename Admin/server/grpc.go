package server

import (
	"net"

	"github.com/HsiaoCz/foo/Admin/pb/pv1"
	"github.com/HsiaoCz/foo/Admin/service"
	"google.golang.org/grpc"
)

func ResGrpcServer(network string, addr string) error {
	conn, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pv1.RegisterFooServer(server, service.NewAdminSer())
	return server.Serve(conn)
}
