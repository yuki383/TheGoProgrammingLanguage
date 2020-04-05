package main

import (
	"fmt"
	"os"
)

func main() {
	anagram := isAnagram(os.Args[1], os.Args[2])
	var not string
	fmt.Printf("Args: %s, %s\n", os.Args[1], os.Args[2])
	if !anagram {
		not = "not"
	}
	fmt.Printf("%s and %s is %s anagram\n", os.Args[1], os.Args[2], not)
}

func isAnagram(origin, compare string) bool {
	om := recordRunes(origin)
	cm := recordRunes(compare)

	isSame := true

	for k, v := range om {
		if cm[k] != v {
			isSame = false
		}
	}

	return isSame
}

func recordRunes(s string) map[byte]int {
	m := make(map[byte]int)
	b := []byte(s)

	for _, v := range b {
		m[v]++
	}

	return m

}
