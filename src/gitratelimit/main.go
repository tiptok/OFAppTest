package main

import (
	"bytes"
	"fmt"
	"github.com/juju/ratelimit"
	"log"
	"time"
)
//限流
func main(){
	bucketWithRate :=ratelimit.NewBucketWithRate(500,500)//rate 500 interval=1e9/500
	log.Println("rate :",bucketWithRate.Rate())


	bucket:= ratelimit.NewBucket(time.Millisecond*10,100)//1000/10  == 100 (rate)
	//t:=bucket.Take(50)
	//log.Println(t)
	log.Println(bucket.Take(1))
	//t=bucket.Take(2)
	//log.Println(t)
	log.Println(bucket.Available())
	log.Println(bucket.Rate())//增加速率
	buf:= bytes.NewBuffer(nil)
	for i :=0;i<1000;i++{
		buf.WriteString(fmt.Sprintf("%d-",i))
	}
	rd := ratelimit.Reader(buf,bucket)
	data :=make([]byte,64)
	count :=1
	go func(){
		for{
			time.Sleep(time.Second*5)
			log.Println("current bucket:",bucket.Available())
		}
	}()
	for ;count>0;{
		count,err:=rd.Read(data)
		if err!=nil{
			log.Fatal(err)
		}
		log.Println("read data:",count,string(data))
		//data =data[:0]
	}
}
