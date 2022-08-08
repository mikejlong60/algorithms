package chapter3

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func TestEqualityOfNodesInDfAndBfSearch(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	prop := propcheck.ForAll(UndirectedGraphGen(1, 1000),
		"Generate a random graph and do a Tree search starting from some root.",
		func(graph propcheck.Pair[map[int]*Node, int]) []int {
			var tree []Edge
			dfStart := time.Now()
			_, _, dfTree := DFSearch(graph.A[graph.B], make(map[int]*Node), tree)
			fmt.Printf("DFS on a graph of size:%v took %v\n", len(graph.A), time.Since(dfStart))
			bfStart := time.Now()
			bfTree, _, _ := BFSearch(graph.A, graph.B)
			fmt.Printf("BFS on a graph of size:%v took %v\n", len(graph.A), time.Since(bfStart))
			bf := func(e []Edge) []int {
				var r []int
				for _, b := range e {
					r = append(r, b.u)
					r = append(r, b.v)
				}
				return r
			}
			df := func(e Edge) []int {
				var r []int
				r = append(r, e.u)
				r = append(r, e.v)
				return r
			}
			lt := func(l, r int) bool {
				if l < r {
					return true
				} else {
					return false
				}
			}
			eq := func(l, r int) bool {
				if l == r {
					return true
				} else {
					return false
				}
			}
			dfConnectedComponent := sets.ToSet(arrays.FlatMap(dfTree, df), lt, eq)
			bfConnectedComponent := sets.SetMinus(sets.ToSet(arrays.FlatMap(bfTree, bf), lt, eq), []int{-1}, eq) //Remove the default -1 first node
			if len(bfConnectedComponent) == 1 && len(dfConnectedComponent) == 0 {                                //If root node has no connections then df has no tree and this is fine.
				return []int{}
			} else {
				return sets.SetMinus(bfConnectedComponent, dfConnectedComponent, eq)
			}
		},
		func(xs []int) (bool, error) {
			var errors error
			if len(xs) > 0 {
				errors = multierror.Append(errors, fmt.Errorf("Depth-first and breadth-first search should have produced the same set of components but differed(bf - df=%v\n", xs))
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[propcheck.Pair[map[int]*Node, int]](t, result)
	fmt.Println(rng)
}

func TestAllConnectedComponentsGeneration(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}
	eq := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}

	prop := propcheck.ForAll(UndirectedGraphGen(1, 50),
		"Generate the set of connected components for a given graph.",
		func(graph propcheck.Pair[map[int]*Node, int]) [][]Edge {
			var tree []Edge
			makeAllUnseenNodes := func(nodeMap map[int]*Node) []int {
				var r []int
				for i, _ := range nodeMap {
					r = append(r, i)
				}
				return r
			}
			var unseenNodes = makeAllUnseenNodes(graph.A)

			toNodeIdSet := func(tree []Edge) []int {
				var r []int
				for _, node := range tree {
					r = append(r, node.v)
					r = append(r, node.u)
				}
				r = sets.ToSet(r, lt, eq)
				return r
			}

			var allConnectedComponents [][]Edge
			for len(unseenNodes) > 0 {
				_, _, dfTree := DFSearch(graph.A[unseenNodes[0]], make(map[int]*Node), tree) //Build a tree from the first node of the unseen nodes.
				if len(dfTree) == 0 {                                                        //If tree is empty the node has no connections but is still connected to itself
					dfTree = []Edge{{unseenNodes[0], unseenNodes[0]}}
				}
				allConnectedComponents = append(allConnectedComponents, dfTree)
				dfNodes := toNodeIdSet(dfTree)
				unseenNodes = sets.SetMinus(unseenNodes, dfNodes, eq)
			}
			return allConnectedComponents
		},
		func(xs [][]Edge) (bool, error) { //TODO take the original graph and make sure the xs(connected components array) contains every element in the original graph.
			//You must pass the original graph to this function as well
			var errors error
			//	fmt.Println(len(xs))
			for x, y := range xs {
				fmt.Printf("Connected component:%v, %v\n", x, y)

			}
			//	fmt.Println(xs)
			//if len(xs) > 0 {
			//	errors = multierror.Append(errors, fmt.Errorf("Depth-first and breadth-first search should have produced the same set of components but differed(bf - df=%v\n", xs))
			//}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	fmt.Println(rng)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[propcheck.Pair[map[int]*Node, int]](t, result)
	fmt.Println(rng)
}
