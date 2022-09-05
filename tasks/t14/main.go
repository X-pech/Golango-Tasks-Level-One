package main

import (
	"fmt"
	"reflect"
)

func whoAreYou(value interface{}) string {
	/*
		reflect is a module for so-called "reflection"
		it has a method TypeOf that returns type of a value
		and Type type has a String() method for string
		represenation of this type signature
	*/
	return reflect.TypeOf(value).String()
}

func main() {
	fmt.Println(whoAreYou(5))
	fmt.Println(whoAreYou(5.0))
	fmt.Println(whoAreYou(false))
	fmt.Println(whoAreYou("Ayanami Rei"))
	fmt.Println(whoAreYou([]int{2, 2, 8}))
	fmt.Println(whoAreYou(map[int]string{5: "five"}))

}
