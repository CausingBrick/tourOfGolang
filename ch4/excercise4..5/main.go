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
	// var x []string
	// for _, v := range "11111111112223333444555666777888999" {
	// 	x = append(x, string(v))
	// }
	// fmt.Println(x)
	// eliminate(x)
	// fmt.Println(x)
	ages := map[string]int{}
	if age, ok := ages["nick"]; !ok { /* ... */
		println(age, ok)
	}
}
