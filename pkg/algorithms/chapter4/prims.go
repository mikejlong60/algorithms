package chapter4

import (
	"fmt"
	"github.com/greymatter-io/golangz/heap"
	"math"
)

type PrimsEdge struct {
	u      *PrimsNode
	v      *PrimsNode
	length int
}

type PrimsNode struct {
	id            string
	connectionsTo heap.Heap[PrimsEdge, *PrimsNode]
}

func (w PrimsNode) String() string {
	return fmt.Sprintf("PrimsNode{id:%v, connectionsTo:%v}", w.id, w.connectionsTo)
}
func (w PrimsEdge) String() string {
	return fmt.Sprintf("PrimsEdge{u:%v, v:%v, length:%v}", w.u.id, w.v.id, w.length)
}

func primsEdgeLt(l, r *PrimsEdge) bool {
	if l.length < r.length {
		return true
	} else {
		return false
	}
}

func primsNodeLt(l, r *PrimsNode) bool {
	if l.id < r.id {
		return true
	} else {
		return false
	}
}

func extractor(edge *PrimsEdge) *PrimsNode {
	return edge.v
}

func minSpanningTree(xs []*PrimsNode, xxs []*PrimsNode) ([]*PrimsNode, []*PrimsNode) {
	//1. Is xss len == xs len ? If true return
	//2. Find minimum-cost connected node among all edges in map in xss.
	//2. Add that node to xxs
	//3 Remove all nodes from connections pointing back to any nodes in xs
	//4. call minSpanningTree again with xs and updated xxs

	minNode := func(xs []*PrimsNode) *PrimsNode {
		var m = &PrimsEdge{length: math.MaxInt}
		for _, y := range xs {
			b, err := heap.FindMin(y.connectionsTo)
			if err == nil && b.length < m.length {
				m = b
			}
		}
		return m.v
	}

	removeV := func(xs []*PrimsNode, m *PrimsNode) []*PrimsNode {
		for _, y := range xs {
			b := heap.FindPosition(y.connectionsTo, m)
			if b != -1 {
				conns, _ := heap.HeapDelete(y.connectionsTo, b, primsEdgeLt)
				y.connectionsTo = conns
			}
		}
		return xs
	}

	if len(xs) == len(xxs) {
		return xs, xxs
	} else {
		m := minNode(xxs)
		xxs = append(xxs, m)
		newXs := removeV(xs, m)
		return minSpanningTree(newXs, xxs)
	}
}

func MinSpanningTree(xs []*PrimsNode) []*PrimsNode {
	var xxs = []*PrimsNode{}
	xxs = append(xxs, xs[0])
	_, r := minSpanningTree(xs, xxs)
	return r
}
