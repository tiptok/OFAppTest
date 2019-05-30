package handler

import (
	"fmt"
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

func(s *Example)Push(c *gin.Context){
	type PushMessage struct {
		To string `json:"to"`
		From string `json:"from"`
		Message string `json:"msg"`
	}
	var pushMsg PushMessage
	err:= c.ShouldBind(&pushMsg)
	if err!=nil{
		c.JSON(500, map[string]string{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, map[string]string{
		"message": fmt.Sprintf("%s send msg to %s msgcentent:%s",pushMsg.From,pushMsg.To,pushMsg.Message),
	})
}
