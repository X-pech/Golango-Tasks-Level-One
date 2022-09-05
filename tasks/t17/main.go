package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

/*
Function itself, classic lower_bound implementation on generics
*/
func lwbnd[V constraints.Ordered](arr []V, value V, left int, right int) int {
	if right-left <= 1 {
		return left
	}

	m := (left + right) / 2
	if arr[m] >= value {
		return lwbnd(arr, value, left, m)
	} else {
		return lwbnd(arr, value, m, right)
	}
}

/*
Starting point for function
*/
func LowerBound[V constraints.Ordered](arr []V, value V) int {
	return lwbnd(arr, value, 0, len(arr))
}

/*
Function itself, classic upper_bound implementation on generics
*/
func upbnd[V constraints.Ordered](arr []V, value V, left int, right int) int {
	if right-left <= 1 {
		return right
	}

	m := (left + right) / 2
	if arr[m] > value {
		return upbnd(arr, value, left, m)
	} else {
		return upbnd(arr, value, m, right)
	}
}

/*
Starting point for function
*/
func UpperBound[V constraints.Ordered](arr []V, value V) int {
	return upbnd(arr, value, -1, len(arr)-1)
}

/*
versions with custom comp(l V, r V) int function which should return
 -1 if l should be before r
  0 if l is replaceable with r
  1 if l should be after r
*/
func lwbndcmp[V any](arr []V, value V, comp func(left V, right V) int, left int, right int) int {
	if right-left <= 1 {
		return left
	}

	m := (left + right) / 2
	res := comp(value, arr[m])
	if res <= 0 {
		return lwbndcmp(arr, value, comp, left, m)
	} else {
		return lwbndcmp(arr, value, comp, m, right)
	}
}

func LowerBoundWithComp[V any](arr []V, value V, comp func(left V, right V) int) int {
	return lwbndcmp(arr, value, comp, 0, len(arr))
}

func upbndcmp[V any](arr []V, value V, comp func(left V, right V) int, left int, right int) int {
	if right-left <= 1 {
		return right
	}

	m := (left + right) / 2
	res := comp(value, arr[m])
	if res < 0 {
		return upbndcmp(arr, value, comp, left, m)
	} else {
		return upbndcmp(arr, value, comp, m, right)
	}
}

func UpperBoundWithComp[V any](arr []V, value V, comp func(left V, right V) int) int {
	return upbndcmp(arr, value, comp, -1, len(arr)-1)
}

func main() {
	arr := []int{1, 2, 5, 10, 22, 31, 233, 1202}
	fmt.Printf("LB(5): %d\n", LowerBound(arr, 5))
	fmt.Printf("LB(-5): %d\n", LowerBound(arr, -5))
	fmt.Printf("LB(2222): %d\n", LowerBound(arr, 2222))

	fmt.Printf("UB(5): %d\n", UpperBound(arr, 5))
	fmt.Printf("UB(1202): %d\n", UpperBound(arr, 1202))
	fmt.Printf("UB(-5): %d\n", UpperBound(arr, -5))

	cmp := func(l int, r int) int {
		return l - r
	}

	fmt.Printf("UBC(-5): %d\n", UpperBoundWithComp(arr, -5, cmp))
	fmt.Printf("UBC(5): %d\n", UpperBoundWithComp(arr, 5, cmp))
	fmt.Printf("LBC(5): %d\n", LowerBoundWithComp(arr, 5, cmp))
	fmt.Printf("LBC(2222): %d\n", LowerBoundWithComp(arr, 2222, cmp))

}
