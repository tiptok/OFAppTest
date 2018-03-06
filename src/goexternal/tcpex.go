package goexternal

import (
	"sync"

	"log"

	"github.com/tiptok/gotransfer/comm"
	"github.com/tiptok/gotransfer/conn"
)

const (
	bizIP   = "127.0.0.1"
	bizPort = 8050
)

type tcpExserver struct {
	BizManager comm.DataContext
}

func (srv *tcpExserver) OnConnect(c *conn.Connector) bool {
	node := NewTcpBizNode(c, bizIP, bizPort)
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
			log.Println(err)
		}
	}()
	key := c.RemoteAddress
	node, isEx := srv.BizManager.GetOk(key)
	log.Println("OnClose:", key, " ManageExists:", isEx)
	node.(*TcpBizNode).Disposed()
	srv.BizManager.Delete(key)
	_, isEx = srv.BizManager.GetOk(key)
	log.Println("OnClose:", key, " ManageExists:", isEx)
}

/*
   链接交换节点
*/
type TcpBizNode struct {
	/*分发地址+*/
	Key string
	/*源*/
	Src *conn.Connector
	/*目标*/
	Dst   *conn.Connector
	mutex *sync.RWMutex
}

//发送给 src
func (node *TcpBizNode) SendToSrc(d conn.TcpData) bool {
	if node.Src.IsConneted {
		node.Src.SendChan <- d
	} else {
		return false
	}
	return true
}

//发送给 dst
func (node *TcpBizNode) SendToDst(d conn.TcpData) bool {
	if node.Dst.IsConneted {
		node.Dst.SendChan <- d
	} else {
		return false
	}
	return true
}

func (node *TcpBizNode) GetKey() string {
	return node.Key
}

//关闭释放
func (node *TcpBizNode) Disposed() {
	node.mutex.Lock()
	node.Src.Close()
	//node.Dst.Close()
	node.mutex.Unlock()
}

func NewTcpBizNode(src *conn.Connector, dstIP string, dstPort int) *TcpBizNode {
	var tcpClient conn.TcpClient
	tcpClient.NewTcpClient(dstIP, dstPort, 500, 500)
	handler := &TcpExClient{}
	if !tcpClient.Start(handler) {
		//client 启动失败
	}
	node := &TcpBizNode{
		Key:   src.RemoteAddress,
		Src:   src,
		Dst:   tcpClient.Conn,
		mutex: new(sync.RWMutex),
	}
	handler.BizNode = node
	return node
}

type TcpExClient struct {
	BizNode *TcpBizNode
}

func (cli *TcpExClient) OnConnect(c *conn.Connector) bool {
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
