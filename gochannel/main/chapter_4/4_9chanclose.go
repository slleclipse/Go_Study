package chapter_4

import "fmt"

func main() {

	dataChan := make(chan int, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() {
		<- syncChan1
		for {
			if elem, ok := <- dataChan; ok {
				fmt.Printf("Received: %d [receiver]\n", elem)
			} else {
				break
			}
		}
		fmt.Println("Done. [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			dataChan <- i
			fmt.Printf("Sent: %d [sender]\n", i)
		}
		close(dataChan)//虽然通道关闭，但对接收操作并无影响，接收方依然可以在接受完所有值后自行结束
		syncChan1 <- struct{}{}
		fmt.Println("Done. [sender]")
		syncChan2 <- struct{}{}
	}()

	<- syncChan2
	<- syncChan2
}
