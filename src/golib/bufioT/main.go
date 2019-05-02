package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s) //NewReaderSize(rd, defaultBufSize)  4096
	// Peek 返回缓存的一个切片，该切片引用缓存中前 n 字节数据
	// 该操作不会将数据读出，只是引用
	// 引用的数据在下一次读取操作之前是有效的
	b, _ := br.Peek(5)
	b[0] = 66 //修改切片的值 原来的值也变化了
	fmt.Println(b, br.Buffered())
	fmt.Println(br, br.Buffered())

	// br.ReadByte
	// br.ReadBytes
	// br.ReadLine
	// br.ReadRune
	// br.ReadSlice
	// br.ReadString
	//br.Reset

	bf := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(bf) //bufio.NewReaderSize
	fmt.Println(bw.Available(), bw.Buffered())
	bw.WriteString("hello 下午")

	bw.Flush()
	fmt.Println(bf)
}
