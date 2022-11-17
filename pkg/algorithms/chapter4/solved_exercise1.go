package chapter4

import (
	"fmt"
	"sync/atomic"
)

func peakOfAList(xs []int, currentGoRoutines *uint64, maxGoRoutines uint64) int {
	if len(xs) == 3 {
		if xs[2] > xs[0] && xs[2] > xs[1] { //peak is last
			return xs[2]
		} else if xs[1] > xs[0] && xs[1] > xs[2] { //peak is second
			return xs[1]
		} else if xs[0] > xs[1] && xs[0] > xs[2] { //peak is first
			return xs[0]
		} else { //they are all equal so just return first
			return xs[0]
		}
	} else if len(xs) == 2 {
		if xs[0] > xs[1] {
			return xs[0]
		} else {
			return xs[1]
		}
	} else {
		//slice the array in half and send it off recursively
		i := len(xs) / 2
		left := xs[0:i]
		right := xs[i:]

		async := func(xs []int, currentGoRoutines *uint64, maxGoRoutines uint64, f func(xss []int, currentGoRoutines *uint64, maxGoRoutines uint64) int, c chan int) {
			go func() {
				actual := f(xs, currentGoRoutines, maxGoRoutines)
				c <- actual
			}()
		}
		if atomic.LoadUint64(currentGoRoutines) < maxGoRoutines {
			//fmt.Println("spawning a new go routine")
			atomic.AddUint64(currentGoRoutines, 1)
			//fmt.Printf("Current # go routines:%v\n", atomic.LoadUint64(currentGoRoutines))
			c := make(chan int, 2)
			async(left, currentGoRoutines, maxGoRoutines, peakOfAList, c)
			async(right, currentGoRoutines, maxGoRoutines, peakOfAList, c)

			diasResults := make([]int, 2)
			diasResults[0], diasResults[1] = <-c, <-c

			return peakOfAList([]int{diasResults[0], diasResults[1]}, currentGoRoutines, maxGoRoutines)
		} else {
			//fmt.Println("NOT spawning a new go routine")
			a := peakOfAList(left, currentGoRoutines, maxGoRoutines)
			b := peakOfAList(right, currentGoRoutines, maxGoRoutines)

			return peakOfAList([]int{a, b}, currentGoRoutines, maxGoRoutines)
		}
	}
}

func PeakOfAList(xs []int, maxGoRoutines uint64) int {
	var currentGoRoutines uint64
	r := peakOfAList(xs, &currentGoRoutines, maxGoRoutines)
	fmt.Printf("Go rountines: %v\n", currentGoRoutines)
	return r
}
