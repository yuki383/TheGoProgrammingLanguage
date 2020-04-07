package main

import (
	"fmt"
	"treesort"
)

func main() {
	values := []int{2, 3, 1, 6, 5, 9, 0, 7, 8}
	treesort.Sort(values)

	for _, v := range values {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")
}
