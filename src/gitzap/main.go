package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"strconv"
	"time"
)

func main(){
	log.Println(Decimal(127.014526,2))

	//https://blog.csdn.net/skh2015java/article/details/81771808
	logger,_:=zap.NewProduction()
	defer logger.Sync()

	logger.Info("Test Debug",zap.String("name","tiptok"))

	sugar :=logger.Sugar()
	sugar.Infow("failed too fetch URL","" +
		"url","github\tiptok",
		"time",time.Now().Unix(),
	)
	sugar.Infof("failed to fetch url:%s","github\tiptok")

	sugar.Infow("failed too fetch URL.struct log",
		zap.String("url","github\aaa"),
		zap.Int64("time",time.Now().Unix()),
	)
}

func Decimal(value float64,reserveSize int) float64 {
	var f = fmt.Sprintf("%%.%df", reserveSize)
	value, _ = strconv.ParseFloat(fmt.Sprintf(f, value), 64)
	return value
}
