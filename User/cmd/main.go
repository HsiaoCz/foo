package main

import (
	"log/slog"

	"github.com/HsiaoCz/foo/User/config"
	"github.com/HsiaoCz/foo/User/logger"
	"github.com/HsiaoCz/foo/User/pkg"
	"github.com/HsiaoCz/foo/User/server"
)

var (
	network   = "tcp"
	addr      = config.Conf.App.Addr
	startTime = config.Conf.App.Stime
	machineID = config.Conf.App.MachineID
)

func main() {
	logger.InitSlogger()
	if err := config.ParseConfig(); err != nil {
		slog.Error("parse config err:", err)
		return
	}
	if err := logger.InitLogger(logger.NewZapLoggerConf()); err != nil {
		slog.Error("init logger err:", err)
		return
	}
	if err := server.ResGrpcServer(network, addr); err != nil {
		slog.Error("register grpc server err:", err)
		return
	}
	if err := pkg.Init(startTime, int64(machineID)); err != nil {
		slog.Error("init identity err:", err)
		return
	}
}
