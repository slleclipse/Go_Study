package main

import (
	"awesomeProject/main/interface/two"
	"fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type File struct {
}
func (f *File) Read(buf []byte) (n int, err error) {
	return
}
func (f *File) Write(buf []byte) (n int, err error) {
	return
}
func (f *File) Seek(off int64, whence int) (pos int64, err error) {
	return
}
func (f *File) Close() error {
	return nil
}

type Writer interface {
	Write(buf []byte) (n int, err error)
}


func main() {
	var a Integer = 1
	var b LessAdder = &a
	b.Add(2)
	fmt.Println(a)

	/*var file1 two.IStream = new(File)
	var file2 one.ReadWriter = file1
	var file3 two.IStream = file2*/

	var file1 two.IStream = new(File)
	var file4 Writer = file1

	if file5, ok := file4.(*File); ok {
		fmt.Println("file5: ", file5)
	}

	var v1 interface{} = 1

	switch v := v1.(type) {
		case int:
			fmt.Println("v: ",v, " int ")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("nil type")
	}

	switch v := file1.(type) {
	case Writer: // 现在v的类型是int
		if v, ok := v.(Writer); ok { // 现在v的类型是Stringer
			fmt.Println("sdaf: ",v)
		} else {
			// ...
		}
	default:

	}
	//fmt.Println(file1, file4)

}
