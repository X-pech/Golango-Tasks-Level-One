package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
	"wbtechl1tasks/taskutils"
)

func pivot[V any](arr []V, left int, right int) int {
	// Choosing pivot by random
	rand.Seed(time.Now().Unix())
	return left + rand.Intn(right-left)
}

type CompFunc[V any] func(V, V) int // Three-way comparator {-1, 0, 1}

func partition[V any](arr []V, comp CompFunc[V], left int, right int, pivotIndex int) (int, int) {
	pivotElement := arr[pivotIndex]

	// These two would be beginning and the end of the segment with comp == 0
	// pbegin would be incremented from the left
	// And pend would be decremented from the right
	// (let us increment it first for convinience)
	pbegin := left
	pend := 0

	i := left

	for {
		if i >= right {
			break
		}

		if comp(pivotElement, arr[i]) > 0 {
			pbegin++
		} else if comp(arr[i], pivotElement) > 0 {
			pend++
		}

		i++
	}

	pend = right - pend

	// Then lets make a partition
	li := left   // comp < 0
	ei := pbegin // comp == 0
	ri := pend   // comp > 0

	for {
		// while anything inside bound
		if !(li < pbegin || ri < right || ei < pend) {
			break
		}

		for {
			// while li is in bounds and comp < 0
			if !(li < pbegin && comp(arr[li], pivotElement) < 0) {
				break
			}
			li++
		}

		for {
			// while ei is in bounds and comp == 0
			if !(ei < pend && comp(arr[ei], pivotElement) == 0) {
				break
			}
			ei++
		}

		for {
			// while ri is in bounds and comp > 0
			if !(ri < right && comp(arr[ri], pivotElement) > 0) {
				break
			}
			ri++
		}

		// send equal elements to the center
		if li != pbegin && ei != pend && comp(arr[li], pivotElement) == 0 {
			arr[li], arr[ei] = arr[ei], arr[li]
			ei++
		}

		if ei != pend && ri != right && comp(arr[ri], pivotElement) == 0 {
			arr[ri], arr[ei] = arr[ei], arr[ri]
			ei++
		}

		// then swap them if the order is incorrect
		if li != pbegin && ri != right && comp(arr[li], arr[ri]) > 0 {
			arr[li], arr[ri] = arr[ri], arr[li]
		}
	}

	return pbegin, pend
}

func qsort[V any](arr []V, comp CompFunc[V], left int, right int, wg *sync.WaitGroup) {

	// To decrease a wg counter
	defer wg.Done()

	// Checking if we really need to sort it
	if right-left <= 1 {
		return
	}

	// Choosing a pivot
	pivotIndex := pivot(arr, left, right)

	// making a partition
	pbegin, pend := partition(arr, comp, left, right, pivotIndex)

	// concurrently runnning them
	wg.Add(2)
	go qsort(arr, comp, left, pbegin, wg)
	go qsort(arr, comp, pend, right, wg)
}

func QuickSort[V any](arr []V, comp CompFunc[V]) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go qsort(arr, comp, 0, len(arr), &wg)
	// it will wait until all the "leaf" recurrent goroutines stop
	wg.Wait()
}

// Additional checker to test it
func checkArray[V any](arr []V, comp CompFunc[V]) bool {
	n := len(arr)
	for i := 1; i < n; i++ {
		if comp(arr[i-1], arr[i]) > 0 {
			return false
		}
	}
	return true
}

func main() {
	c := make(chan []int)
	go taskutils.GenerateSlice(300000, math.MaxInt32, c)
	arr := <-c
	close(c)

	// taskutils.PrintSlice(arr)
	QuickSort(arr, func(l int, r int) int {
		return l - r
	})
	// taskutils.PrintSlice(arr)

	fmt.Printf("This is: %t\n", checkArray(arr, func(l int, r int) int {
		return l - r
	}))

}
