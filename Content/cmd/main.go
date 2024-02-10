package main

import (
	"log/slog"

	"github.com/HsiaoCz/foo/Admin/server"
	"github.com/HsiaoCz/foo/Content/etc"
	"github.com/HsiaoCz/foo/Content/logger"
)

var (
	addr = etc.Conf.App.AppPort
)

func main() {
	logger.InitSlogger()
	if err := etc.ParseConfig(); err != nil {
		slog.Error("parse config err:", err)
		return
	}
	if err := logger.InitLogger(logger.NewZapLoggerConf()); err != nil {
		slog.Error("logger init zap logger err:", err)
		return
	}
	if err := server.ResGrpcServer("tcp", addr); err != nil {
		slog.Error("register grpc server err:", err)
		return
	}
}
