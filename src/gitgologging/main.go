package main

import (
	"os"
	"time"

	"github.com/op/go-logging"
)

/*
window下不支持终端颜色
color support github.com/daviddengcn/go-colortext
github.com/multiformats/go-multiaddr
https://blog.csdn.net/liuzhijun301/article/details/80433557
*/
var log = logging.MustGetLogger("OF")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func main() {
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")

	logging.SetBackend(backend1Leveled, backend2Formatter)

	var exit chan int
	log.Info("message info")
	log.Debug("message debug")
	log.Error("message error")
	log.Critical("message critical")
	go func() {
		for {
			time.Sleep(time.Second * 2)
		}
	}()
	<-exit
}
