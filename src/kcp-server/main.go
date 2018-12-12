package main

import (
	"bytes"
	"encoding/hex"
	"log"
	"net"

	kcp "github.com/xtaci/kcp-go"
)

func main() {
	listen, err := kcp.Listen("0.0.0.0:8090")
	if err != nil {
		log.Println(err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go OnServerHandleConn(conn)
	}
}

//OnServerHandleConn  服务端数据处理
func OnServerHandleConn(conn net.Conn) {
	datas := bytes.NewBuffer(nil)
	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Println(err)
		} else {
			//datas.WriteString("success rec:")
			datas.Write(buf[0:n])
			log.Printf("Recv -%s -%d -%s\n", conn.RemoteAddr().String(), n, hex.EncodeToString(datas.Bytes()))
		}
		log.Println("STRING:", string(datas.Bytes()))
		conn.Write(datas.Bytes())
		datas.Reset()
	}
}
