package main

import (
	"fmt"
	"unsafe"
)

type emptyStruct struct {
	A struct{}
	B struct{}
	C struct{}
	D struct{}
	E struct{}
}

func main() {
	var empty emptyStruct
	fmt.Println(unsafe.Sizeof(struct{}{})) //0
	fmt.Println(unsafe.Sizeof(empty))      //0
}
