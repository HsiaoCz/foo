package main

import (
	"log/slog"

	"github.com/HsiaoCz/foo/Serach/etc"
	"github.com/HsiaoCz/foo/Serach/logger"
	"github.com/HsiaoCz/foo/Serach/server"
	"go.uber.org/zap"
)

var (
	network = "tcp"
	addr    = etc.Conf.App.AppPort
)

func main() {
	logger.InitSlogger()
	if err := etc.ParseConfig(); err != nil {
		slog.Error("parse config err:", err)
		return
	}
	if err := logger.InitLogger(logger.NewZapLoggerConf()); err != nil {
		slog.Error("init logger err:", err)
		return
	}
	if err := server.RegisterGrpc(network, addr); err != nil {
		zap.L().Error("register grpc server err:%v\n", zap.Error(err))
		return
	}
}
