package main

import (
	"log"
	"os"
	"text/template"
	"time"
	"tourOfGolang/ch4/github"
)

// !+template
const templ = `
{{.TotalCount}} issues:
{{range .Items}}--------------------
Number:{{.Number}}
User:  {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age:   {{.CreatedAt | daysAgo}} hours ago
{{end}}
`

// !-template

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours())
}
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	err = report.Execute(os.Stdout, result)
	if err != nil {
		log.Fatal(err)
	}

	// must()
}

//
func must() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	report := template.Must(template.
		New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))

	err = report.Execute(os.Stdout, result)
	if err != nil {
		log.Fatal(err)
	}
}

/*
// !+ textoutput
$ go run main.go repo:shadowsocks/shadowsocks/

1422 issues:
--------------------
Number:1460
User:  SunR54
Title: SS端口和SSH端口有什么联系吗
Age:   2019-05-29T15:30:35Z
--------------------
Number:1459
User:  uebb
Title: 为什么SSR成功连接 没办法打开youtube
Age:   2019-05-28T08:22:30Z
--------------------
Number:1458
User:  walkingaway
Title: 免费ss账号，长期更新
Age:   2019-05-27T01:12:07Z
--------------------
Number:1457
User:  xiaotana
Title: super Wingy
Age:   2019-05-21T15:08:19Z
--------------------
...
// !- textoutput
*/
