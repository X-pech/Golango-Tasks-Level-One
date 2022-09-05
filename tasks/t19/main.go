package main

import (
	"fmt"
	"wbtechl1tasks/taskutils"
)

func reversed(s *string) string {

	/*
		first of all, strings are value-type
		you can not change random character
		second of all, strings are []byte
		so if we reverse the order of string
		elements we can get invalid
		characters because of unicode
		so we need to convert it into a rune slice
		and then reverse the order of the runes
	*/
	res := []rune(*s)
	// n := len(res)

	// for i := 0; i < n/2; i++ {
	// 	res[i], res[n-i-1] = res[n-i-1], res[i]
	// }
	taskutils.ReverseSlice(res)
	return string(res)
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(reversed(&s))
}
