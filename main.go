package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/testground/learning-example/pkg/env"
	"github.com/testground/learning-example/pkg/rabbit"
	"github.com/testground/learning-example/pkg/redis"
)

// Initializes the services in use by the server: blockchain, rabbitMq, redis, etc.
func initialize() {
	env.GetEnv()

	log.Printf("Waiting for redis and rabbitMq connections to initialize")
	redis.InitClient()
	rabbit.GetConnection()

	log.Printf("Server initialization complete")
}

func startServer() {
	serverEnv := env.GetEnv()

	initialize()

	port := serverEnv.Port
	log.Printf("Server started and listening at port %s\n", port)

	listenPort := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(listenPort, nil))
}

func main() {
	startServer()
}
