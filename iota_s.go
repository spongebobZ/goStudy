package main

import "fmt"

func main() {
	const (
		a = 1 << iota
		b = 2 << iota
		c
		d
	)
	fmt.Println(a, b, c, d)
}
