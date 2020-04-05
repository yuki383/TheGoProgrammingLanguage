package main

import "fmt"

// 
func main() {
	s := []string{ "hoge", "fuga", "fuga", "foo"}
	result := removeRelated(s)

	fmt.Printf("%v\n", result)

}

func removeRelated(s []string) []string {
	var current int

	for i := 0; i < len(s)-1; i++ {
		if s[current] != s[i+1] {
			s[current+1] = s[i+1]
			current++
			continue
		}
	}
	return s[:current+1]

}