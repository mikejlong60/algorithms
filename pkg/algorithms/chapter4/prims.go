package chapter4

import (
	"fmt"
	"github.com/greymatter-io/golangz/heap"
	"math"
)

type PrimsEdge struct {
	u      string
	v      string
	length int
}

type PrimsNode struct {
	id            string
	connectionsTo heap.Heap[PrimsEdge, string]
}

func (w PrimsNode) String() string {
	return fmt.Sprintf("PrimsNode{id:%v, connectionsTo:%v}", w.id, w.connectionsTo)
}
func (w PrimsEdge) String() string {
	return fmt.Sprintf("PrimsEdge{u:%v, v:%v, length:%v}", w.u, w.v, w.length)
}

func primsEdgeLt(l, r *PrimsEdge) bool {
	if l.length < r.length {
		return true
	} else {
		return false
	}
}

func extractor(edge *PrimsEdge) string {
	return edge.v
}

func minSpanningTree(xs []*PrimsNode, xxs []*PrimsEdge) ([]*PrimsNode, []*PrimsEdge) {
	//1. Is minEdge == nil return
	//2. Find minimum-cost connected edge(minEdge) among all edges in array of PrimsNodes.
	//3. Delete all edges that point to minEdge.v in original xs array.
	//4. Add that minEdge to PrimsEdge result (xxs) array
	//4. call minSpanningTree again with updated xs and xxs

	deleteAllEdgesPointingToV := func(xs []*PrimsNode, v *PrimsEdge) []*PrimsNode {
		for _, y := range xs {
			p := heap.FindPosition(y.connectionsTo, v.v)
			if p > -1 {
				y.connectionsTo, _ = heap.HeapDelete(y.connectionsTo, p, primsEdgeLt)
			}
		}
		return xs
	}

	minEdge := func(xxs []*PrimsNode) *PrimsEdge {
		var lowestEdge = &PrimsEdge{
			u:      "",
			v:      "",
			length: math.MaxInt,
		}
		for _, y := range xxs {
			a, err := heap.FindMin(y.connectionsTo)
			if err == nil && a.length < lowestEdge.length {
				lowestEdge = a
			}
		}
		if lowestEdge.length == math.MaxInt {
			return nil
		} else {
			return lowestEdge
		}
	}

	e := minEdge(xs)
	if e == nil {
		return xs, xxs
	}
	xs = deleteAllEdgesPointingToV(xs, e)
	xxs = append(xxs, e)
	return minSpanningTree(xs, xxs)
}

func MinSpanningTree(xs []*PrimsNode) ([]*PrimsEdge, int) {
	totalCost := func(xs []*PrimsEdge) int {
		var r int
		for _, b := range xs {
			r = b.length + r
		}
		return r
	}
	_, r := minSpanningTree(xs, []*PrimsEdge{})
	return r, totalCost(r)
}
