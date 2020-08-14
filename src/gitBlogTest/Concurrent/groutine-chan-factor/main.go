package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	resultChan := make(chan interface{}, 100)
	var wgSend sync.WaitGroup
	var wgReceive sync.WaitGroup

	wgReceive.Add(1)
	go doReceiveWorker(resultChan, &wgReceive) //&waitGroup 需要指针传入
	//wgSend.Add(10)
	for i := 1; i <= 10; i++ {
		wgSend.Add(1)
		go doSendWorker(i, 10, resultChan, &wgSend)
	}
	wgSend.Wait()
	time.Sleep(time.Second)
	fmt.Println("**********************  结束发送任务  **********************")

	close(resultChan)
	wgReceive.Wait()
}

// 发送
func doSendWorker(workId int, totalNumber int, ch chan<- interface{}, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		fmt.Printf("done send :%v \n", workId)
	}()
	for i := 1; i <= totalNumber; i++ {
		ch <- fmt.Sprintf("%v-%v", workId, i)
		time.Sleep(time.Second * 5)
		printWorker(workId, i)
	}
}

// 接收
func doReceiveWorker(ch <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		totalReceiveNumber int
		stopChan           chan struct{} = make(chan struct{})
	)

	//定时打印处理结果
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("当前处理数量:", totalReceiveNumber)
			case <-stopChan:
				return
			}
		}
	}()

	for data := range ch {
		totalReceiveNumber++
		fmt.Printf("receive :%v \n", data)
	}

	//结束信号
	stopChan <- struct{}{}

	fmt.Println("**********************  结束接收  **********************")
}

func printWorker(workId int, totalNumber int) {
	fmt.Printf("send workId:%v sn:%v \n", workId, totalNumber)
}
