package main

import (
	"fmt"
)

// TODO need to be done.
func eliminate(x []string) {
	for i, l := 0, len(x)-1; i < len(x)-1; i++ {
		if x[i] == x[i+1] {
			copy(x[i:], x[i+1:])
			l--
			fmt.Println(x)
		}

	}
}

func main() {
	ages := map[string]int{}
	if age, ok := ages["nick"]; !ok { /* ... */
		println(age, ok)
	}
}
