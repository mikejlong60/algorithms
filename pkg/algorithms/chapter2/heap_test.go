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

func gtWhenNonDefaultChild(l, r int) bool {
	if l > r && r != -1 {
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

func TestHeapInsertAndStartHeapAndHeapifyUpWithInts(t *testing.T) {
	g := propcheck.ChooseArray(0, 10, propcheck.ChooseInt(0, 10))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	insertIntoHeap := func(xss []int) []int {
		var r = StartHeap[int](len(xss), -1)
		for _, x := range xss {
			r = HeapInsert(r, x, lt, isZeroVal, -1)
		}
		return r
	}
	insert := func(p []int) []int {
		xss := insertIntoHeap(p)
		return xss
	}
	validateIsAHeap := func(p []int) (bool, error) {
		var errors error
		for idx, _ := range p {
			errors = parentIsGreater(p, idx, gtWhenNonDefaultChild)
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}
	validateHeapMin := func(p []int) (bool, error) {
		var errors error
		var sorted = make([]int, len(p))
		copy(sorted, p)
		sorting.QuickSort(sorted, lt)
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
		var sorted = make([]string, len(p))
		copy(sorted, p)
		sorting.QuickSort(sorted, lt)
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
}

func TestHeapDelete(t *testing.T) {

	insertIntoHeap := func(xss []int) []int {
		var r = StartHeap[int](len(xss), -1)
		for _, x := range xss {
			r = HeapInsert(r, x, lt, isZeroVal, -1)
		}
		return r
	}

	var deleteFromHeap = func(xss []int) []int {
		r := insertIntoHeap(xss)
		l := len(r)
		r = HeapDelete(r, l/2, lt, isZeroVal, intZeroVal)
		if l > 0 {
			r = HeapDelete(r, l-1, lt, isZeroVal, intZeroVal)
		}
		if l > 1 {
			r = HeapDelete(r, l-2, lt, isZeroVal, intZeroVal)
		}
		if l > 2 {
			r = HeapDelete(r, l-3, lt, isZeroVal, intZeroVal)
		}
		return r
	}

	validateIsAHeap := func(p []int) (bool, error) {
		var errors error
		for idx, _ := range p {
			errors = parentIsGreater(p, idx, gtWhenNonDefaultChild)
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}
	validateHeapMin := func(p []int) (bool, error) {
		var errors error
		var sorted = make([]int, len(p))
		copy(sorted, p)
		sorting.QuickSort(sorted, lt)
		if len(p) > 0 && FindMin(p) != sorted[0] {
			errors = multierror.Append(errors, fmt.Errorf("FindMin returned:%v but should have returned:%v", FindMin(p), sorted[0]))
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	g0 := propcheck.ChooseArray(5, 10, propcheck.ChooseInt(1, 2000))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapDelete  \n",
		deleteFromHeap,
		validateIsAHeap, validateHeapMin,
	)
	result := prop.Run(propcheck.RunParms{6, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
