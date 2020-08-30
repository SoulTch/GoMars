package main

import "fmt"

type a struct {
	x int
}

type b struct {
	x int
	a
}

func main() {
	a := b{x: 3, a: a{4}
	fmt.Println(a.x)
}
