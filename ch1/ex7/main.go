package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			message := fmt.Sprintf("fetch: %v\n", err)
			return errors.New(message)
		}
		defer resp.Body.Close()

		fmt.Printf("%d: %s\n", resp.StatusCode, resp.Status)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			message := fmt.Sprintf("fetch: %v\n", err)
			return errors.New(message)
		}
	}
	return nil
}
