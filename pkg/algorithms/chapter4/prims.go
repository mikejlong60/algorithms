package chapter4

import (
	"fmt"
	"github.com/greymatter-io/golangz/heap"
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

func primsLt(l, r *PrimsEdge) bool {
	if l.length < r.length {
		return true
	} else {
		return false
	}
}

func extractor(edge *PrimsEdge) *PrimsNode {
	return edge.u
}

func minSpanningTree([]*PrimsNode) map[string]*PrimsNode {

	return make(map[string]*PrimsNode)
}
