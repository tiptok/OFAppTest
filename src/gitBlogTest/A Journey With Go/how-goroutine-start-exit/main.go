package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var a int
	go func() {
		var skip int
		for {
			_, file, line, ok := runtime.Caller(skip)
			if !ok {
				break
			}
			fmt.Printf("%s:%d \n", file, line)
			skip++
		}
		wg.Done()
	}()
	a++
	fmt.Println("value:", a)
	wg.Wait()
}
