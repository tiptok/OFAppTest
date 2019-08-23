package main

import (
	"flag"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)



func main(){
	var topic,channel string
	fs :=flag.NewFlagSet("producer",1)
	flag.StringVar(&topic,"topic","my2","set public topic name")
	flag.StringVar(&channel,"channel","tip-chan","set consumer channel name")
	fs.String("addr","127.0.0.1:4150","set nsqd address")
	addr :=fs.Lookup("addr").Value.(flag.Getter).Get().(string)
	p,err:= nsq.NewProducer(addr,nsq.NewConfig())
	if err!=nil{
		log.Fatal(err)
	}
	go func(){ //producer  use nsqd tcp port
		for {
			message :=fmt.Sprintf("%v",time.Now().Unix())
			if err:=p.Publish(topic,[]byte(message));err!=nil{
				log.Println("publish ",topic,err)
				log.Println(err)
			}
			time.Sleep(time.Second*5)
		}
	}()
	go func(){//consumer1  user nsqlookup http port
		c,err:= nsq.NewConsumer(topic,channel,nsq.NewConfig())
		c.AddHandler(&ConsumerHandler{})
		if err!=nil{
			log.Println(err)
			return
		}
		if err:=c.ConnectToNSQLookupd("127.0.0.1:4161");err!=nil{  //connect ot nsqlookup by http
			log.Println(err)
		}
		for {
			time.Sleep(time.Second*5)
		}
	}()
	go func(){//consumer2
		c,err:= nsq.NewConsumer(topic,channel,nsq.NewConfig())
		c.AddHandler(&ConsumerHandler2{})
		if err!=nil{
			log.Println(err)
			return
		}
		if err:=c.ConnectToNSQLookupd("127.0.0.1:4161");err!=nil{  //connect ot nsqlookup by http
			log.Println(err)
		}
		for {
			time.Sleep(time.Second*5)
		}
	}()
	time.Sleep(time.Second*60*60)
}

type ConsumerHandler struct{}

func(*ConsumerHandler)HandleMessage(msg *nsq.Message)error{
	log.Println("recv:",string(msg.Body))
	return nil
}

type ConsumerHandler2 struct{}

func(*ConsumerHandler2)HandleMessage(msg *nsq.Message)error{
	log.Println("handler2 recv:",string(msg.Body))
	return nil
}