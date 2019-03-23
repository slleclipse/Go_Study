package main

import "fmt"

func main() {

	stack := make([]int, 0, 10)

	push := func(x int) {
		length := len(stack)
		if cap(stack) == length {
			fmt.Println("Stack is full!")
		}
		stack = stack[:length + 1]
		stack[length] = x
	}

	pop := func() int{
		length := len(stack)
		if length == 0 {
			fmt.Println("Stack is empty!")
		}
		n := stack[length - 1]
		stack = stack[:length - 1]
		return n
	}

	push(1)

	fmt.Println(pop())
}
