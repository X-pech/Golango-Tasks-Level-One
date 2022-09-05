package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Не уверен нуждается ли это в комментировании, честно говоря
*/
func multiplierRoutine(from <-chan int, to chan<- int) {
	for {
		to <- (<-from * 2)
	}
}

func printerRoutine(from <-chan int) {
	for {
		fmt.Printf("%d ", <-from)
	}
}

func main() {
	ch1 := make(chan int, 20)
	ch2 := make(chan int, 20)

	go multiplierRoutine(ch1, ch2)
	go printerRoutine(ch2)

	for i := 0; i < 20; i++ {
		ch1 <- rand.Int()
		time.Sleep(1 * time.Microsecond)
	}
}
