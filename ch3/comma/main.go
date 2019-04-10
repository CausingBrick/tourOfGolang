package main

import "fmt"

func comma(s string) string {
	l := len(s)
	if l < 4 {
		return s
	}
	return comma(s[:l-3]) + "," + s[l-3:]
}

func main() {
	fmt.Println(comma("1234546546546546"))
}
