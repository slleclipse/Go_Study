package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {


	//concurrent computing
	u := & Vector{1.0,2.0, 3.0, 4.0, 5,6,7,8,9,10,11,12,13,14,15,16}
	u.DoAll( *u )

	//once
	/*ch1 := make(chan int)
	ch2 := make(chan int)
	twoprint(ch1, ch2)
	runtime.Gosched()
	<-ch1
	<-ch2*/

	newch := make(chan string)
	go func() {
		newch <- "jj"
		s := <- newch
		fmt.Println("func: ", s)

	}()
    runtime.Gosched()

	s := <-newch
	newch <- "hh"
	fmt.Println("new chan", s)
}



type Vector []float64
// 分配给每个CPU的计算任务
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1 // 发信号告诉任务管理者我已经计算完成了
}
const NCPU = 4 // 假设总共有16核
func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU) // 用于接收每个CPU的任务完成信号
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
	// 等待所有CPU的任务完成
	for i := 0; i < NCPU; i++ {
		<-c // 获取到一个数据，表示一个CPU计算完成了
	}
	// 到这里表示所有计算已经结束
}

func (v Vector) Op(f float64) float64 {
	return f
}


var a string
var once sync.Once
func setup() {
	a = "hello, world"
}
func doprint(ch chan int) {
	once.Do(setup)
	print(a)
	ch <- 1
}
func twoprint(ch1 chan int, ch2 chan int) {
	go doprint(ch1)
	go doprint(ch2)

}