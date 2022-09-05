package main

import "wbtechl1tasks/taskutils"

func main() {

	n := 30

	c := make(chan []int)
	go taskutils.GenerateSlice(n, n*2, c)
	arr1 := <-c
	taskutils.PrintSlice(arr1)
	go taskutils.GenerateSlice(n, n*2, c)
	arr2 := <-c
	taskutils.PrintSlice(arr2)

	close(c)

	/*
		Creating set-like structure,
		adding every element from first slice to it
	*/
	map1 := make(map[int]struct{}, n)

	for i := range arr1 {
		map1[arr1[i]] = struct{}{}
	}

	res := make([]int, 0)

	/*
		Checking for every element from second slice
		if it is presented in set. If it is, adding it
		to result slice
	*/
	for i := range arr2 {
		if _, status := map1[arr2[i]]; status {
			res = append(res, arr2[i])
		}
	}

	taskutils.PrintSlice(res)

}
