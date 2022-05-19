package chapter2

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"github.com/hashicorp/go-multierror"
	"math/rand"
	"testing"
	"time"
)

var shuffle = func(toBeShuffled []int) []int {
	r := make([]int, len(toBeShuffled))
	copy(r, toBeShuffled)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})
	return r
}
var lt = func(l, r int) bool {
	if l < r {
		return true
	} else {
		return false
	}
}

var eq = func(l, r int) bool {
	if l == r {
		return true
	} else {
		return false
	}
}

var parentIsLess = func(heap []int, lastIdx int) error {
	var parentIdx = ParentIdx(lastIdx)
	//var currentParsent = heap[lastIdx]
	var childIdx = lastIdx
	var errors error
	for parentIdx > 0 {
		fmt.Printf("parents value:%v child's value:%v\n", heap[parentIdx], heap[childIdx])
		if heap[parentIdx] > heap[childIdx] {
			errors = multierror.Append(errors, fmt.Errorf("parent:%v value was not less than child's:%v\n", heap[parentIdx], heap[childIdx]))
			//fmt.Printf("parent:%v value was not less than child's:%v\n", heap[parentIdx], currentParent)
		}
		childIdx = parentIdx
		parentIdx = ParentIdx(childIdx)
		//currentParent = heap[parentIdx]
	}
	return errors
}

func TestHeapifyUp(t *testing.T) {
	g0 := sets.ChooseSet(0, 100, propcheck.ChooseInt(0, 1000), lt, eq)
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapifyUp  \n",
		func(xs []int) []int {
			//xss := shuffle(xs)
			//TODO verify this with random arrays
			xss := []int{2, 4, 5, 10, 9, 7, 11, 15, 17, 20, 17, 15, 8, 16, 3}

			start := time.Now()
			fmt.Printf("Array Before HeapifyUp:%v\n", xss)
			lastIdx := len(xss) - 1
			r := HeapifyUp(xss, lastIdx)
			fmt.Printf("HeapUp algorithm for an array of length:%v took:%v\n", len(xss), time.Since(start))
			fmt.Printf("Array Before HeapifyUp:%v\n", r)
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

func TestHeapifyDown(t *testing.T) {
	g0 := sets.ChooseSet(0, 100, propcheck.ChooseInt(0, 1000), lt, eq)
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapifyUp  \n",
		func(xs []int) []int {
			xss := shuffle(xs)
			start := time.Now()
			fmt.Printf("Array Before HeapifyUp:%v\n", xss)
			lastIdx := len(xss) - 1
			r := HeapifyDown(xss, lastIdx)
			fmt.Printf("HeapUp algorithm for an array of length:%v took:%v\n", len(xss), time.Since(start))
			fmt.Printf("Array Before HeapifyUp:%v\n", r)
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
