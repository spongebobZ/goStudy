package main

import (
	"fmt"
)

func main() {
	birth := []int{1, 9, 9, 2, 0, 6, 0, 8}
	var birth_len int
	birth_len = len(birth)
	const (
		a = &birth_len
		b = a
	)
	var ptrlist [b]*int
	for i := 0; i < b; i++ {
		ptrlist[i] = &birth[b-1-i]
	}
	ptr_len := len(ptrlist)
	for j := 0; j < ptr_len; j++ {
		fmt.Println(*ptrlist[j])
	}
}
