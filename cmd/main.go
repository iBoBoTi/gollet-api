package main

import (
	"fmt"
	"github.com/iBoBoTi/gollet-api/infrastructure/config"
	"github.com/iBoBoTi/gollet-api/infrastructure/server"
	"github.com/joho/godotenv"
	"log"
	"strconv"
)

func main() {
	fmt.Println("Welcome Gollet")

	// Load env vars
	conf := config.NewConfig()
	port, err := strconv.Atoi(conf.ServerPort)
	if err != nil {
		log.Fatalf("invalid port: %v", err)
	}

	env := conf.GinMode
	if env != "release" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("couldn't load env vars: %v", err)
		}
	}

	s, err := server.NewWebServerFactory(server.InstanceGin, port, conf.CtxTimeout)
	if err != nil {
		log.Fatalf("couldn't create server: %v", err)
	}

	s.Start()
}
