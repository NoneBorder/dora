package dora

import "github.com/astaxie/beego/logs"

var Logger = logs.GetBeeLogger()

func init() {
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(4)
}

// SetLogger sets a new logger.
func SetLogger(adaptername string, config string) error {
	return logs.SetLogger(adaptername, config)
}

// SetLevel sets the global log level used by the simple logger.
func SetLevel(l int) {
	logs.SetLevel(l)
}

func Debug(format string, v ...interface{}) {
	logs.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	logs.Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	logs.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	logs.Error(format, v...)
}

func Fatal(format string, v ...interface{}) {
	logs.Critical(format, v...)
}
