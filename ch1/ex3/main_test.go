package main

import (
	"strings"
	"testing"
)

func BenchmarkSlower(b *testing.B) {
	data := strings.Split("Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad mini", " ")
	for i := 0; i < b.N; i++ {
		Slower(data)
	}
}

func BenchmarkFaster(b *testing.B) {
	data := strings.Split("Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad mini", " ")
	for i := 0; i < b.N; i++ {
		strings.Join(data, " ")
	}
}
