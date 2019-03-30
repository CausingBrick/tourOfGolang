package main

import (
	"fmt"
	exercise "tourOfGolang/ch2/exercise"
	popcount "tourOfGolang/ch2/popCount"
)

const num = 89

func main() {
	fmt.Println(num, exercise.PopCount(num), popcount.PopCount(num), exercise.PopCountEB(num))
}
