package main

import "fmt"

func main() {
	/*
		set and map are always implemented
		in the same ways
		with hashtables or with trees

		golang has no sets so let's implement
		set on the base of a map

		we will emphasize the fact it is set
		by making a value of type struct{}
		that is a classic way for golang to say
		"we need the fact of existence, not the value"
		(nil is not suitable for the fact of "existence"
		and bool is a value itself)

	*/
	set := make(map[string]struct{})
	words := []string{"cat", "cat", "dog", "tree", "cat"}
	for i := range words {
		set[words[i]] = struct{}{}
	}

	for word := range set {
		fmt.Printf("%s ", word)
	}
	fmt.Println()
}
