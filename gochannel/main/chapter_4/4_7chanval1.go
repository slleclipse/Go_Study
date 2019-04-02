package chapter_4

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() { //接收
		for {
			if elem, ok := <- mapChan; ok {
				elem["count"]++
			} else {
					break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()

	go func() { //发送
		countMap := make(map[string]int)
		for i:= 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()

	<- syncChan
	<- syncChan
}