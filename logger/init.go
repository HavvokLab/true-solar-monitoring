package logger

var loggerInstance Logger

func InitLogger(logName string) {
	loggerInstance = NewLogger(&LoggerOption{
		LogLevel:    LOG_LEVEL_INFO,
		LogName:     logName,
		SkipCaller:  1,
		LogSize:     1024,
		LogBackup:   3,
		LogCompress: true,
	})
}

func GetLogger() Logger {
	return loggerInstance
}
