package main

import "fmt"

type a struct {
	a int
	b int
}

func main() {
	a := func() {}
	b := a
	fmt.Printf("%d %d", &a, &b)
}
