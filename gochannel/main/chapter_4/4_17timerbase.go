package chapter_4

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(time.Second * 2)
	fmt.Printf("Present time: %v.\n", time.Now())
	expirationTime := <- timer.C
	fmt.Printf("Expiration time: %v.\n", expirationTime)
	fmt.Printf("Stop timer: %v.\n", timer.Stop())
}
