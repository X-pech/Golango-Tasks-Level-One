package main

import (
	"fmt"
	"math/rand"
	"time"
	"wbtechl1tasks/taskutils"
)

func eraseElement[V any](arr []V, i int) []V {
	// ... is for unpacking values
	// like * in Python
	// so we can pass every element from slice
	// as a distinct argument
	return append(arr[:i], arr[i+1:]...)
}

func main() {
	rand.Seed(time.Now().Unix())
	n := 20 + rand.Intn(30)
	c := make(chan []int)
	go taskutils.GenerateSlice(n, 100, c)
	arr := <-c
	close(c)

	fmt.Println(len(arr))
	taskutils.PrintSlice(arr)
	toDel := rand.Intn(n)
	fmt.Println(toDel)

	arr = eraseElement(arr, toDel)
	fmt.Println(len(arr))
	taskutils.PrintSlice(arr)
}
