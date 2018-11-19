package base

import (
	"webcrawler/helper/log/field"
)

//Option 日志记录器的选项
type Option interface {
	Name() string
}

//记录日志时是否带有调用方的代码位置
type OptWithLocation struct {
	Value bool
}

func (opt OptWithLocation) Name() string {
	return "with location"
}

type MyLogger interface {
	Name() string
	Level() LogLevel
	Format() LogFormat
	Options() []Option

	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Debugln(v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Infoln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Warnln(v ...interface{})

	//会增加记录的额外字段
	WithFields(fields ...field.Field) MyLogger
}
