package ioT

import (
	"fmt"
	"io"
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
