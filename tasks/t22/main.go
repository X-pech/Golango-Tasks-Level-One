package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	/*
		(1<<20) * (1<<20) > 32 => int is not suitable there
		by the way we can not see condition "less than...",
		only "greater than (1 << 20)" so let's just use
		big arithmetics module
	*/
	rand.Seed(time.Now().Unix())
	a := big.NewInt(1 << 20 * rand.Int63n(1<<20))
	b := big.NewInt(1 << 20 * rand.Int63n(1<<20))
	c := big.NewInt(0)
	template := "%v %c %v = %v\n"
	fmt.Printf(template, a, '+', b, c.Add(a, b))
	fmt.Printf(template, a, '-', b, c.Sub(a, b))
	fmt.Printf(template, a, '*', b, c.Mul(a, b))
	fmt.Printf(template, a, '/', b, c.Div(a, b))

}
