package server

import (
	"net"

	"github.com/HsiaoCz/foo/Comment/data"
	"github.com/HsiaoCz/foo/Comment/pb/pv1"
	"github.com/HsiaoCz/foo/Comment/service"
	"google.golang.org/grpc"
)

func RegisterGrpc(network, addr string) error {
	conn, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pv1.RegisterFooServer(server, service.NewCommService(data.NewCommUseCase()))
	return server.Serve(conn)
}
