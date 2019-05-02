package main

import (
	"log"
	"net"
	"strconv"

	"github.com/tiptok/OFAPPTest/src/grpcServer/inf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = "9090"
)

func main() {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	svr := grpc.NewServer()
	inf.RegisterUserDataServer(svr, &UserData{})
	log.Printf("grpc server in: %s", port)
	svr.Serve(lis)
}

type UserData struct{}

func (d *UserData) GetUser(ctx context.Context, request *inf.UserReq) (response *inf.UserRsp, err error) {
	response = &inf.UserRsp{
		Name: strconv.Itoa(int(request.UserId)) + ":tip",
	}
	return response, err
}
