package main

import "fmt"

type a struct {
	a int
	b int
}

func main() {
	x := a{a: 1}

	fmt.Printf("%d", x.a)
}
