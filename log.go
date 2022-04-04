package log

import (
	"errors"
	"github.com/ml444/glog/config"
)

var logger ILogger
var Config *config.Config

func init() {
	Config = config.NewDefaultConfig()
}

func InitLog() error {
	if logger != nil {
		return errors.New("logger has init")
	}
	l := NewLogger(Config)
	err := l.Init()
	if err != nil {
		return err
	}
	logger = l
	return nil
}

func Debug(args ...interface{}) { logger.Debug(args...) }
func Info(args ...interface{})  { logger.Info(args...) }
func Error(args ...interface{}) { logger.Error(args...) }
func Warn(args ...interface{})  { logger.Warn(args...) }
func Print(args ...interface{}) { logger.Print(args...) }
func Panic(args ...interface{}) { logger.Panic(args...) }
func Fatal(args ...interface{}) { logger.Fatal(args...) }

func Debugf(template string, args ...interface{}) { logger.Debugf(template, args...) }
func Infof(template string, args ...interface{})  { logger.Infof(template, args...) }
func Errorf(template string, args ...interface{}) { logger.Errorf(template, args...) }
func Warnf(template string, args ...interface{})  { logger.Warnf(template, args...) }
func Printf(template string, args ...interface{}) { logger.Printf(template, args...) }
func Panicf(template string, args ...interface{}) { logger.Panicf(template, args...) }
func Fatalf(template string, args ...interface{}) { logger.Fatalf(template, args...) }

func Exit() error {
	if logger != nil {
		logger.Stop()
		return logger.Sync()
	}
	return errors.New("logger not open")
}
