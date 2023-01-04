package log

import (
	"crgo/infra/conf"
	"log"

	"go.uber.org/zap"
)

var Debug func(args ...interface{})

var Info func(args ...interface{})

var Warn func(args ...interface{})

var Error func(args ...interface{})

var Debugf func(format string, args ...interface{})

var Infof func(format string, args ...interface{})

var Warnf func(format string, args ...interface{})

var Errorf func(format string, args ...interface{})
var Fatalf func(format string, args ...interface{})

func InitLogger() {
	var logger *zap.Logger
	var err error
	if conf.IsProd() {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}
	log.Println("Log init")

	//zap.ReplaceGlobals(logger)
	// Zap 提供了两种类型的日志，分别是Logger 与 SugaredLogger。其中，Logger是默认的，每个要写入的字段都得指定对应类型的方法，这种方式可以不使用反射，因此效率更高。为了避免在异常情况下丢失日志（尤其是在崩溃时），logger.Sync()会在进程退出之前落盘所有位于缓冲区中的日志条目。

	// 而SugaredLogger的性能稍微低于Logger，但是它提供了一种更灵活的打印方式：
	sugarLogger := logger.Sugar() //去掉这行性能更高， 但是要指定类型了。
	Debug = sugarLogger.Debug
	Debugf = sugarLogger.Debugf
	Info = sugarLogger.Info
	Infof = sugarLogger.Infof
	Warn = sugarLogger.Warn
	Warnf = sugarLogger.Warnf
	Error = sugarLogger.Error
	Errorf = sugarLogger.Errorf
	Fatalf = sugarLogger.Fatalf
}
