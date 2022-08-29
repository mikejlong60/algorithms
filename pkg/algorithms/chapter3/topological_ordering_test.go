package chapter3

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

type NodeForTopoOrdering2 struct {
	Id            int
	Connections   map[int]*NodeForTopoOrdering2
	IncomingEdges map[int]*NodeForTopoOrdering2
}

func topo(m map[int]*NodeForTopoOrdering2, accum []*NodeForTopoOrdering2) (map[int]*NodeForTopoOrdering2, []*NodeForTopoOrdering2) {
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

	simpleDirectedGraph := propcheck.Id(make(map[int]*NodeForTopoOrdering2))

	//TODO Add a real graph

	prop := propcheck.ForAll(simpleDirectedGraph,
		"Generate a directed graph from which you compute a topological ordering.",
		func(graph map[int]*NodeForTopoOrdering2) propcheck.Pair[map[int]*NodeForTopoOrdering2, []*NodeForTopoOrdering2] {
			r1, r2 := topo(graph, []*NodeForTopoOrdering2{})
			return propcheck.Pair[map[int]*NodeForTopoOrdering2, []*NodeForTopoOrdering2]{r1, r2}
		},
		func(xs propcheck.Pair[map[int]*NodeForTopoOrdering2, []*NodeForTopoOrdering2]) (bool, error) {
			var errors error
			if len(xs.B) > 0 {
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
	propcheck.ExpectSuccess[map[int]*NodeForTopoOrdering2](t, result)
	fmt.Println(rng) //
}
