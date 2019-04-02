package chapter_4

import "fmt"

func main() {
	chanCap := 5
	intChan := make(chan int, chanCap)

	//如果同时有多个case满足条件，那么运行时系统会通过一个伪随机的算法选择一个case
	for i := 0; i < chanCap; i++ {
		select {
		case intChan <- 1:
		case intChan <- 2:
		case intChan <- 3:
		}
	}

	for i := 0; i < chanCap; i++ {
		fmt.Printf("%d\n", <-intChan)
	}
}
