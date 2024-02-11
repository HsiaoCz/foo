package server

import (
	"net"

	"github.com/HsiaoCz/foo/Content/pb/pv1"
	"github.com/HsiaoCz/foo/Content/service"
	"google.golang.org/grpc"
)

func RegisterGrpc(network, addr string) error {
	conn, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pv1.RegisterContentServer(server, service.NewContentService())
	return server.Serve(conn)
}
