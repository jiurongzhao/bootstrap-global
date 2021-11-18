package log

import "fmt"

type LoggerConfig interface {
	Load() (Logger, error)
}

type Logger interface {
	Debug(msg string, params ...interface{})
	Info(msg string, params ...interface{})
	Warn(msg string, params ...interface{})
	Error(msg string, params ...interface{})
}

var adapters = make(map[string]LoggerConfig)
var instance Logger

func Register(name string, adapter LoggerConfig) {
	_, ok := adapters[name]
	if ok {
		panic(fmt.Errorf("adapter [" + name + "] already existing, don't register again."))
	}
	adapters[name] = adapter
}

func InitGlobalLogger(name string) error {
	adapter, ok := adapters[name]
	if !ok {
		return fmt.Errorf("not found adapter with [" + name + "].")
	}
	if res, err := adapter.Load(); err != nil {
		return err
	} else {
		instance = res
		return nil
	}
}

func Debug(msg string, params ...interface{}) {
	instance.Debug(msg, params...)
}
func Info(msg string, params ...interface{}) {
	instance.Info(msg, params...)
}
func Warn(msg string, params ...interface{}) {
	instance.Warn(msg, params...)
}
func Error(msg string, params ...interface{}) {
	instance.Error(msg, params...)
}
