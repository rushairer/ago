package main

import (
	"fmt"
	"log"

	"PACKAGENAME/api/bootstrap"
	"PACKAGENAME/core/config"

	"github.com/gin-gonic/gin"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("ago/api crashed, error:", err)
		}
	}()

	log.Println("starting...")
	server := gin.Default()
	bootstrap.SetupServer(server)

	log.Println("running...")
	if gin.IsDebugging() {
		err := server.RunTLS(":443", "./api/resources/dev.cert.pem", "./api/resources/dev.key.pem")
		if err != nil {
			log.Println(err)
		}
	} else {
		err := server.Run(fmt.Sprintf(":%s", config.ServerPort))
		if err != nil {
			log.Println(err)
		}
	}
}
