package chapter2

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sorting"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func lt(l, r int) bool {
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
var isZeroVal = func(s int) bool {
	if s == -1 {
		return true
	} else {
		return false
	}
}

func insertIntoHeap(xss []int, zeroVal int) []int {
	var r = StartHeap(len(xss), zeroVal)

	start := time.Now()
	for _, x := range xss {
		r = HeapInsert(r, x, lt, isZeroVal, zeroVal)
	}
	fmt.Printf("Heap HeapInsert algorithm for an array of length:%v took:%v\n", len(xss), time.Since(start))
	return r
}

var parentIsLess = func(heap []int, lastIdx int) error {
	var pIdx = parentIdx(lastIdx)
	var cIdx = lastIdx
	var errors error
	for pIdx > 0 {
		//		fmt.Printf("parents value:%v child's value:%v\n", heap[parentIdx], heap[childIdx])
		if heap[pIdx] > heap[cIdx] {
			errors = multierror.Append(errors, fmt.Errorf("parent:%v value was not less than or equal to child's:%v\n", heap[pIdx], heap[cIdx]))
		}
		cIdx = pIdx
		pIdx = parentIdx(cIdx)
	}
	return errors
}

var parentLess = func(xss []int) (bool, error) {
	lastIdx := len(xss) - 1
	errors := parentIsLess(xss, lastIdx)
	if errors != nil {
		return false, errors
	} else {
		return true, nil
	}
}

var topIsMin = func(heap []int) (bool, error) {
	s := sorting.QuickSort(heap, lt)
	var errors error
	if s[0] != FindMin(heap) {
		errors = multierror.Append(errors, fmt.Errorf("top of heap was not smallest. expected:%v actual:%v\n", s[0], heap[0]))
	}
	if errors != nil {
		return false, errors
	} else {
		return true, nil
	}
}

//func TestHeapifyDown(t *testing.T) {
//	g0 := sets.ChooseSet(5, 10, propcheck.ChooseInt(0, 10), lt, eq)
//	biggerInt := propcheck.ChooseInt(100, 1000)
//	g1 := propcheck.Product(g0, biggerInt)
//	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
//	insertIntoHeap := func(xss []int) []int {
//		var r = StartHeap[int](len(xss), -1)
//		for _, x := range xss {
//			r = HeapInsert(r, x, lt, isZeroVal, -1)
//		}
//		fmt.Println("======================")
//		fmt.Println(xss)
//		fmt.Println(r)
//		fmt.Println("======================")
//		return r
//	}
//	prop := propcheck.ForAll(g1,
//		"Validate HeapifyDown  \n",
//		func(p propcheck.Pair[[]int, int]) propcheck.Pair[[]int, int] {
//			//Insert this guaranteed larger element at element position 3 in the heap array
//			xss := insertIntoHeap(p.A)
//			xss[2] = p.B
//			r := HeapifyDown(xss, 2, lt)
//			return propcheck.Pair[[]int, int]{r, p.B}
//		},
//		func(p propcheck.Pair[[]int, int]) (bool, error) {
//			//Find the big element in the array so you can tell that it got pushed down
//			var bigElementPos = 0
//			for i, x := range p.A {
//				if x == p.B {
//					bigElementPos = i
//				}
//			}
//			errors := parentIsLess(p.A, bigElementPos)
//			if errors != nil {
//				return false, errors
//			} else {
//				return true, nil
//			}
//		},
//	)
//	result := prop.Run(propcheck.RunParms{1, rng})
//	propcheck.ExpectSuccess[propcheck.Pair[[]int, int]](t, result)
//	fmt.Println(rng)
//}

func TestHeapInsertAndStartHeapAndHeapifyUp(t *testing.T) {
	g := propcheck.ChooseArray(0, 10, propcheck.ChooseInt(0, 10))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	insertIntoHeap := func(xss []int) []int {
		var r = StartHeap[int](len(xss), -1)
		for _, x := range xss {
			r = HeapInsert(r, x, lt, isZeroVal, -1)
		}
		fmt.Println("======================")
		fmt.Println(xss)
		fmt.Println(r)
		fmt.Println("======================")
		return r
	}
	insert := func(p []int) []int {
		xss := insertIntoHeap(p)
		return xss
	}
	validateIsAHeap := func(p []int) (bool, error) {
		var errors error
		for idx, _ := range p {
			errors = parentIsLess(p, idx)
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}
	validateHeapMin := func(p []int) (bool, error) {
		var errors error
		sorted := sorting.QuickSort(p, lt)
		if len(p) > 0 && FindMin(p) != sorted[0] {
			errors = multierror.Append(errors, fmt.Errorf("FindMin returned:%v but should have returned:%v", FindMin(p), sorted[0]))
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	prop := propcheck.ForAll(g,
		"Validate HeapifyUp  \n",
		insert,
		validateIsAHeap, validateHeapMin,
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
	fmt.Println(rng)
}

//func TestHeapDelete(t *testing.T) {
//
//	var deleteFromHeap = func(xss []int) []int {
//		r := insertIntoHeap(xss, -1) //TODO top element is not correct for heap after insertion.  This is a problem in your HeapInsert function
//		//Delete half of the elements from the heap
//		//		n := len(r) / 2
//		fmt.Printf("Now a heap:%v\n", r)
//		m := len(r) - 1
//		var new []int
//		for i := m; i >= 0; i-- {
//			//0 is the zero val for integers
//			new = HeapDelete(r, i, lt, 0)
//			fmt.Printf("Just deleted element:%v and now heap is:%v\n", i, new)
//			//if i >= n { //Only delete half the elements starting at the top
//			//	break
//			//}
//		}
//		fmt.Printf("Just deleted all elements and now heap is:%v\n", new)
//		return r
//	}
//
//	g0 := propcheck.ChooseArray(5, 10, propcheck.ChooseInt(1, 20))
//	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
//	prop := propcheck.ForAll(g0,
//		"Validate HeapDelete  \n",
//		deleteFromHeap,
//		propcheck.AssertionAnd(parentLess, topIsMin),
//	)
//	result := prop.Run(propcheck.RunParms{100, rng})
//	propcheck.ExpectSuccess[[]int](t, result)
//	fmt.Println(rng)
//}
