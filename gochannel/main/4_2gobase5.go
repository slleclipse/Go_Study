package main

import (
	"fmt"
	"time"
)

func main() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		/*go func(who string) {
			fmt.Println("Hello ", who)
		}(name)*/
		go func() {
			fmt.Println("Hello ", name)
		}()
	}
	time.Sleep(time.Millisecond)
}