package main

import (
	"fmt"
)

//20170904
//1.交替输出 数字 字母
var c_seq = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}

func Test_c_seq() {
	chan_n := make(chan bool, 1)
	chan_c := make(chan bool, 1)
	done := make(chan struct{})

	go func() {
		for i := 0; i < 11; i++ {
			<-chan_c
			fmt.Print(i)
			chan_n <- true
		}
		done <- struct{}{}
	}()

	go func() {
		for i := 0; i < 11; i++ {
			<-chan_n
			fmt.Print(c_seq[i])
			chan_c <- true
		}
	}()
	chan_n <- true //控制先后
	<-done
	fmt.Println("App stop")
}
