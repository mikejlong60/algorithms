package chapter4

import (
	"fmt"
	"sync/atomic"
)

func async[A any](xs []A, currentGoRoutines *uint64, maxGoRoutines uint64, gt func(a, b A) bool, f func([]A, *uint64, uint64, func(A, A) bool) A, c chan A) {
	go func() {
		actual := f(xs, currentGoRoutines, maxGoRoutines, gt)
		c <- actual
	}()
}

func peakOfAList[A any](xs []A, currentGoRoutines *uint64, maxGoRoutines uint64, gt func(a, b A) bool) A {
	if len(xs) == 3 {
		if gt(xs[2], xs[0]) && gt(xs[2], xs[1]) { //peak is last
			return xs[2]
		} else if gt(xs[1], xs[0]) && gt(xs[1], xs[2]) { //peak is second
			return xs[1]
		} else if gt(xs[0], xs[1]) && gt(xs[0], xs[2]) { //peak is first
			return xs[0]
		} else { //they are all equal so just return first
			return xs[0]
		}
	} else if len(xs) == 2 {
		if gt(xs[0], xs[1]) {
			return xs[0]
		} else {
			return xs[1]
		}
	} else {
		//slice the array in half and send it off recursively
		i := len(xs) / 2
		left := xs[0:i]
		right := xs[i:]

		if atomic.LoadUint64(currentGoRoutines) < maxGoRoutines {
			//fmt.Println("spawning a new go routine")
			atomic.AddUint64(currentGoRoutines, 1)
			//fmt.Printf("Current # go routines:%v\n", atomic.LoadUint64(currentGoRoutines))
			c := make(chan A, 2)

			f := peakOfAList[A]
			async(left, currentGoRoutines, maxGoRoutines, gt, f, c)
			async(right, currentGoRoutines, maxGoRoutines, gt, f, c)

			diasResults := make([]A, 2)
			diasResults[0], diasResults[1] = <-c, <-c

			return peakOfAList([]A{diasResults[0], diasResults[1]}, currentGoRoutines, maxGoRoutines, gt)
		} else {
			//fmt.Println("NOT spawning a new go routine")
			a := peakOfAList(left, currentGoRoutines, maxGoRoutines, gt)
			b := peakOfAList(right, currentGoRoutines, maxGoRoutines, gt)

			return peakOfAList([]A{a, b}, currentGoRoutines, maxGoRoutines, gt)
		}
	}
}

func PeakOfAList[A any](xs []A, maxGoRoutines uint64, gt func(a, b A) bool) A {
	var currentGoRoutines uint64
	r := peakOfAList(xs, &currentGoRoutines, maxGoRoutines, gt)
	fmt.Printf("Go rountines: %v\n", currentGoRoutines)
	return r
}
