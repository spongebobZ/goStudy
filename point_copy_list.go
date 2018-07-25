package main

import (
	"fmt"
)

func main() {
	birth := []int{1, 9, 9, 2, 0, 6, 0, 8}
	birth_len := len(birth)
	var ptrlist [8]*int
	for i := 0; i < birth_len; i++ {
		ptrlist[i] = &birth[birth_len-1-i]
	}
	ptr_len := len(ptrlist)
	for j := 0; j < ptr_len; j++ {
		fmt.Print(*ptrlist[j])
	}
}
