package main

import (
	"go.uber.org/zap"
	"time"
)

func main(){
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
