package main

import "fmt"

func main()  {

	instance := map[string] struct{
		x int
	} {
		"jj": {1},
		"hh": {3},
	}

	element, ok:= instance["hh"]
	if ok {
		fmt.Println(element.x)
	}

	instance1 := map[string] int {"a": 1}
	instance1["a"] = 0
	fmt.Println(instance1["a"])
}
