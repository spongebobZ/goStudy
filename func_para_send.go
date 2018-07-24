package main

import (
	"fmt"
)

func main() {
	/*a值传递示例,b引用传递示例*/
	a, b := 1, 2
	fmt.Printf("in the func a is %d,outside a is %d\n", change1(a), a)
	fmt.Printf("in the func b is %d,outside b is %d", change2(&b), b)

}
func change1(num int) int {
	num = 100
	return num
}
func change2(num *int) int {
	*num = 200
	return *num
}
