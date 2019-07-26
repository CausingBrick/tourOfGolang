package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(0)
	}
	fmt.Println(outline(nil, doc))
}

func outline(stack []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
	return stack
}
