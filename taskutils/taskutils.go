package taskutils

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func ReverseSlice[V any](arr []V) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}

func GenerateStrings(count int, lengths int, c chan<- string) {
	for i := 0; i < count; i++ {
		go GenerateString(lengths, c)
	}
}

func GenerateString(length int, c chan<- string) {
	rand.Seed(time.Now().UnixNano())
	str := make([]rune, length)
	for j := range str {
		str[j] = letters[rand.Intn(len(letters))]
	}
	c <- string(str)
}

func GenerateSlice(length int, maxv int, c chan<- []int) {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, length)
	for j := range arr {
		arr[j] = rand.Intn(maxv)
	}
	c <- arr
}

func ReadSlice(arr *[]int) {
	var n int
	fmt.Scan(&n)
	var b int
	for i := 0; i < n; i++ {
		fmt.Scan(&b)
		*arr = append(*arr, b)
	}
}

func PrintSlice(arr []int) {
	for i := range arr {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}
