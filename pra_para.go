package main

import "fmt"

func main() {
	/*指针与值的关系，a是一个指针，并且a引用了整数10的指针地址；b是一个整数，引用了指针a保存的值；d是一个指针，并且d引用了a的指针地址，当修改d指针保存的值时，a指针的值同样也被修改了*/
	c := 10
	var a *int
	a = &c
	b := *a
	d := a
	*d = 20

	fmt.Println(b, a, *a, d)
}
