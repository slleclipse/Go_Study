package main

import (
	"fmt"
)

var counter1 int = 0

func main() {

	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go count(chs[i], i)
	}
	for i := 0; i < 10; i++ {
		//fmt.Println(<-chs[i])
		x, ok := <- chs[i]
		fmt.Println(x, ok)
	}

}

func count(ch chan int,i int) {
	counter1++
	ch <- counter1
	//fmt.Println("counting:", counter1)
	fmt.Println(i,"is over")
}