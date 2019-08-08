package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

type NamesData struct {
	Names []string `json:"names"`
}

func main(){
	defer func(){
		if p:=recover();p!=nil{
			log.Println(p)
		}
		log.Println("app stop.")
	}()
	f, _ := os.Create("gin.log")
	jsonData,err:=ioutil.ReadFile("jd.txt")
	if err!=nil{
		log.Println(err)
		return
	}
	var namesData NamesData
	if err:=json.Unmarshal(jsonData,&namesData);err!=nil{
		log.Println(err)
		return
	}
	firstName = namesData.Names
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()
	nameHandler := r.Group("/name")
	{
		nameHandler.GET("/gen",GenName )
	}
	log.Println("app start.")
	r.Run(":8090")
}
func ServerSuccess(g *gin.Context, msg string) {
	g.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": msg,
	})
}

var firstName = []string{}

func GenName(g *gin.Context) {
	rspMsg := ""
	t,o :=0,0
	if v:=g.Query("times");len(v)!=0{
		t,_ = strconv.Atoi(v)
	}
	if v:=g.Query("onece");len(v)!=0{
		o,_ = strconv.Atoi(v)
	}
	rspMsg = getName(t,o)
	g.String(200,rspMsg)
}

func getName(t,o int)string{
	if t==0 || o==0{
		log.Println("t / o not 0")
		return ""
	}
	var times int =t
	var once int =o
	l :=len(firstName)
	getChar :=func()string{
		index:=rand.Intn(l)%(l)
		if index<=1{
			index =1
		}
		return firstName[index-1]
	}
	buf:=bytes.NewBufferString("")
	for t:=0;t<times;t++{
		for i:=0;i<once;i++{
			buf.WriteString(fmt.Sprintf("%6s",getChar()+getChar()))
		}
		//返回姓名
		buf.WriteString("\n")
	}
	return buf.String()
}
