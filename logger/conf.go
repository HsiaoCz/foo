package logger

type ZapLoggerConf struct {
	Filename  string
	MaxSize   int
	MaxBackup int
	MaxAge    int
	LogLevel  string
	AppMode   string
}
