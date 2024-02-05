package logger

import "github.com/HsiaoCz/foo/etc"

type ZapLoggerConf struct {
	Filename  string
	MaxSize   int
	MaxBackup int
	MaxAge    int
	LogLevel  string
	AppMode   string
}

func NewZapLoggerConf() ZapLoggerConf {
	return ZapLoggerConf{
		Filename:  etc.Conf.Log.Filename,
		MaxSize:   etc.Conf.Log.MaxSize,
		MaxBackup: etc.Conf.Log.MaxBackups,
		MaxAge:    etc.Conf.Log.MaxAge,
		LogLevel:  etc.Conf.Log.Level,
		AppMode:   etc.Conf.App.Mode,
	}
}
