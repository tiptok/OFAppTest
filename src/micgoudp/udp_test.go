package micgoudp

import (
	"log"
	"testing"

	"time"

	"github.com/tiptok/gotransfer/conn"
)

func TestUDP(t *testing.T) {
	var udpcli conn.UpdClient
	udpcli.NewUpdClient("192.168.3.54", 8067, 500, 500)
	if !udpcli.Start(&transUpdClientHandler{}) {
		log.Println("Udp Start Error.")
	}
	time.Sleep(time.Second * 1000)
}

/*
	udp Client 处理
*/
type transUpdClientHandler struct {
}

func (trans *transUpdClientHandler) OnConnect(c *conn.Connector) bool {
	d := conn.NewTcpData([]byte("Hello"))
	c.SendChan <- d
	return true
}
func (trans *transUpdClientHandler) OnReceive(c *conn.Connector, d conn.TcpData) bool {
	sAddr, _ := c.LocalAddr()
	slocal := c.RemoteAddress
	c.SendChan <- d
	log.Println(sAddr, " ", slocal)
	return true
}
func (trans *transUpdClientHandler) OnClose(c *conn.Connector) {
	//从列表移除
	log.Println(c.RemoteAddress, "On Close...")
}
