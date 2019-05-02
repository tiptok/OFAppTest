package ogoexternal

import (
	"github.com/tiptok/gotransfer/conn"

	"testing"

	"github.com/tiptok/gotransfer/comm"
)

func TestGoEx(t *testing.T) {
	exit := make(chan int, 1)
	var srv conn.TcpServer
	tcpExserver := &tcpExserver{}
	tcpExserver.BizManager = comm.DataContext{}
	go func() {
		srv.NewTcpServer(8085, 500, 500)
		srv.Start(tcpExserver)
	}()
	<-exit
}
