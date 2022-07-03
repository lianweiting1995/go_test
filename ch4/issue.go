package ch4

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch4/github"
)

func IssueFunc() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s", item.Number, item.User.Login, item.Title)
	}
}
