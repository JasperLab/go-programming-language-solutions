package main

import (
	"fmt"
	"gopl/ch4/4.10/github"
	"log"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	args = append(args, "sort:created")
	//args = append(args, "order=asc")
	result, err := github.SearchIssues(args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	now := time.Now()
	month_ago := now.AddDate(0, -1, 0)
	year_ago := now.AddDate(-1, 0, 0)
	new_count, old_count, oldest_count := 0, 0, 0
	for _, item := range result.Items {
		if item.CreatedAt.After(month_ago) {
			if new_count == 0 {
				fmt.Println("\n== Less than a month old")
			}
			new_count++
		} else if item.CreatedAt.After(year_ago) {
			if old_count == 0 {
				fmt.Println("\n== Less than a year old")
			}
			old_count++
		} else {
			if oldest_count == 0 {
				fmt.Println("\n== More than a year old")
			}
			oldest_count++
		}
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number,
			item.User.Login, item.Title)
	}
}
