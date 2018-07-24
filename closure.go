package main

import (
	"fmt"
)

func addnum() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	num := addnum() /*相当于num:=func() int {
		i += 1
		return i
	}，即addnum()的返回值，每调用一次num()，就相当于在不退出addnum()的前提下，执行一次return func()*/
	fmt.Println(num())
	fmt.Println(num())
	fmt.Println(num())
	num2 := addnum()
	fmt.Println(num2())
	fmt.Println(num2())
	fmt.Println(num2())
	num3 := addnum()
	fmt.Println(num3)
	fmt.Println(num3)
	fmt.Println(num3)
}
