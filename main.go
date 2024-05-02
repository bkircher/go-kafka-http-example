package main

import (
	"kafka-http-example/config"
	"kafka-http-example/service"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	service.Run(cfg)
}
