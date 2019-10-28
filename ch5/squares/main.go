package main

import "fmt"

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	s := squares()
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

/*
1
4
9
1
4
9
*/
