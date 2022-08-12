package main

import "sync"

var once sync.Once

func GetConfigWithOnce() *Config {
	once.Do(func() {
		println("Creating config")
		config = &Config{}
	})
	println("Get config from old instance")
	return config
}
