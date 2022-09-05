package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func sb() {
	s := new(strings.Builder)
	s.WriteRune('a')
	fmt.Println(s.String())

}

func emptystruct() {
	fmt.Println(unsafe.Sizeof(struct{ bool }{}))

}

func mapiter() {
	m := make(map[int]int)
	m[1] = 0
	m[2] = 3543
	m[5] = 2184909
	m[-1123] = -123

	for i := range m {
		fmt.Println(i)
	}
}

func makenew() {
	var arr = []int{2, 2, 2}
	arr[2] = 2

	var arr2 []int
	arr2 = append(arr2, 2)
	arr2[0] = 23

	arr3 := make([]int, 123, 222)
	arr3[0] = 123

	var dict2 = map[int]int{
		2: 3,
		3: 4,
	}
	dict2[5] = 6

	dict3 := make(map[int]int)
	dict3[2] = 3

	dict4 := make(map[int]int, 2)
	dict4[2] = 0
}

func update(p *int) {

	b := 2

	p = &b

}

func t10() {

	var (
		a = 1

		p = &a
	)

	fmt.Println(*p)

	update(p)

	fmt.Println(*p)

}

func someAction(v []int8, b int8) {
	v[0] = 100

	v = append(v, b)

}

func t13() {
	var a = []int8{1, 2, 3, 4, 5}

	someAction(a, 6)

	fmt.Println(a)
}

func t14() {

	slice := []string{"a", "a"}

	func(slice []string) {

		fmt.Printf("%p\n", slice)
		slice = append(slice, "a")
		fmt.Printf("%p\n", slice)
		slice = append(slice, "a")
		fmt.Printf("%p\n", slice)

		slice[0] = "b"

		slice[1] = "b"

		fmt.Println(slice)

	}(slice)

	fmt.Printf("%p\n", slice)
	fmt.Println(slice)
}

func main() {
	t14()
}
