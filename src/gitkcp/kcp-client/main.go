package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"time"

	kcp "github.com/xtaci/kcp-go"
)

func main() {
	conn, err := kcp.Dial("127.0.0.1:8090")
	if err != nil {
		log.Println(err)
		return
	}
	go OnClientHandleConn(conn)
	for {
		data := fmt.Sprintf("hello kcp %d", time.Now().Unix())
		_, err := conn.Write([]byte(data))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Second * 10)
	}
}

//OnClientHandleConn  服务端数据处理
func OnClientHandleConn(conn net.Conn) {
	datas := bytes.NewBuffer(nil)
	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Println(err)
		} else {
			// datas.WriteString("success rec:")
			datas.Write(buf[0:n])
			log.Printf("Client Recv -%s -%d -%s %s\n", conn.RemoteAddr().String(), n, hex.EncodeToString(datas.Bytes()), string(datas.Bytes()))
		}
		// log.Println("ASCII STRING:", string(datas.Bytes()))
		// conn.Write(datas.Bytes())
		datas.Reset()
	}
}
