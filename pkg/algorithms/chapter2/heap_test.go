package chapter2

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
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

var eq = func(l, r int) bool {
	if l == r {
		return true
	} else {
		return false
	}
}
var isZeroVal = func(s int) bool {
	if s == 0 {
		return true
	} else {
		return false
	}
}

var insertIntoHeap = func(xss []int) []int {
	var r = StartHeap[int](len(xss))
	start := time.Now()
	for _, x := range xss {
		r = HeapInsert(r, x, lt, isZeroVal)
	}
	fmt.Printf("Heap HeapInsert algorithm for an array of length:%v took:%v\n", len(xss), time.Since(start))
	return r
}

var parentIsLess = func(heap []int, lastIdx int) error {
	var parent = parentIdx(lastIdx)
	var childIdx = lastIdx
	var errors error
	for parent > 0 {
		//		fmt.Printf("parents value:%v child's value:%v\n", heap[parentIdx], heap[childIdx])
		if heap[parent] > heap[childIdx] {
			errors = multierror.Append(errors, fmt.Errorf("parent:%v value was not less than child's:%v\n", heap[parent], heap[childIdx]))
		}
		childIdx = parent
		parent = parentIdx(childIdx)
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

func TestHeapifyDown(t *testing.T) {
	g0 := sets.ChooseSet(5, 10, propcheck.ChooseInt(0, 10), lt, eq)
	biggerInt := propcheck.ChooseInt(100, 1000)
	g1 := propcheck.Product(g0, biggerInt)
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	insertIntoHeap := func(xss []int) []int {
		var r = StartHeap[int](len(xss))
		for _, x := range xss {
			r = HeapInsert(r, x, lt, isZeroVal)
		}
		fmt.Println("======================")
		fmt.Println(xss)
		fmt.Println(r)
		fmt.Println("======================")
		return r
	}
	prop := propcheck.ForAll(g1,
		"Validate HeapifyDown  \n",
		func(p propcheck.Pair[[]int, int]) propcheck.Pair[[]int, int] {
			//Insert this guaranteed larger element at element position 3 in the heap array
			xss := insertIntoHeap(p.A)
			xss[2] = p.B
			r := HeapifyDown(xss, 2, lt)
			return propcheck.Pair[[]int, int]{r, p.B}
		},
		func(p propcheck.Pair[[]int, int]) (bool, error) {
			//Find the big element in the array so you can tell that it got pushed down
			var bigElementPos = 0
			for i, x := range p.A {
				if x == p.B {
					bigElementPos = i
				}
			}
			errors := parentIsLess(p.A, bigElementPos)
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{1, rng})
	propcheck.ExpectSuccess[propcheck.Pair[[]int, int]](t, result)
	fmt.Println(rng)
}

func TestHeapInsertAndFindMinAndHeapifyUp(t *testing.T) {
	g0 := propcheck.ChooseArray(5, 1000, propcheck.ChooseInt(1, 10000))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate FindMin and FindMin and HeapifyUp  \n",
		insertIntoHeap,
		propcheck.AssertionAnd(parentLess, topIsMin),
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
	fmt.Println(rng)
}

func TestHeapDelete(t *testing.T) {

	var deleteFromHeap = func(xss []int) []int {
		r := insertIntoHeap(xss) //TODO top element is not correct for heap after insertion.  This is a problem in your HeapInsert function
		//Delete half of the elements from the heap
		//		n := len(r) / 2
		fmt.Printf("Now a heap:%v\n", r)
		m := len(r) - 1
		var new []int
		for i := m; i >= 0; i-- {
			//0 is the zero val for integers
			new = HeapDelete(r, i, lt, 0)
			fmt.Printf("Just deleted element:%v and now heap is:%v\n", i, new)
			//if i >= n { //Only delete half the elements starting at the top
			//	break
			//}
		}
		fmt.Printf("Just deleted all elements and now heap is:%v\n", new)
		return r
	}

	g0 := propcheck.ChooseArray(5, 10, propcheck.ChooseInt(1, 20))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapDelete  \n",
		deleteFromHeap,
		propcheck.AssertionAnd(parentLess, topIsMin),
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
	fmt.Println(rng)
}
