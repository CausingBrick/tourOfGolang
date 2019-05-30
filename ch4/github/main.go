package main

import (
	"fmt"
	"log"
	"os"
	"tourOfGolang/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Issues count:", result.TotalCount)
	for _, val := range result.Items {
		fmt.Println(val)
	}
}
