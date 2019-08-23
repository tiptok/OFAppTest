package syncT

import (
	"fmt"
	"log"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

//atomic.Value  原子操作
func TestAtomicValue(t *testing.T){
	var v atomic.Value
	var m int64
	go func(){//write
		for{
			m = time.Now().Unix()
			v.Store(m)
			log.Println("Store:",m)
			time.Sleep(time.Second*10)

		}
	}()
	go func(){//read
		for{
			time.Sleep(time.Second*1)
			value :=v.Load()
			log.Println("Get:",value.(int64))
		}
	}()

	/*unsafe.Pointer*/
	bytes:=[]byte{104, 101, 108, 108, 111}
	p :=unsafe.Pointer(&bytes)
	str :=(*string)(p)
	log.Println("unsafe.pointer:",str,*str)
	time.Sleep(time.Second*60)
}

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
