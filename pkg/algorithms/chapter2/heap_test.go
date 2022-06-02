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

func gt(l, r int) bool {
	if l > r {
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
var intZeroVal = -1

var isZeroVal = func(s int) bool {
	if s == intZeroVal {
		return true
	} else {
		return false
	}
}

func parentIsGreater[A any](heap []A, lastIdx int, parentGT func(l, r A) bool) error {
	var pIdx = ParentIdx(lastIdx)
	var cIdx = lastIdx
	var errors error
	for pIdx > 0 {
		if parentGT(heap[pIdx], heap[cIdx]) {
			errors = multierror.Append(errors, fmt.Errorf("parent:%v value was not less than or equal to child's:%v\n", heap[pIdx], heap[cIdx]))
		}
		cIdx = pIdx
		pIdx = ParentIdx(cIdx)
	}
	return errors
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

func TestHeapInsertAndStartHeapAndHeapifyUpWithInts(t *testing.T) {
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
			errors = parentIsGreater(p, idx, gt)
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

func TestHeapInsertAndStartHeapAndHeapifyUpWithStrings(t *testing.T) {
	zero := "Donald Trump"
	isZeroVal := func(s string) bool {
		if s == zero {
			return true
		} else {
			return false
		}
	}
	gt := func(l, r string) bool {
		if l > r {
			return true
		} else {
			return false
		}
	}

	lt := func(l, r string) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}
	g := propcheck.ChooseArray(0, 10, propcheck.String(5))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	insertIntoHeap := func(xss []string) []string {
		var r = StartHeap[string](len(xss), zero)
		for _, x := range xss {
			r = HeapInsert(r, x, lt, isZeroVal, zero)
		}
		fmt.Println("======================")
		fmt.Println(xss)
		fmt.Println(r)
		fmt.Println("======================")
		return r
	}
	insert := func(p []string) []string {
		xss := insertIntoHeap(p)
		return xss
	}
	validateIsAHeap := func(p []string) (bool, error) {
		var errors error
		for idx, _ := range p {
			errors = parentIsGreater(p, idx, gt)
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}
	validateHeapMin := func(p []string) (bool, error) {
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
		"Validate HeapifyUp for String \n",
		insert,
		validateIsAHeap, validateHeapMin,
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]string](t, result)
	fmt.Println(rng)
}

func TestHeapDelete(t *testing.T) {

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

	var deleteFromHeap = func(xss []int) []int {
		r := insertIntoHeap(xss)
		fmt.Printf("Now a heap:%v\n", r)
		m := len(r) / 2 // Delete an element near middle of heap
		//TODO test deletion of other elements
		r = HeapDelete(r, m, lt, isZeroVal, intZeroVal)
		fmt.Printf("Just deleted element:%v and now heap is:%v\n", m, r)
		return r
	}

	validateIsAHeap := func(p []int) (bool, error) {
		var errors error
		for idx, _ := range p {
			errors = parentIsGreater(p, idx, gt)
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

	g0 := propcheck.ChooseArray(5, 10, propcheck.ChooseInt(1, 20))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapDelete  \n",
		deleteFromHeap,
		validateIsAHeap, validateHeapMin,
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
	fmt.Println(rng)
}
