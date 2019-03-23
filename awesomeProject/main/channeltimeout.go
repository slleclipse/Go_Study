package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//time out
	ch := make(chan int)
	timeout := make(chan int)

	fmt.Println("Cpu num", runtime.NumCPU())

	go func () {
		time.Sleep(10000000000)
		timeout <- 1
	}()

	go func() {
		ch <- 1
	}()

	select {
	case <- ch:
		fmt.Println("not wait")
	case <- timeout:
		fmt.Println("timeout 1 seconds")
	}
}


