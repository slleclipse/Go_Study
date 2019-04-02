package chapter_4

import "fmt"

func main()  {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	close(intChan)
	syncChan := make(chan struct{}, 1)
	go func() {
		Loop:
			for {
				select {
				case elem, ok := <- intChan:
					if !ok {
						fmt.Println("End.")
						break Loop
					}
					fmt.Println("Received: ", elem)
				}
			}
		syncChan <- struct{}{}
	}()

	<- syncChan
}
