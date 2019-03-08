package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/tiptok/OFAPPTEST/src/gitGin/docs"

	"github.com/tiptok/OFAPPTEST/src/gitGin/mid/jwt"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	//
	//gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//Listen

	sd := r.Group("/sd")
	sd.Use(AuthMidware())
	sd.Use(NoCache)
	sd.Use(jwt.JWTMiddleware())
	sd.Use(Secure)
	{
		sd.GET("/health", HealthCheck)
		sd.GET("/cpu", CPUCheck)
	}

	demo := r.Group("/demo")
	{
		demo.GET("/user/:name/*action", DemoUserAction)
		demo.GET("/player/:name", DemoUserAction) //匹配路由 demo/player/xxx
		demo.GET("/welcome", DemoWelcome)
		demo.POST("/form_post", Demoform_post)
		demo.GET("/auth", DemoAuth)

		demo.POST("/loginJSON", LoginJSON)
		demo.POST("/loginForm", LoginForm)
		demo.Any("/sbq", DemoSBQ)

		demo.GET("/somexml", func(c *gin.Context) {
			c.XML(http.StatusOK, gin.H{"message": "hello"})
		})
		demo.GET("/someyaml", func(c *gin.Context) {
			c.YAML(http.StatusOK, gin.H{"message": "hello", "title": "azmaze"})
		})
		demo.GET("/somejson", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "hello", "title": "azmaze"})
		})
	}

	alarmflag := r.Group("/alarmflag")
	{
		alarmflag.POST("/save", AlarmFlagSave)
		alarmflag.POST("/delete", AlarmFlagDelete)
		alarmflag.POST("/update", AlarmFlagUpdate)
	}
	r.Run(":8081")
}

// @Summary A ping tool
// @Description Response ping
// @Tags ping
// @Accept  json
// @Produce  json
// @Success 200
// @Router /ping [get]
func ping(c *gin.Context) {
	c.String(200, "pong")
}

// @Summary Shows OK as the ping-pong result
// @Description Shows OK as the ping-pong result
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200 {string} plain "OK"
// @Router /sd/health [get]
func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}

// @Summary Check Cup used info
// @Description Check Cup used info
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200 {string} plain "70%"
// @Router /sd/cpu [get]
func CPUCheck(c *gin.Context) {
	message := "70%"
	c.String(http.StatusOK, "\n"+message)
}

//http://localhost:8081/demo/user/jim/wingame
//Response : jim Do /wingame
func DemoUserAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	action = strconv.Quote(action)
	//http.StatusRequestEntityTooLarge
	message := name + " Do " + action
	c.String(http.StatusOK, "\n"+message)
}

//http://localhost:8081/demo/player/tiptok
//tiptok Do

//http://localhost:8081/demo/welcome?lastname=tiptok   hello geust tiptok
//demo/welcome?firstname=ccc&&lastname=tiptok hello ccc tiptok
func DemoWelcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "geust")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "hello %s %s", firstname, lastname)
}

//Post formData
func Demoform_post(c *gin.Context) {
	message := c.PostForm("message")
	errmessage := c.DefaultPostForm("errm", "none")
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    errmessage,
	})
}

// @Summary Get jwt token
// @Description Get jwt token
// @Tags demo
// @Accept  json
// @Produce  json
// @Success 200 {string} plain "70%"
// @Router /demo/auth [get]
func DemoAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	token, err := jwt.GenerateToken(username, password)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

//http://localhost:8081/demo/sbq?name=tip&address=fz get
func DemoSBQ(c *gin.Context) {
	var p Person
	if c.ShouldBindQuery(&p) == nil {
		log.Println(p.Name)
		log.Println(p.Address)
	}
	if c.ShouldBind(&p) == nil {
		log.Println("ShouldBind", p.Name)
		log.Println("ShouldBind", p.Address)
	}
	c.String(200, "Success")
}
