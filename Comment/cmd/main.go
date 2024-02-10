package main

import (
	"log/slog"

	"github.com/HsiaoCz/foo/Comment/etc"
	"github.com/HsiaoCz/foo/Comment/logger"
	"github.com/HsiaoCz/foo/Comment/server"
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
		slog.Error("init zap logger err:", err)
		return
	}
	if err := server.RegisterGrpc("tcp", addr); err != nil {
		slog.Error("register grpc err:", err)
		return
	}
}
