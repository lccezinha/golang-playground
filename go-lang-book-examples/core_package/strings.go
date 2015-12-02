package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		// true
		strings.Contains("test", "t"),

		// 2
		strings.Count("test", "t"),

		// true
		strings.HasPrefix("test", "te"),

		// true
		strings.HasSuffix("test", "st"),

		// 1
		strings.Index("test", "e"),

		// a-b
		strings.Join([]string{"a", "b"}, "-"),

		// == "aaaa"
		strings.Repeat("a", 4),

		// bbaa
		strings.Replace("aaaa", "a", "b", 2),

		// xxxa
		strings.Replace("aaaa", "a", "x", 3),

		// []string{"a", "b", "c", "d", "e"}
		strings.Split("a-b-c-d-e", "-"),

		// test
		strings.ToLower("TEST"),

		// TEST
		strings.ToUpper("test"),
	)
}