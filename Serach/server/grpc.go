package server

import (
	"net"

	"github.com/HsiaoCz/foo/Serach/pb/pv1"
	"github.com/HsiaoCz/foo/Serach/service"
	"google.golang.org/grpc"
)

func RegisterGrpc(network, addr string) error {
	conn, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pv1.RegisterFooServer(server, service.NewSerachService())
	return server.Serve(conn)
}
