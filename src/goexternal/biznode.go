package goexternal

import (
	"sync"

	"github.com/tiptok/gotransfer/conn"
)

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
	if node.Src.IsConneted{
		node.Src.Close()
	}
	node.Dst.Close()
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
