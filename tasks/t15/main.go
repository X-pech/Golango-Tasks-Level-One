package main

import (
	"fmt"
)

func createHugeString(length int) string {
	s := make([]rune, length)
	for i := range s {
		s[i] = 'X'
	}
	return string(s)
}

func sliceProof() {
	slice1 := make([]int, 20)
	slice2 := slice1[5:10]
	slice1[6] = 435
	fmt.Printf("%d == %d!!!", slice2[1], slice1[6])
}

/*
Maybe the issue here is that someFunc does not return value
But operates with the global value?
So we can't correctly use it as a goroutine
Because working with global value
Is not thread-safe at all
So...

And, by the way
Slicing is not copying values
*/

func someFunc() string {

	v := createHugeString(1 << 10)

	justSlice := make([]byte, 100)
	copy(justSlice, v)
	return string(justSlice)

}

func main() {

	sliceProof()

	justString := someFunc()
	fmt.Println(justString)

}
