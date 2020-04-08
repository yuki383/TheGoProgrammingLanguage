package main

import (
	"fmt"
	"issues"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	result, err := issues.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	t := ByTime(result.Items)
	sort.Sort(t)

	m,_ := t.findFirst(compareByTime(time.Hour * 24 * 30))
	y,_ := t.findFirst(compareByTime(time.Hour * 24 * 365))

	ltMonth := t[:m]
	ltYear := t[m:y]
	other := t[y:]

	fmt.Println("less than one month (30 days)")
	for _, item := range ltMonth {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.String(), item.User.Login, item.Title)
	}

	fmt.Println("\nless than one year (365 days)--------------------------")
	for _, item := range ltYear {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.String(), item.User.Login, item.Title)
	}

	fmt.Println("\nmore--------------------------")
	for _, item := range other {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.String(), item.User.Login, item.Title)
	}
}

type ByTime []*issues.Issue

func (b ByTime) Len() int           { return len(b) }
func (b ByTime) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByTime) Less(i, j int) bool { return b[i].CreatedAt.After(b[j].CreatedAt) }

func (b ByTime) findFirst(fn func(issue *issues.Issue) bool) (int, bool) {
	for i, v := range b {
		if fn(v) {
			return i, true
		}
	}

	return 0, false

}

func compareByTime(d time.Duration) func(i *issues.Issue) bool {
	return func(i *issues.Issue) bool {
		if time.Now().After(i.CreatedAt.Add(d)) {
			return true
		}
		return false
	}
}
