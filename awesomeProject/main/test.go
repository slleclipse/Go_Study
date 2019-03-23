package main

import "fmt"

func main(){
	var c int = 0
	/*const (
		c0 = iota
		c1 = iota
		c2 = iota
	)*/
	/*const (
		c0 = iota
		c1
		c2
	)*/
	/*const (
		c0 = 1 << iota
		c1 = 1 << iota
		c2 = 3
		c3 = 1 << iota
	)*/

	//const usage
	const (
		c0 = iota * 42
		c1 float64 = iota * 42
		c2 = iota * 42
		c3 = 1 << iota
	)

	//exchange value
	a := 0
	b := 1
	a, b = b, a
	fmt.Println("c0: ", c0, "c1: ", c1, "c2: ", c2, "c3: ", c3)
	fmt.Println("hello world:" , "a = ",a ,"b = ", b, "c = ", c)

	ok := true
	var no bool
	if a != b {
		fmt.Println("ok: ", ok)
	} else {
		fmt.Println("no: ", no)
	}

	//return multi value
	firstName, _, _ := getName("May")
	fmt.Println("firstName", "Hello " + firstName)
	fmt.Println("length:", len(firstName), "is %c",firstName[0])

	//for usage
	for i:=0; i < len(firstName); i++ {
		ch := firstName[i]
		fmt.Println(i, ch)
	}
	for i, ch1 := range firstName {
		fmt.Println(i, ch1)
	}

	//array slice
	var array [3]int = [3]int{3, 4, 5}
	//newArray := [3]int{3,6,8}
	anotherSlice := make([]int, 3)
	for _, num := range anotherSlice {
		fmt.Println( "another num:", num)
	}

	var mySlice []int = array[:]
	mySlice = append(mySlice,6, 7, 8)
	for _, num := range mySlice {
		fmt.Println( "num", num)
	}
	fmt.Println("len", len(mySlice))
	fmt.Println("cap", cap(mySlice))

	var newSlice []int = make([]int,5)
	for _, num := range newSlice {
		fmt.Println( "num", num)
	}
	fmt.Println("len", len(newSlice))
	fmt.Println("cap", cap(newSlice))

	//map usage
	type Person struct {
		id string
		name string
		address string
	}

	var personMap map[string] Person = make(map[string] Person)
	personMap["1"] = Person{"1", "long", "hangqiao"}

	person := Person{"2", "xiaohui", "fudian"}
	personMap[person.id] = person

	delete(personMap, "3")
	person1, ok := personMap["2"]
	fmt.Println("map len: ", len(personMap))
	fmt.Println("person1: ", person1.id, "ok: ", ok)

	//function
	j := 5
	d := func () func(){
		i := 10
		return func() {
			fmt.Println("j: ", j, "i: ", i)
		}
	}()
	d()
	j *= 2
	d()

	//interface
	var a Integer

}

func getName(first string) (firstName, lastName, nikeName string) {
	firstName = first
	return firstName, "Chan", "Chibi Maruko"
}

type PathError struct {
	Op string
	Path string
	err error
}

func (err *PathError) Error() string {
	return err.Op + err.Path + err.err.Error();
}

func Stat(name string) (fi FileInfo, err error) {

}
/*func Foo(param int) (n int, err error) {
	var err error =
	return param,err
}*/