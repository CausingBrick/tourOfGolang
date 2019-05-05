package main

// function compute the number of non-zero digits for sha256
import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

// init pc
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// popcout function returns the population count of [32]byte
func popcount(x [32]byte) int {
	var count int
	for i := uint(0); i < 32; i++ {
		count += int(pc[x[i]])
	}
	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	fmt.Printf("%b\n", c1)
	fmt.Println(popcount(c1))
}
