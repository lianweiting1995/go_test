package ch4

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

const TIMELAYOUT = "2006-01-02 15:04:05"

func IssueFunc() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	localT, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)
	pDate := "2021-12-12 12:00:06"
	pTime, err := time.ParseInLocation(TIMELAYOUT, pDate, localT)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Items {
		createdDate := item.CreatedAt.UTC().Format(TIMELAYOUT)
		createdTime, err := time.ParseInLocation(TIMELAYOUT, createdDate, localT)
		if err != nil {
			log.Fatal(err)
		}
		result := createdTime.Before(pTime)
		fmt.Printf("#%-5d %s %s %s %s %t\n", item.Number, item.User.Login, item.Title, createdDate, pDate, result)
	}
}
