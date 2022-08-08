package chapter3

import (
	"testing"
)

func TestTopologicalOrdering(t *testing.T) {
	//rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	//nodeEq := func(l, r *NodeForTopoOrdering) bool {
	//	if l.Id == r.Id {
	//		return true
	//	} else {
	//		return false
	//	}
	//}
	//computeIncomingEdges := func(g map[int]*NodeForTopoOrdering) {
	//	for _, fromNode := range g {
	//		for _, toNode := range g {
	//			if arrays.Contains(toNode.Connections, fromNode, nodeEq) {
	//				fromNode.IncomingEdges = append(fromNode.IncomingEdges, toNode)
	//			}
	//		}
	//	}
	//
	//}
	//prop := propcheck.ForAll(DirectedGraphGen(1, 1000),
	//	"Generate a directed graph from which you compute a topological ordering.",
	//	func(graph propcheck.Pair[map[int]*Node, int]) []int {
	//
	//		var tree []Edge
	//		dfStart := time.Now()
	//		_, _, dfTree := DFSearch(graph.A[graph.B], make(map[int]*Node), tree)
	//		fmt.Printf("DFS on a graph of size:%v took %v\n", len(graph.A), time.Since(dfStart))
	//		bfStart := time.Now()
	//		bfTree, _, _ := BFSearch(graph.A, graph.B)
	//		fmt.Printf("BFS on a graph of size:%v took %v\n", len(graph.A), time.Since(bfStart))
	//		bf := func(e []Edge) []int {
	//			var r []int
	//			for _, b := range e {
	//				r = append(r, b.u)
	//				r = append(r, b.v)
	//			}
	//			return r
	//		}
	//		df := func(e Edge) []int {
	//			var r []int
	//			r = append(r, e.u)
	//			r = append(r, e.v)
	//			return r
	//		}
	//		lt := func(l, r int) bool {
	//			if l < r {
	//				return true
	//			} else {
	//				return false
	//			}
	//		}
	//		eq := func(l, r int) bool {
	//			if l == r {
	//				return true
	//			} else {
	//				return false
	//			}
	//		}
	//		dfConnectedComponent := sets.ToSet(arrays.FlatMap(dfTree, df), lt, eq)
	//		bfConnectedComponent := sets.SetMinus(sets.ToSet(arrays.FlatMap(bfTree, bf), lt, eq), []int{-1}, eq) //Remove the default -1 first node
	//		if len(bfConnectedComponent) == 1 && len(dfConnectedComponent) == 0 {                                //If root node has no connections then df has no tree and this is fine.
	//			return []int{}
	//		} else {
	//			return sets.SetMinus(bfConnectedComponent, dfConnectedComponent, eq)
	//		}
	//	},
	//	func(xs []int) (bool, error) {
	//		var errors error
	//		if len(xs) > 0 {
	//			errors = multierror.Append(errors, fmt.Errorf("Depth-first and breadth-first search should have produced the same set of components but differed(bf - df=%v\n", xs))
	//		}
	//		if errors != nil {
	//			return false, errors
	//		} else {
	//			return true, nil
	//		}
	//	},
	//)
	//result := prop.Run(propcheck.RunParms{100, rng})
	//propcheck.ExpectSuccess[propcheck.Pair[map[int]*Node, int]](t, result)
	//fmt.Println(rng)
}
