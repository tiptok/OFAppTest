package syncT

import (
	"fmt"
	"log"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	runtime.GOMAXPROCS(2)
	var Index int32 = 0
	chW := make(chan int, 20)
	atomic.AddInt32(&Index, 1)
	for i := 0; i < 10; i++ {
		go AddInt(&Index, chW)
	}
	for i := 0; i < 10; i++ {
		go AddInt(&Index, chW)
	}
	for i := 0; i < 20; i++ {
		<-chW
	}
	//log.Printf("Is True:%v", atomic.CompareAndSwapInt32(&Index, 1, 0))
	log.Println("Index:", Index)
	t.Log("End")
}

func AddInt(num *int32, ch chan int) {
	for i := 0; i < 200; i++ {
		//*num += 1
		atomic.AddInt32(num, 1)
		//
	}
	ch <- 1
	log.Println("*num:", *num)
}

func TestLoadInt(t *testing.T) {
	runtime.GOMAXPROCS(2)
	var op int32 = 0
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				time.Sleep(time.Nanosecond)
				atomic.AddInt32(&op, 1)
				if i == 50 {
					fmt.Println("ops:", atomic.LoadInt32(&op), op)
				}
			}
		}()
	}
	time.Sleep(time.Second)
}
