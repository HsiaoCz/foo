package main

import (
	"log/slog"

	"github.com/HsiaoCz/foo/etc"
	"github.com/HsiaoCz/foo/logger"
	"github.com/HsiaoCz/foo/router"
	"go.uber.org/zap"
)

// foo i have no idea what i want to do
// so i named foo
// jsut foo

var (
	addr = "127.0.0.1:3001"
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
	zap.L().Info("hello my man the server is running", zap.String("addr", addr))
	if err := router.Router(addr); err != nil {
		slog.Error("register router err:", err)
		return
	}
}
