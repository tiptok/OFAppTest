package main

import (
	"fmt"
	"log"
	"math/rand"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/tiptok/OFAppTest/src/ogrpc/grpcServer/inf"
)

const (
	port = "9090"
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%s", port), grpc.WithInsecure())

	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(conn)
	}
	cli := inf.NewUserDataClient(conn)
	getUser(cli)
}

func getUser(cli inf.UserDataClient) {
	iRand := rand.Intn(100)
	req := inf.UserReq{
		UserId: int32(iRand),
	}
	rsp, err := cli.GetUser(context.Background(), &req)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("response:%v", rsp)
}
