package main

import (
	"fmt"
	"time"
)

func main() {
	var strChan = make(chan string, 3)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go receive1(strChan, syncChan1, syncChan2)
	go send1(strChan, syncChan1, syncChan2)
	<-syncChan2
	<-syncChan2
}

func receive1(strChan <-chan string, syncChan1 <-chan struct{}, syncChan2 chan<- struct{}) {
	<- syncChan1
	fmt.Println("Received a sync signal and wait a second...[reciever]")
	time.Sleep(time.Second)

	for elem := range strChan{
		fmt.Println("Received: ", elem, "[receiver]")
	}
	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
}

/*func send(strChan chan<- string, syncChan1 chan<- struct{}, syncChan2 chan<- struct{}) {
	for i := 0; i < 3; i++ {
		strChan <- strconv.Itoa(i)
		fmt.Println("Send:", i, "[sender]")
	}
	close(strChan)
	syncChan1 <- struct {}{}
	syncChan2 <- struct{}{}
}*/

func send1(strChan chan<- string, syncChan1 chan<- struct{}, syncChan2 chan<- struct{}) {

	for _, elem := range []string{"a", "b", "c", "d"}{
		strChan <- elem
		fmt.Println("Send:", elem, "[sender]")
		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("Sent a sync siganl. [sender]")
		}
	}
	fmt.Println("Wait 2 seconds...[sender]")
	time.Sleep(time.Second * 2)
	close(strChan)
	syncChan2 <- struct{}{}
}