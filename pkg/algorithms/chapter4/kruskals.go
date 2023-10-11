package chapter4

import (
	"github.com/greymatter-io/golangz/heap"
)

func kruskals(h heap.Heap[PrimsEdge, string], r map[string]*PrimsEdge, lt func(l, r *PrimsEdge) bool, expectedSize int) (heap.Heap[PrimsEdge, string], map[string]*PrimsEdge, func(l, r *PrimsEdge) bool, int) {
	if len(r) == expectedSize {
		return h, r, lt, expectedSize
	} else {
		a, _ := heap.FindMin(h)

		_, alreadySeen := r[a.v]
		if !alreadySeen {
			r[a.v] = a
		}
		h, _ = heap.HeapDelete(h, 0, lt)
		return kruskals(h, r, lt, expectedSize)
	}
}

func Kruskals(g []*PrimsEdge) []*PrimsEdge {

	// NOTE - this would have been less code if I had just used a sorted array instead of my heap.  But
	// maintaining familiarity with my heap is important because AAC uses it.
	toHeap := func(xs []*PrimsEdge, lt func(l, r *PrimsEdge) bool) heap.Heap[PrimsEdge, string] {
		exf := func(e *PrimsEdge) string {
			return e.v
		}
		h := heap.New[PrimsEdge, string](exf)

		for _, b := range xs {
			h = heap.HeapInsert(h, b, lt)
		}
		return h
	}

	lt := func(l, r *PrimsEdge) bool {
		if l.length < r.length {
			return true
		} else {
			return false
		}
	}

	toArray := func(xs map[string]*PrimsEdge) []*PrimsEdge {
		z := []*PrimsEdge{}
		for _, x := range xs {
			z = append(z, x)
		}
		return z
	}

	numberOfNodesMinus1 := func(xs []*PrimsEdge) int {
		var r = map[string]interface{}{}
		var c = 0
		for _, j := range xs {
			var there bool
			_, there = r[j.u]
			if !there {
				r[j.u] = nil
				c = c + 1
			}
			_, there = r[j.v]
			if !there {
				r[j.v] = nil
				c = c + 1
			}
		}
		return c - 1
	}
	_, r, _, _ := kruskals(toHeap(g, lt), map[string]*PrimsEdge{}, lt, numberOfNodesMinus1(g))

	return toArray(r)
}
