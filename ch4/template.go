package ch4

import (
	"log"
	"os"
	"text/template"
	"time"

	"gopl.io/ch4/github"
)

func InitTemplate() {
	var templ string = `{{.TotalCount}} issues
{{range.Items}}
-----------------------------
Number:	{{.Number}}
User:	{{.User.Login}}
Title:	{{.Title | printf "%.64s"}}
Age:	{{.CreatedAt | daysAge}} days
{{end}}`

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	resport, err := template.New("issue").Funcs(template.FuncMap{
		"daysAge": daysAge,
	}).Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	if err = resport.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAge(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
