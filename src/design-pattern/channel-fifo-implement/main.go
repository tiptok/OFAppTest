package main

import (
	"fmt"
	"time"
)

// 通过channel实现fifo模式
func main() {
	var writeRspChan = make(chan error)
	var writeChan = make(chan int)
	var end = make(chan bool)
	var exit = make(chan bool)

	go func() {
		for {
			select {
			case v := <-writeChan:
				if v%2 == 0 {
					time.Sleep(time.Second)
					writeRspChan <- func() error { return fmt.Errorf("error in pos:%v", v) }()
				} else {
					writeRspChan <- nil
				}
			case <-exit:
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		pos := i
		err := func(index int) error {
			fmt.Println("write:", index)
			//fifo
			writeChan <- index
			return <-writeRspChan
		}(pos)
		if err != nil {
			fmt.Println(err)
		}
	}
	exit <- true
	//1.阻塞方法1
	//var exit= make(chan os.Signal)
	//signal.Notify(exit,syscall.SIGINT, syscall.SIGTERM)
	//<-exit

	// 再阻塞之前,需要有运行的业务代码
	go func() {
		for {
			time.Sleep(time.Second * 5)
		}
	}()
	//2.阻塞方法2
	<-end

	//3.阻塞方法3
	//select{}
	//fmt.Println("end...")
}
