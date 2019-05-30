package main

import (
	"ddmicro/app/service/main/account/model/proto/example"
	"ddmicro/app/service/main/push/handler"
	"github.com/micro/go-micro"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-log"
	"github.com/micro/go-web"
)

func main() {
	// New Service
	service := web.NewService(
		web.Name("go.micro.api.push"),
		web.Version("latest"),
	)

    // Initialise service
	service.Init()
	// Create RESTful handler (using Gin)
	example := new(handler.Example)

	msvr := micro.NewService()
	msvr.Init()
	example.Client = go_dmicro_srv_account.NewExampleService("go.micro.srv.account",msvr.Client())

	router := gin.Default()
	router.GET("/push", example.Anything)
	router.GET("/push/hello/:name", example.Hello)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
