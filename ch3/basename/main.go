package main

// basename remove route and '.' suffix
// e.g. a=>a, a.b=>a, a/b.c=>b
import (
	"fmt"
	"strings"
)

func basename(s string) string {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// same function as basename.Use srings to deal this problem.
func basename1(s string) string {
	dot := strings.LastIndex(s, ".")
	slash := strings.LastIndex(s, "/")
	if dot == -1 {
		dot = len(s)
	}
	return s[slash+1 : dot]
}

func main() {
	fmt.Println(basename("a/b/v/b.a.c.go"))
	fmt.Println(basename1("bgo.md"))
}
