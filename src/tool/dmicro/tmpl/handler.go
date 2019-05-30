package tmpl

var (
	HandlerFNC = `package handler

import (
	"context"

	example "{{.Dir}}/proto/example"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) Call(ctx context.Context, req *example.Request, rsp *example.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}
`

	HandlerSRV = `package handler

import (
	"context"

	"github.com/micro/go-log"

	example "{{.Dir}}/model/proto/example"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) Call(ctx context.Context, req *example.Request, rsp *example.Response) error {
	log.Log("Received Example.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Example) Stream(ctx context.Context, req *example.StreamingRequest, stream example.Example_StreamStream) error {
	log.Logf("Received Example.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&example.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Example) PingPong(ctx context.Context, stream example.Example_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&example.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
`

	SubscriberFNC = `package subscriber

import (
	"context"

	"github.com/micro/go-log"

	example "{{.Dir}}/proto/example"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}
`

	SubscriberSRV = `package subscriber

import (
	"context"
	"github.com/micro/go-log"

	example "{{.Dir}}/proto/example"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *example.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
`

	HandlerAPI = `package handler

import (
    "log"

	"github.com/gin-gonic/gin"
)

type Example struct{}

func (s *Example) Anything(c *gin.Context) {
	log.Print("Received Example.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func (s *Example) Hello(c *gin.Context) {
	log.Print("Received Example.Hello API request")

	name := c.Param("name")
	log.Print(name)
	

	c.JSON(200, map[string]string{
		"message": "Hi, hello "+name,
	})
}
`

	HandlerWEB = `package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	example "github.com/micro/examples/template/srv/proto/example"
)

func ExampleCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	exampleClient := example.NewExampleService("go.micro.srv.template", client.DefaultClient)
	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
`
)
