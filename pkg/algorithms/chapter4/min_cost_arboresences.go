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
	return fmt.Sprintf("Edge{u:%v, v:%v, weight:%v}", w.u.id, w.v.id, w.weight)
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

	//isArborescence := func(g []*Node, xs []Edge) bool {
	//	if len(xs) == len(g)-1 {
	//		return true
	//	} else {
	//		return false
	//	}
	//
	//}
	//Assign new weights to each Nodes entering node.
	for _, b := range gs {
		b.nodesEntering = relativeCost(b.nodesEntering)
	}

	isCycle := func(xs []Edge) bool {
		var seen = make(map[string]struct{})
		for _, x := range xs {
			_, uthere := seen[x.u.id]
			if !uthere {
				seen[x.u.id] = struct{}{}
				_, vthere := seen[x.v.id]
				if vthere {
					return true
				}
				//} else {
				//	return true
			}
		}
		return false
	}

	buildCycleEdges := func(xs []Edge, end Edge) []Edge {
		cycle := []Edge{}
		for i := len(xs) - 1; i >= 0; i-- {
			if xs[i].v == end.u {
				break
			} else {
				cycle = append(cycle, xs[i])
			}
		}
		return cycle
	}
	//Choose 1 Edge for each Node that has the least weight(it's zero) as long as it's not a cycle
	var result = []Edge{}
	for _, b := range gs {
		leastEdge := b.nodesEntering[0]
		result = append(result, leastEdge)
		if isCycle(result) {

			//TODO RESULT is  cycle!!! ow fix it
			//Here is where you start to deal with the cycle:
			//The Steps:
			//    1. The beginning of the cycle is the node leastEdge.v you are pointing to.  The end of the cycle is the node leastEdge.u.
			//    2. Trace backwards from the current node until you reach the node that you were pointing to here(leastEdge.u), recording
			//       each Edge and that is the complete cycle.
			cycle := buildCycleEdges(result, leastEdge)
			fmt.Println(cycle)
			//    3. Given that set of Edges, grab all the edges that point to any node in the cycle. And then choose the minimum weight
			//       of all those. That edge is the one that enters the cycle. Call it A.
			//    4. Then remove the edge from the cycle that points to A. That's where the cycle gets broken.
			fmt.Println("Detected cycle")
		}
	}

	//	if isArborescence(g, result) {
	return result
	//	} else {
	//		return MinCost(g, r)
	//	}
}
