package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
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
	rest := len(s) % 3
	for i := 0; i < len(s); i++ {
		if (i-rest)%3 == 0 {
			buf.WriteRune(',')
		}
		fmt.Fprintf(&buf, "%s", string(s[i]))
	}

	return buf.String()
}
