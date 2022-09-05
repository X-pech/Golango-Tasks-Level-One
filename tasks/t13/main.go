package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Int()
	b := rand.Int()

	fmt.Printf("%d %d\n", a, b)

	a += b    // a = a0 + b0, b = b0
	b = a - b // a = a0 + b0, b = a0 + b0 - b0 = a0
	a = a - b // b = a0, a = a0 + b0 - a0 = b0

	fmt.Printf("%d %d\n", a, b)

}
