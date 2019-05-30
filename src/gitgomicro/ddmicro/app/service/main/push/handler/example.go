package handler

import (
	"bytes"
	"context"
	"ddmicro/app/service/main/account/model/proto/example"
	"encoding/json"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Example struct{
	Client go_dmicro_srv_account.ExampleService
}

func (s *Example) Anything(c *gin.Context) {
	log.Print("Received Example.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func (s *Example) Hello(c *gin.Context) {
	log.Print("Received Example.Hello API request")

	name := c.Param("name")

	//直接调用rpc接口
	s.CallDirect(name,"go.micro.svr.account")

	//rsp :=&go_dmicro_srv_account.Response{}
	//s.CallHttp("http://localhost:8080/account/call",&go_dmicro_srv_account.Request{Name:name},rsp)
	//fmt.Println("CallHttp response:",rsp.Msg)

	s.CallHttp2("http://localhost:8080/account/call",name)
	//fmt.Println("CallHttp response:",rsp.Msg)

	//s.CallHttp3("http://localhost:8080/rpc","go.micro.srv.account","Example.Call",&go_dmicro_srv_account.Request{Name:name})

	c.JSON(200, map[string]string{
		"message": "Hi, hello "+name,
	})
}
//CallDirect  直接通过go-micro 实例micro-client 调用服务
func (s *Example) CallDirect(name string,svrNS string)error{
	//cli :=example.
	log.Println("CallDirect")
	rsp,err:= s.Client.Call(context.Background(),&go_dmicro_srv_account.Request{Name:name})
	if err!=nil{
		log.Println("CallDirect error:",err)
		return err
	}
	log.Println(fmt.Sprintf("CallDirect service:%v Response:%v",svrNS,rsp))
	return nil
}
//CallHttp 通过 http content-type:app/protobuf  url:127.0.0.1/svr-name/method-name
func(s *Example)CallHttp(url string,request proto.Message,rsp proto.Message)error{
	log.Println("CallHttp",url)
	req,err :=proto.Marshal(request)//
	if err!=nil{
		log.Println("CallHttp err",err)
		return err
	}
	r, err := http.Post(url, "application/protobuf", bytes.NewReader(req))
	if err != nil {
		fmt.Println("CallHttp err",err)
		return err
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("CallHttp err",err)
		return err
	}
	if err := proto.Unmarshal(b, rsp); err != nil {
		fmt.Println("CallHttp err",err)
		return err
	}
	return nil
}
//CallHttp2  通过 http content-type:app/protobuf  url:127.0.0.1/svr-name/method-name
func(s *Example)CallHttp2(url string,name string)error{
	log.Println("CallHttp",url)
	req,err :=proto.Marshal(&go_dmicro_srv_account.Request{Name:name})//
	if err!=nil{
		log.Println("CallHttp err",err)
		return err
	}
	r, err := http.Post(url, "application/protobuf", bytes.NewReader(req))
	if err != nil {
		fmt.Println("CallHttp err",err)
		return err
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("CallHttp err",err)
		return err
	}
	rsp :=&go_dmicro_srv_account.Response{}
	if err := proto.Unmarshal(b, rsp); err != nil {
		fmt.Println("CallHttp err",err)
		return err
	}else{
		fmt.Println("CallHttp response:",rsp.Msg)
	}
	return nil
}
//CallHttp3  通过 http 请求Rpc服务  url:127.0.0.1:8080/rpc
func(s *Example)CallHttp3(url string,svr string,method string,req interface{})error{
	log.Println("CallHttp",url)
	type RpcRequest struct {
		Service string `json:"service"`
		Method  string `json:"method"`
		Request interface{} `json:"request"`
	}
	rpcReq :=&RpcRequest{
		Service:svr,
		Method:method,
		Request:req,
	}
	jsData,err :=json.Marshal(rpcReq)//
	if err!=nil{
		log.Println("CallHttp err",err)
		return err
	}
	r, err := http.Post(url, "application/json", bytes.NewReader(jsData))
	if err != nil {
		fmt.Println("CallHttp err",err)
		return err
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("CallHttp err",err)
		return err
	}
	if len(b)>0{
		log.Println("CallHttp response:",string(b))
	}
	return nil
}