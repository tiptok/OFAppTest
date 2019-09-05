package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var port int
var total int32

//测试负载代理
func main(){
	flag.IntVar(&port,"p",8081,"server listen port.")
	flag.Parse()


	//1.http.HandleFunc
	http.HandleFunc("/test",test)
	http.HandleFunc("/get",get)
	http.ListenAndServe(fmt.Sprintf(":%d",port),nil)


	//2.http.NewServeMux
	//mux :=http.NewServeMux()
	//mux.HandleFunc("/",test)
	//2.1
	//server :=http.Server{
	//	Addr:fmt.Sprintf(":%d",port),
	//	Handler:mux,
	//}
	//server.ListenAndServe()
	//2.2
	//if err:=http.ListenAndServe(fmt.Sprintf(":%d",port),mux);err!=nil{
	//	log.Fatal(err)
	//}
}

func test(w http.ResponseWriter,req *http.Request){
	atomic.AddInt32(&total,1)
	log.Println("test on port:",port,total)
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("test on port:%d total:%d",port,total)))
}

func get(w http.ResponseWriter,req *http.Request){
	//atomic.AddInt32(&total,1)
	log.Println("get on port:",port,total)
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("get on port:%d total:%d",port,total)))
}
