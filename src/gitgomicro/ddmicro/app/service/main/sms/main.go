package main

import (
	"ddmicro/app/service/main/sms/handler"
    //example "sms/model/proto/example"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-log"
    "github.com/micro/go-web"
)

func main() {
	// New Service
	service := web.NewService(
		web.Name("go.micro.api.sms"),
	)
	service.Init()

	// Create RESTful handler (using Gin)
	example := new(handler.Example)
	router := gin.Default()
	router.GET("/sms", example.Anything)
	router.GET("/sms/hello/:name", example.Hello)
	//router.POST("/push",example.Push)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
