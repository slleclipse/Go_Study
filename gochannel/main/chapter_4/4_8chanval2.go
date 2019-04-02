package chapter_4

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

var mapChan1 = make(chan map[string]*Counter, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() { //接收
		for {
			if elem, ok := <- mapChan1; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()

	go func() { //发送
		countMap := map[string]*Counter {
			"count": &Counter{},
		}
		for i:= 0; i < 5; i++ {
			mapChan1 <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan1)  //在发送端关闭通道，不能在接收端关闭通道
		syncChan <- struct{}{}
	}()

	<- syncChan
	<- syncChan
}

