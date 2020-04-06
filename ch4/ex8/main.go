package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	os.Exit(run())
}

func run() int {
	var letter, mark, num int
	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			return 1
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(r) {
			letter++
		}
		if unicode.IsMark(r) {
			mark++
		}
		if unicode.IsNumber(r) {
			num++
		}
	}

	fmt.Printf("rune\tletter\n")
	fmt.Printf("letter\t%d\n",letter)

	fmt.Printf("rune\tmark\n")
	fmt.Printf("mark\t%d\n",mark)

	fmt.Printf("rune\tnumber\n")
	fmt.Printf("number\t%d\n",num)

	return 0

}

func checkRune(r rune, judge func(r rune) bool, counter map[rune]int) {
	if judge(r) {
		counter[r]++
	}
}