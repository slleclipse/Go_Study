package main

import (
	"errors"
	"fmt"
)

func main() {

	funcA()
}

func funcA() (err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic recover! p:",p)
			str, ok := p.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic")
			}
		}
	}()

	return funcB()
}

func funcB() error {
	panic("hhhh")
}