package main

import (
	"fmt"
	"time"
)

func main() {
	//ch := make(chan int, 10)
	ch := make(chan int)
	go produce(ch)
	go consume(ch)

	time.Sleep(10000000000)
}
//带缓冲的channel是异步的方式进行，非缓冲channel是同步的方式进行
func produce(ch chan int) {
	fmt.Println("start produce")
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("produce: ", i)

	}
	fmt.Println("finished produce")
}

func consume(ch chan int) {
	fmt.Println("start consume")
	for i := 0; i < 10; i++ {
		x := <- ch
		fmt.Println("consume ", x)
	}
	fmt.Println("finished consume")
}
