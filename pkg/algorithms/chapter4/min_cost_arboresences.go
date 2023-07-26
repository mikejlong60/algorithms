package chapter4

import "fmt"

type ArbNode struct {
	id string
}

type ArbNodeEdge struct {
	u      *ArbNode //the Id of the beginning node of the edge
	v      *ArbNode //the Id of the ending node of the edge
	length int
}

func (w ArbNodeEdge) String() string {
	return fmt.Sprintf("ArbNodeEdge{u:%v, v:%v, length:%v}", w.u, w.v, w.length)
}

func (w ArbNode) String() string {
	return fmt.Sprintf("ArbNode{id:%v}", w.id)
}

func MinCost(g []*ArbNodeEdge, r *ArbNode) []*ArbNodeEdge {
	//r := []*ArbNodeEdge{}
	return g //TODO implement this
}
