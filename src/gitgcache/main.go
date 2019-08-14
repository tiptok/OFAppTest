package main

import (
	"github.com/bluele/gcache"
	"log"
	"time"
)

func main(){
	gc :=gcache.New(10).LRU().Build()
	gc.Set("tip","tok")
	if v,err :=gc.Get("tip");err!=nil{
		panic(err)
	}else{
		log.Println("Get:",v)
	}
	gc.SetWithExpire("help","go",time.Second*10)
}
