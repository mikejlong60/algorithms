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

func TestHeapInsertAndStartHeapAndHeapifyUpWithInts(t *testing.T) {
	g := propcheck.ChooseArray(0, 20000, propcheck.ChooseInt(0, 1000000))
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
		if len(p) > 0 && FindMin(p).ts != sorted[0].ts {
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

func TestHeapDelete1(t *testing.T) {

	insertIntoHeap := func(xss []int) []*Cache {
		var r = StartHeap(1)
		for _, x := range xss {
			r = HeapInsert(r, &Cache{x, fmt.Sprintf("ts:%v", x)}, lt)
		}
		return r
	}

	var deleteFromHeap = func(xss []int) []*Cache {
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
		if len(p) > 0 && FindMin(p) != sorted[0] {
			errors = multierror.Append(errors, fmt.Errorf("FindMin returned:%v but should have returned:%v", FindMin(p), sorted[0]))
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	g0 := propcheck.ChooseArray(0, 14, propcheck.ChooseInt(1, 2000))
	rng := propcheck.SimpleRNG{30579879} //time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapDelete  \n",
		deleteFromHeap,
		validateIsAHeap, validateHeapMin,
	)
	result := prop.Run(propcheck.RunParms{100, rng}) //The 3rd iteration paniced with array out or bounds
	propcheck.ExpectSuccess[[]int](t, result)
}

func TestHeapDelete2(t *testing.T) {

	insertIntoHeap := func(xss []int) []*Cache {
		var r = StartHeap(1)
		for _, x := range xss {
			r = HeapInsert(r, &Cache{x, fmt.Sprintf("ts:%v", x)}, lt)
		}
		return r
	}

	var deleteFromHeap = func(xss []int) []*Cache {
		r := insertIntoHeap(xss)
		var err error
		r, err = HeapDelete(r, 0, lt)
		log.Info(err)
		r, err = HeapDelete(r, 1, lt)

		log.Info(err)
		r, err = HeapDelete(r, 2, lt)
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
		if len(p) > 0 && FindMin(p) != sorted[0] {
			errors = multierror.Append(errors, fmt.Errorf("FindMin returned:%v but should have returned:%v", FindMin(p), sorted[0]))
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	g0 := propcheck.ChooseArray(0, 14, propcheck.ChooseInt(1, 2000))
	rng := propcheck.SimpleRNG{30579879} //time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapDelete  \n",
		deleteFromHeap,
		validateIsAHeap, validateHeapMin,
	)
	result := prop.Run(propcheck.RunParms{100, rng}) //The 3rd iteration paniced with array out or bounds
	propcheck.ExpectSuccess[[]int](t, result)
}
