package main

import "fmt"

func main() {
	var a APrivate = BImplementAPrivate{}
	a.MethodClose(0)
	fmt.Println("end ...")
}

type APrivate interface {
	MethodClose(v int)
	//不想被其他用户实现这个接口时，定义一个私有方法
	private()
}

type BImplementAPrivate struct {
}

func (b BImplementAPrivate) MethodClose(v int) {

}

// 如果没有实现private() 则 BImplementAPrivate 没有实现接口 APrivate
func (b BImplementAPrivate) private() {

}
