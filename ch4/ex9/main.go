package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
}

func run() error {
	count := make(map[string]int)
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		wordfreq(scanner, count)
	}

	for k, v := range count {
		fmt.Printf("%s\t%d\n", k, v)
	}

	return nil
}

func wordfreq(scanner *bufio.Scanner, m map[string]int) {
	for scanner.Scan() {
		t := scanner.Text()
		m[t]++
	}
}
