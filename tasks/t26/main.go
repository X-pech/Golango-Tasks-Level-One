package main

import (
	"fmt"
	"strings"
	"wbtechl1tasks/taskutils"
)

func checkStringAsSet(s string) bool {
	// runes := []rune(strings.ToLower(s))
	runes := strings.ToLower(s)
	set := make(map[rune]struct{})
	for _, r := range runes {
		if _, ok := set[r]; ok {
			return false
		}
		set[r] = struct{}{}
	}
	return true
}

func main() {
	c := make(chan string)
	go taskutils.GenerateString(20, c)
	s := <-c
	close(c)

	fmt.Println("abchsdkw")
	fmt.Println(checkStringAsSet("abchsdkw"))

	fmt.Println("abB")
	fmt.Println(checkStringAsSet("abB"))

	fmt.Println(s)
	fmt.Println(checkStringAsSet(s))
}
