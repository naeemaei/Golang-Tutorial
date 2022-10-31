package main

import "sync"

var (
	mx     = sync.Mutex{}
	config *Config
)

func GetConfigWithSingleton() *Config {
	if config == nil {
		mx.Lock()
		defer mx.Unlock()
		if config == nil {
			config = &Config{}
			println("Creating config")
		}
	}
	println("Get config from old instance")
	return config
}
