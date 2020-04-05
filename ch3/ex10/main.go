package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", commaWithFloat(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaWithBuf(s string) string {
	var buf bytes.Buffer
	b := []byte(s)
	rest := len(b) % 3
	for i, k := range b {
		if (i-rest)%3 == 0 {
			buf.WriteRune(',')
		}

		buf.WriteByte(k)
	}

	return buf.String()
}

func commaWithFloat(s string) string {
	nums := s
	var floats string
	if strings.Contains(s, ".") {
		splited := strings.Split(s, ".")
		nums = splited[0]
		floats = splited[1]
	}

	s = commaWithBuf(nums)
	return s + "." + floats
}
