package chapter4

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sorting"
	"github.com/hashicorp/go-multierror"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func gtWhenNonDefaultChild(l, r *Cache) bool {
	if l.ts > r.ts && r != nil {
		return true
	} else {
		return false
	}
}

// Find the mimimum and compares it to the actual min in the initial array.
// If array is empty or filled with nil pointers that is OK and FindMin should not fail
// but return a Golang error, not panic.
func minimumCorrectValue(p, sorted []*Cache) bool {
	min, err := FindMin(p)
	if len(p) > 0 && err == nil {
		return min.ts == sorted[0].ts
	} else {
		return true
	}
}

func parentIsGreater(heap []*Cache, lastIdx int, parentGT func(l, r *Cache) bool) error {
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

func TestHeapInsertWithEmptyHeap(t *testing.T) {
	g := propcheck.ChooseArray(0, 100, propcheck.ChooseInt(0, 100))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	insertIntoHeap := func(xss []int) []*Cache {
		var r = []*Cache{}
		for _, x := range xss {
			r = HeapInsert(r, &Cache{x, fmt.Sprintf("ts:%v", x)}, lt)
		}
		return r
	}
	insert := func(p []int) []*Cache {
		xss := insertIntoHeap(p)
		return xss
	}
	validateIsAHeap := func(p []*Cache) (bool, error) {
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
	validateHeapMin := func(p []*Cache) (bool, error) {
		var errors error
		var sorted = make([]*Cache, len(p))
		copy(sorted, p)
		sorting.QuickSort(sorted, lt)
		if !minimumCorrectValue(p, sorted) {
			errors = multierror.Append(errors, fmt.Errorf("FindMin should have returned:%v", sorted[0]))
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

func TestHeapInsertWithNonEmptyHeapAndHolesAtEndOfHeap(t *testing.T) {
	g := propcheck.ChooseArray(0, 10000, propcheck.ChooseInt(0, 100))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	insertIntoHeap := func(xss []int) []*Cache {
		var r = StartHeap(len(xss) + 21)
		for _, x := range xss {
			r = HeapInsert(r, &Cache{x, fmt.Sprintf("ts:%v", x)}, lt)
		}
		return r
	}
	insert := func(p []int) []*Cache {
		xss := insertIntoHeap(p)
		return xss
	}
	validateIsAHeap := func(p []*Cache) (bool, error) {
		var errors error
		for idx, x := range p {
			if x != nil { //Heap could have some empty elements at end
				errors = parentIsGreater(p, idx, gtWhenNonDefaultChild)
			}
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	ltNoNilCheck := func(l, r *Cache) bool {
		if l.ts < r.ts {
			return true
		} else {
			return false
		}
	}

	validateHeapMin := func(p []*Cache) (bool, error) {
		var errors error
		var sorted = make([]*Cache, len(p))
		copy(sorted, p)
		for idx, x := range sorted { //Trim off em,pty elements at end to work around your sorting bug
			if x == nil { //Heap could have some empty elements at end
				sorted = sorted[0:idx]
				break
			}
		}

		sorting.QuickSort(sorted, ltNoNilCheck)
		if !minimumCorrectValue(p, sorted) {
			errors = multierror.Append(errors, fmt.Errorf("FindMin should have returned:%v", sorted[0]))
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	prop := propcheck.ForAll(g,
		"Validate HeapInsertWithNonEmptyHeapAndHolesAtEndOfHeap  \n",
		insert,
		validateIsAHeap, validateHeapMin,
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}

func TestHeapInsertWithNonEmptyHeapAndNoHolesAtEndOfHeap(t *testing.T) {
	g := propcheck.ChooseArray(0, 100, propcheck.ChooseInt(0, 100))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	insertIntoHeap := func(xss []int) []*Cache {
		var r = StartHeap(len(xss))
		for _, x := range xss {
			r = HeapInsert(r, &Cache{x, fmt.Sprintf("ts:%v", x)}, lt)
		}
		return r
	}
	insert := func(p []int) []*Cache {
		xss := insertIntoHeap(p)
		return xss
	}
	validateIsAHeap := func(p []*Cache) (bool, error) {
		var errors error
		for idx, x := range p {
			if x != nil { //Heap could have some empty elements at end
				errors = parentIsGreater(p, idx, gtWhenNonDefaultChild)
			}
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	ltNoNilCheck := func(l, r *Cache) bool {
		if l.ts < r.ts {
			return true
		} else {
			return false
		}
	}

	validateHeapMin := func(p []*Cache) (bool, error) {
		var errors error
		var sorted = make([]*Cache, len(p))
		copy(sorted, p)
		for idx, x := range sorted { //Trim off em,pty elements at end to work around your sorting bug
			if x == nil { //Heap could have some empty elements at end
				sorted = sorted[0:idx]
				break
			}
		}

		sorting.QuickSort(sorted, ltNoNilCheck)
		if !minimumCorrectValue(p, sorted) {
			errors = multierror.Append(errors, fmt.Errorf("FindMin should have returned:%v", sorted[0]))
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
	result := prop.Run(propcheck.RunParms{1000, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}

var it = 0

func TestHeapDelete1(t *testing.T) {

	insertIntoHeap := func(xss []int) []*Cache {
		var r = StartHeap(1)
		for _, x := range xss {
			r = HeapInsert(r, &Cache{x, fmt.Sprintf("ts:%v", x)}, lt)
		}
		return r
	}

	var deleteFromHeap = func(xss []int) []*Cache {
		it = it + 1
		log.Debugf("Iyteration:%v", it)
		r := insertIntoHeap(xss)
		//var l = len(r)
		var err error
		r, err = HeapDelete(r, 5, lt)
		log.Info(err)
		r, err = HeapDelete(r, 4, lt)

		log.Info(err)
		r, err = HeapDelete(r, 3, lt)
		log.Info(err)
		r, err = HeapDelete(r, 2, lt)
		log.Info(err)
		r, err = HeapDelete(r, 1, lt)
		log.Info(err)
		r, err = HeapDelete(r, 0, lt)
		log.Info(err)
		return r
	}

	validateIsAHeap := func(p []*Cache) (bool, error) {
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
	validateHeapMin := func(p []*Cache) (bool, error) {
		var errors error
		var sorted = make([]*Cache, len(p))
		copy(sorted, p)
		sorting.QuickSort(sorted, lt)
		if !minimumCorrectValue(p, sorted) {
			errors = multierror.Append(errors, fmt.Errorf("FindMin should have returned:%v", sorted[0]))
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	g0 := propcheck.ChooseArray(0, 400, propcheck.ChooseInt(1, 2000))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapDelete  \n",
		deleteFromHeap,
		validateIsAHeap, validateHeapMin,
	)
	result := prop.Run(propcheck.RunParms{100, rng}) //The 3rd iteration paniced with array out or bounds
	propcheck.ExpectSuccess[[]int](t, result)
}

func TestHeapDelete2(t *testing.T) {
	var errors error
	insertIntoHeap := func(xss []int) []*Cache {
		var r = StartHeap(1)
		for _, x := range xss {
			r = HeapInsert(r, &Cache{x, fmt.Sprintf("ts:%v", x)}, lt)
		}
		return r
	}

	correctHeapMin := func(p []*Cache) bool {
		var sorted = make([]*Cache, len(p))
		copy(sorted, p)
		sorting.QuickSort(sorted, lt)
		if !minimumCorrectValue(p, sorted) {
			return false
		} else {
			return true
		}
	}

	var deleteAllFromHeap = func(xss []int) []*Cache {
		var r = insertIntoHeap(xss)
		for range r {
			r, _ = HeapDelete(r, 0, lt)
			if !correctHeapMin(r) {
				errors = multierror.Append(errors, fmt.Errorf("Heap property violated"))
			}
		}
		return r
	}

	validateIsAHeap := func(p []*Cache) (bool, error) {
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

	heapWrong := func(p []*Cache) (bool, error) {
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	g0 := propcheck.ChooseArray(0, 9, propcheck.ChooseInt(1, 2000))
	rng := propcheck.SimpleRNG{531874217} //time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapDelete  \n",
		deleteAllFromHeap,
		validateIsAHeap, heapWrong,
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
