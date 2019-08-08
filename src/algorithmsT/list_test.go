package main

import (
	"container/list"
	"fmt"
	"log"
	"sync"
	"testing"
)

func TestList(t *testing.T) {
	l := list.New()
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println("")
	fmt.Print(l.Front().Value) //输出首部元素的值,0
	fmt.Print(l.Back().Value)  //输出尾部元素的值,4

	fmt.Println("")
	l.InsertAfter(6, l.Front()) //首部元素之后插入一个值为6的元素
	PrintList(l)

	l.MoveBefore(l.Front(), l.Back()) //首部头移动到尾部之前
	PrintList(l)

	l.MoveToFront(l.Back()) //将尾部元素移动到首部
	PrintList(l)

	l2 := list.New()
	l2.PushBack("j")
	l2.PushBackList(l) //将l中元素放在l2的末尾
	PrintList(l2)
}

func PrintList(l *list.List) {
	fmt.Println("")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println("")
}

func Test_tmp(t *testing.T){
	input :=20010
	log.Println(fmt.Sprintf("%.2f",float64(input)/float64(100.0)))
}

func Test_G(t *testing.T){
	wg := new(sync.WaitGroup)
	chn := make(chan int, 10)
	defer close(chn)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {

			defer wg.Done()
			if i%4 == 0 {
				return
			}
			chn <- i
		}(i)
	}
	wg.Wait()
	log.Println(cap(chn))
	rmp := make(map[int]bool)
	for i := 0; i < len(chn); i++ {
		if c, ok := <-chn; ok {
			rmp[c] = true
		}
	}
	log.Println(rmp)
}
