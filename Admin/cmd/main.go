package main

import (
	"log/slog"

	"github.com/HsiaoCz/foo/Admin/etc"
	"github.com/HsiaoCz/foo/Admin/logger"
	"github.com/HsiaoCz/foo/Admin/server"
	"go.uber.org/zap"
)

var (
	network = "tcp"
	addr    = "127.0.0.1:3002"
)

func main() {
	if err := etc.ParseConfig(); err != nil {
		slog.Error("parse config err:", err)
		return
	}
	if err := logger.InitLogger(logger.NewZapLoggerConf()); err != nil {
		slog.Error("init logger err:", err)
		return
	}
	if err := server.ResGrpcServer(network, addr); err != nil {
		zap.L().Error("register grpc server err:%v\n", zap.Error(err))
		return
	}
}
