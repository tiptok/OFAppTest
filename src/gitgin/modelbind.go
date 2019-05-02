package main

import "github.com/gin-gonic/gin"
import "net/http"

type login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func LoginForm(c *gin.Context) {
	var data login
	//This will infer what binder to use depending on the content-type header
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if data.User == "tiptok" && data.Password == "123456" {
		c.JSON(http.StatusOK, gin.H{
			"status": "login in",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "unauthorize"})
	}
}

func LoginJSON(c *gin.Context) {
	var data login
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if data.User == "tiptok" && data.Password == "123456" {
		c.JSON(http.StatusOK, gin.H{
			"status": "login in",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "unauthorize"})
	}
}
