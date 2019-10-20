package ioT

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestIO(t *testing.T) {
	data, err := ReadFrom(strings.NewReader("from string..."), 12)
	if err == nil {
		fmt.Println(string(data))
	}

	file, err := os.Open("E:/app/Go_Myeclipse/GoWorkSpace/src/github.com/tiptok/OFAppTest/RecordLog.txt")
	defer file.Close()
}

func ReadFrom(r io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := r.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func TestMultiRead(t *testing.T){
	r1:=strings.NewReader("tip")
	r2:=strings.NewReader("tok")
	mr :=io.MultiReader(r1,r2)
	var buf []byte = make([]byte,10)
	rn,_:=mr.Read(buf) //只会读取到第一个reader
	t.Log("Read:",string(buf),rn)
	io.ReadFull(mr,buf[rn:])//io.ReadFull(mr,buf)
	t.Log(string(buf))

	allocs :=testing.AllocsPerRun(1000,func(){
		io.Copy(os.Stdout,mr)
	})
	t.Log("allocs:",allocs)
}

func TestMultiWriter(t *testing.T){
	r :=strings.NewReader("tiptok")
	var buf1,buf2 bytes.Buffer
	w :=io.MultiWriter(&buf1,&buf2)
	if _,err :=io.Copy(w,r);err!=nil{ //同时写入两个writer
		log.Fatal(err)
	}
	t.Log(buf1.String())
	t.Log(buf2.String())
}