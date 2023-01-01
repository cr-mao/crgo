package log

import (
	"log"

	"crgo/infra/conf"

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

	sugarLogger := logger.Sugar()
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
