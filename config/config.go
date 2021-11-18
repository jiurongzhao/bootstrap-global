package config

import (
	"fmt"
)

type Provider interface {
	Load(name string) (Configer, error)
}

type Configer interface {
	Get(key string) (interface{}, bool)
	Resolve(prefix string, p interface{}) error
}

var configs = make(map[string]Provider)
var instance Configer

func Register(name string, provider Provider) {
	if _, ok := configs[name]; ok {
		return
	}
	configs[name] = provider
}

func InitGlobalInstance(name string, filename string) error {
	config, ok := configs[name]
	if !ok {
		return fmt.Errorf("not found adapter: [" + name + "]\n")
	}

	var err error
	instance, err = config.Load(filename)
	return err
}

func Get(key string) (interface{}, bool) {
	return instance.Get(key)
}

func Resolve(prefix string, p interface{}) error {
	return instance.Resolve(prefix, p)
}
