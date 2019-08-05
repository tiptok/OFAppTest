package main

import (
	"container/list"
	"fmt"
	"log"
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
