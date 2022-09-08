package chapter3

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"testing"
	"time"
)

func topo(m map[int]*NodeForTopoOrdering, accum []*NodeForTopoOrdering) (map[int]*NodeForTopoOrdering, []*NodeForTopoOrdering) {
	if len(m) == 0 {
		return m, accum
	} else {
		//Find next node `n` in map `m` with no incoming edges
		//Iterate over each `n`'s outgoing connections `p` and remove `n` from `p`'s list of incoming connections
		//Append `n` to accum and store in `r2`
		//Remove `n` from `m` and store in `r1`
		return topo(m, accum)
	}
}

func TestTopologicalOrdering(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	makeConnectedComponentsAsNodeForTopoOrdering := func(a propcheck.Pair[map[int]*Node, int]) map[int]*NodeForTopoOrdering {
		graph := make(map[int]NodeForTopoOrdering, len(a.A))
		for _, xs := range a.A { //Convert initial list of nodes to the type from which you can make a topological ordering.
			ie := make(map[int]*NodeForTopoOrdering)
			oe := make(map[int]*NodeForTopoOrdering)
			graph[xs.Id] = NodeForTopoOrdering{Id: xs.Id, IncomingEdges: ie, OutgoingEdges: oe}
		}

		cc := GenerateConnectedComponents(a)
		var nodes = make(map[int]*NodeForTopoOrdering)
		for _, xs := range cc.A {
			if len(xs) > 1 { //Transform first connected component graph that is larger than one node into its equivalent NodeForTopoOrdering map for computing topo ordering
				for _, ys := range xs {
					n := graph[ys.u]
					oe := graph[ys.v]
					oe.IncomingEdges[n.Id] = &n
					n.OutgoingEdges[ys.v] = &oe
					if ys.u != ys.v { //Don't add top-level edge that points to itself
						n.IncomingEdges[ys.u] = &n
					}
					nodes[n.Id] = &n
				}
				break
			}
		}
		return nodes
	}

	prop := propcheck.ForAll(UndirectedGraphGen(1, 100),
		"Generate a directed graph from which you compute a topological ordering.",
		makeConnectedComponentsAsNodeForTopoOrdering,
		//func(graph map[int]*NodeForTopoOrdering) propcheck.Pair[map[int]*NodeForTopoOrdering, []*NodeForTopoOrdering] {
		//	r1, r2 := topo(graph, []*NodeForTopoOrdering{})
		//	return propcheck.Pair[map[int]*NodeForTopoOrdering, []*NodeForTopoOrdering]{r1, r2}
		//},
		func(xs map[int]*NodeForTopoOrdering) (bool, error) {
			var errors error
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[map[int]*NodeForTopoOrdering](t, result)
	fmt.Println(rng) //
}
