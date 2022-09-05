package main

import (
	"fmt"
	"time"
)

/*
I checked the implementation of time.Sleep
And of course this is very low-level
But it is implemented by something like
"low-level" timer so let me implement it
with high-level timer
*/
func GigaSleep(d time.Duration) {
	t := time.NewTimer(d)
	<-t.C
}

func main() {
	fmt.Println("Hi there")
	GigaSleep(time.Second * 2)
	fmt.Println("Goodbye there")
}
