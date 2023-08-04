package chapter4

import (
	"fmt"
	"github.com/greymatter-io/golangz/sets"
	"github.com/greymatter-io/golangz/sorting"
)

type Node struct {
	id            string
	nodesEntering []Edge
}

type Edge struct {
	u      *Node
	v      *Node
	weight int
}

func (w Node) String() string {
	return fmt.Sprintf("Edge{id:%v, nodesEntering:%v}", w.id, w.nodesEntering)
}

func (w Edge) String() string {
	return fmt.Sprintf("Edge{u:%v, v:%v, weight:%v}", w.u, w.v, w.weight)
}

// g is a Set with id as the key
//TODO this works for a graph with no cycles as per the example here: https://www.youtube.com/watch?v=B5H1qlv9hrg.
//But you still need to detect a cycle and complete the algorithm as shown here: https://www.youtube.com/watch?v=mZBcslesf-o

func MinCost(g []*Node, r *Node) []Edge {

	lt := func(l, r *Node) bool {
		if l.id < r.id {
			return true
		} else {
			return false
		}
	}
	eq := func(l, r *Node) bool {
		if l.id == r.id {
			return true
		} else {
			return false
		}
	}

	ltx := func(l, r Edge) bool {
		if l.weight < r.weight {
			return true
		} else {
			return false
		}
	}
	//Make g a set
	var gs = sets.ToSet(g, lt, eq)

	//Get Root node out of set
	gs = sets.SetMinus(gs, []*Node{r}, eq)

	//Calculate the relative weight of each node entering parent node.
	//Pure function
	relativeCost := func(xs []Edge) []Edge {
		sorting.QuickSort(xs, ltx)
		var rxs = []Edge{}
		for _, y := range xs {
			a := y
			a.weight = a.weight - xs[0].weight
			rxs = append(rxs, a)
		}
		return rxs
	}

	//Assign new weights to each Nodes entering node.
	for _, b := range gs {
		b.nodesEntering = relativeCost(b.nodesEntering)
	}

	//Choose 1 Edge for each Node that has the weight
	result := []Edge{}
	for _, b := range gs {
		result = append(result, b.nodesEntering[0])
	}

	return result
}
