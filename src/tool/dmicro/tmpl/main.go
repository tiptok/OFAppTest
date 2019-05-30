package tmpl

var (
	MainFNC = `package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"{{.Dir}}/handler"
	"{{.Dir}}/subscriber"
)

func main() {
	// New Service
	function := micro.NewFunction(
		micro.Name("{{.FQDN}}"),
		micro.Version("latest"),
	)

	// Initialise function
	function.Init()

	// Register Handler
	function.Handle(new(handler.Example))

	// Register Struct as Subscriber
	function.Subscribe("{{.FQDN}}", new(subscriber.Example))

	// Run service
	if err := function.Run(); err != nil {
		log.Fatal(err)
	}
}
`

	MainSRV = `package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"{{.Dir}}/handler"
	//"{{.Dir}}/subscriber"

	example "{{.Dir}}/model/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("{{.FQDN}}"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("{{.FQDN}}", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("{{.FQDN}}", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
`
	MainAPI = `package main

import (
	"{{.Dir}}/handler"
    //example "{{.Dir}}/model/proto/example"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-log"
    "github.com/micro/go-web"
)

func main() {
	// New Service
	service := web.NewService(
		web.Name("{{.FQDN}}"),
		web.Version("latest"),
	)

    // Initialise service
	service.Init()

	// Create RESTful handler (using Gin)
	example := new(handler.Example)
	router := gin.Default()
    
	router.GET("/{{.Dir}}", example.Anything)
	router.GET("/{{.Dir}}/hello/:name", example.Hello)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
`
	MainWEB = `package main

import (
        "github.com/micro/go-log"
	"net/http"

        "github.com/micro/go-web"
        "{{.Dir}}/handler"
)

func main() {
	// create new web service
        service := web.NewService(
                web.Name("{{.FQDN}}"),
                web.Version("latest"),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/example/call", handler.ExampleCall)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
`
)
