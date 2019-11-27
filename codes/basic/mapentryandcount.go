package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Fields(s)
	for _, v := range a {
		m[v]++
	}
	return m
}
func main() {
	fmt.Println(wordCount("I am I"))
}
