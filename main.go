package main

import "fmt"

type a struct {
	a int
	b int
}

func main() {
	x := make([]*a, 5, 5)

	fmt.Printf("%d", x[0].a)
}
