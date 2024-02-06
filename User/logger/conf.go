package logger

import "github.com/HsiaoCz/foo/User/config"

type ZapLoggerConf struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	LogLevel   string
	AppMode    string
}

func NewZapLoggerConf() ZapLoggerConf {
	return ZapLoggerConf{
		Filename:   config.Conf.Log.Filename,
		MaxSize:    config.Conf.Log.MaxSize,
		MaxBackups: config.Conf.Log.MaxBackups,
		MaxAge:     config.Conf.Log.MaxAge,
		LogLevel:   config.Conf.Log.Level,
		AppMode:    config.Conf.App.Mode,
	}
}
