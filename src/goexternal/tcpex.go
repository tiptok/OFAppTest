package goexternal

import (
	"log"

	"time"

	"github.com/tiptok/gotransfer/comm"
	"github.com/tiptok/gotransfer/conn"
)

const (
	bizIP   = "127.0.0.1"
	bizPort = 9928
)

type tcpExserver struct {
	BizManager comm.DataContext
}

func (srv *tcpExserver) OnConnect(c *conn.Connector) bool {
	node := NewTcpBizNode(c, bizIP, bizPort)
	log.Println("Link Size:", len(srv.BizManager.DataStore))
	srv.BizManager.Set(node.Key, node)
	return true
}
func (srv *tcpExserver) OnReceive(c *conn.Connector, d conn.TcpData) bool {
	key := c.RemoteAddress
	node := srv.BizManager.Get(key)
	node.(*TcpBizNode).SendToDst(d)
	return true
}
func (srv *tcpExserver) OnClose(c *conn.Connector) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("tcpExserver OnClose err", err)
		}
	}()
	key := c.RemoteAddress
	node, isEx := srv.BizManager.GetOk(key)
	log.Println("OnClose First:", key, " ManageExists:", isEx)
	//node.(*TcpBizNode).Disposed()

	if tcpnode, isNode := node.(*TcpBizNode); isNode {
		tcpnode.Disposed() //异常
	}
	srv.BizManager.Delete(key)
	_, isEx = srv.BizManager.GetOk(key)
	log.Println("OnClose Second:", key, " ManageExists:", isEx)
	time.Sleep(time.Millisecond * 50)
}

type TcpExClient struct {
	BizNode *TcpBizNode
}

func (cli *TcpExClient) OnConnect(c *conn.Connector) bool {
	log.Println("client On connect", c.RemoteAddress)
	return true
}
func (cli *TcpExClient) OnReceive(c *conn.Connector, d conn.TcpData) bool {
	cli.BizNode.SendToSrc(d)
	return true
}
func (cli *TcpExClient) OnClose(c *conn.Connector) {
	//从列表移除
	log.Println("Client OnClose:", c.RemoteAddress)
}
