package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var buf bytes.Buffer
	for i := 0; i < 100; i++ {
		buf.WriteString("A")
	}
	fmt.Println(buf.String())
	//O(n)

	//faster - prevents accidental copying
	//can only grow or reset
	var sb strings.Builder
	for i := 0; i < 100; i++ {
		sb.WriteString("A")
	}
	fmt.Println(sb.String())
}
