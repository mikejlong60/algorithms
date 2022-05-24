package chapter2

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sorting"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

var lt = func(l, r int) bool {
	if l < r {
		return true
	} else {
		return false
	}
}

var parentIsLess = func(heap []int, lastIdx int) error {
	var parentIdx = ParentIdx(lastIdx)
	var childIdx = lastIdx
	var errors error
	for parentIdx > 0 {
		//		fmt.Printf("parents value:%v child's value:%v\n", heap[parentIdx], heap[childIdx])
		if heap[parentIdx] > heap[childIdx] {
			errors = multierror.Append(errors, fmt.Errorf("parent:%v value was not less than child's:%v\n", heap[parentIdx], heap[childIdx]))
		}
		childIdx = parentIdx
		parentIdx = ParentIdx(childIdx)
	}
	return errors
}

func TestHeapifyUpOddNumberBottomNodes(t *testing.T) {
	xss0 := []int{1, 2, 4, 5, 10, 9, 7, 11, 15, 17, 20, 17, 15, 8, 16, 1}
	xss1 := []int{0, 2, 4, 5, 10, 9, 7, 11, 15, 17, 20, 17, 15, 8, 3, 16}
	xss2 := []int{0, 2, 4, 5, 10, 9, 7, 11, 15, 17, 20, 17, 15, 3, 8, 16}
	xss3 := []int{0, 2, 4, 5, 10, 9, 7, 11, 15, 17, 20, 17, 3, 8, 16, 15}
	xss4 := []int{0, 2, 4, 5, 10, 9, 7, 11, 15, 17, 20, 3, 15, 8, 16, 17}
	xss5 := []int{0, 2, 4, 5, 10, 9, 7, 11, 15, 17, 3, 17, 15, 8, 16, 20}
	xss6 := []int{0, 2, 4, 5, 10, 9, 7, 11, 15, 3, 20, 17, 15, 8, 16, 17}
	xss7 := []int{0, 2, 4, 5, 10, 9, 7, 11, 3, 17, 20, 17, 15, 8, 16, 15}

	var elem = 15 //The element you want to add to the almost-a-heap. It's always 3 but position shifts
	var errors error
	for _, xss := range [][]int{xss0, xss1, xss2, xss3, xss4, xss5, xss6, xss7} {
		r := HeapifyUp(xss, elem, lt)
		fmt.Println(r)
		errors = parentIsLess(r, elem)
		elem = elem - 1
	}
	if errors != nil {
		t.Errorf("\033[31m Test Falsified with: %v  \u001B[30m \n", errors)
	}
}

func TestHeapifyUpEvenNumberBottomNodes(t *testing.T) {
	xss0 := []int{2, 4, 5, 10, 9, 7, 11, 15, 17, 20, 17, 15, 8, 16, 3}
	xss1 := []int{2, 4, 5, 10, 9, 7, 11, 15, 17, 20, 17, 15, 8, 3, 16}
	xss2 := []int{2, 4, 5, 10, 9, 7, 11, 15, 17, 20, 17, 15, 3, 8, 16}
	xss3 := []int{2, 4, 5, 10, 9, 7, 11, 15, 17, 20, 17, 3, 8, 16, 15}
	xss4 := []int{2, 4, 5, 10, 9, 7, 11, 15, 17, 20, 3, 15, 8, 16, 17}
	xss5 := []int{2, 4, 5, 10, 9, 7, 11, 15, 17, 3, 17, 15, 8, 16, 20}
	xss6 := []int{2, 4, 5, 10, 9, 7, 11, 15, 3, 20, 17, 15, 8, 16, 17}
	xss7 := []int{2, 4, 5, 10, 9, 7, 11, 3, 17, 20, 17, 15, 8, 16, 15}

	var elem = 14 //The element you want to add to the almost-a-heap. It's always 3 but position shifts
	var errors error
	for _, xss := range [][]int{xss0, xss1, xss2, xss3, xss4, xss5, xss6, xss7} {
		r := HeapifyUp(xss, elem, lt)
		fmt.Println(r)
		errors = parentIsLess(r, elem)
		elem = elem - 1
	}
	if errors != nil {
		t.Errorf("\033[31m Test Falsified with: %v  \u001B[30m \n", errors)
	}
}

func TestHeapifyDown(t *testing.T) {
	g0 := propcheck.ChooseArray(5, 10, propcheck.ChooseInt(0, 10))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()} //366368000} //time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapifyUp  \n",
		func(xss []int) []int {
			start := time.Now()
			lastIdx := (len(xss) - 1) / 2
			fmt.Println(rng)
			fmt.Printf("Array Before HeapifyDown   :%v\n", xss)
			r := HeapifyDown(xss, lastIdx, lt)
			fmt.Printf("HeapifyDown algorithm for an array of length:%v took:%v\n", len(xss), time.Since(start))
			fmt.Printf("Array After HeapifyDown    :%v\n", xss)
			fmt.Printf("Array After PureHeapifyDown:%v\n", r)
			//
			return r
		},
		func(xss []int) (bool, error) {
			lastIdx := len(xss) - 1
			errors := parentIsLess(xss, lastIdx)
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{1, rng})
	propcheck.ExpectSuccess[[]int](t, result)
	fmt.Println(rng)
}

func TestHeapInsertAndFindMin(t *testing.T) {
	isZeroVal := func(s int) bool {
		if s == 0 {
			return true
		} else {
			return false
		}
	}
	var lt = func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}

	insertIntoHeap := func(xss []int) []int {
		var r = StartHeap[int](len(xss))
		start := time.Now()
		for _, x := range xss {
			r = HeapInsert(r, x, lt, isZeroVal)
		}
		fmt.Printf("Heap HeapInsert algorithm for an array of length:%v took:%v\n", len(xss), time.Since(start))
		return r
	}

	parentLess := func(xss []int) (bool, error) {
		lastIdx := len(xss) - 1
		errors := parentIsLess(xss, lastIdx)
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}
	topIsMin := func(heap []int) (bool, error) {
		s := sorting.QuickSort(heap, lt)
		var errors error
		if s[0] != HeapMin(heap) {
			errors = multierror.Append(errors, fmt.Errorf("top of heap was not smallest. expected:%v actual:%v\n", s[0], heap[0]))
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	g0 := propcheck.ChooseArray(5, 1000, propcheck.ChooseInt(1, 10000))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapMin  \n",
		insertIntoHeap,
		propcheck.AssertionAnd(parentLess, topIsMin),
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
	fmt.Println(rng)
}
