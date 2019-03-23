package main

import "fmt"

func main() {
	test(5, 0)
}
func test(x, y int) {
	z := 0

	func() {
		defer func() {
			if err := recover(); err != nil {
				z = 0
			}
		}()
		z = x / y
	}()

	array := [...]int{1,2,5}

	var array2 [3]int
	array2 = array

	/*slice := [3][]int{{1,2,3},{2},{4}}
	fmt.Println(slice[1])
	slice[1] = append(slice[1], 100,11)

	fmt.Println(slice[1])*/

	fmt.Printf("array: %p, %v\n", &array, array)
	fmt.Printf("array2: %p, %v\n", &array2, array2)
	fmt.Println("x / y = ", z)

	slice := []int{1,4,5}
	test2(array)
	fmt.Printf("pointer before:%p, %v\n", &slice, slice)
	test3(slice)
	fmt.Printf("pointer after:%p, %v\n", &slice, slice)
}

func test2(array [3]int) {
	fmt.Printf("pointer:%p, %v\n", &array, array)
}

func test3(slice []int) {
	//slice = append(slice, 7)
	slice[2] = 10
	fmt.Printf("pointer:%p, %v\n", &slice, slice)
}

