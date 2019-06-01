package main

import (
	"html/template"
	"log"
	"os"
	"tourOfGolang/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	issuesList := template.Must(template.ParseFiles("./templ.tpl"))
	html, err := os.Create("issueslist.html")
	if err != nil {
		log.Fatal("Creat file failed: ", err)
	}
	err = issuesList.Execute(html, result)
	if err != nil {
		log.Fatal(err)
	}

}
