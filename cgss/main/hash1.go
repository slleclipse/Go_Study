package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	TestString := "Hi, pandaman!"

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	Result := Md5Inst.Sum([]byte(""))

	fmt.Printf("%x \n\n", Result)

	ShalInst := sha1.New()
	ShalInst.Write([]byte(TestString))
	Result = ShalInst.Sum([]byte(""))
	fmt.Printf("%x \n\n", Result)

}
