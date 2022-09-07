package chapter3

import "github.com/greymatter-io/golangz/propcheck"

type NodeForTopoOrdering struct {
	Id            int
	OutgoingEdges map[int]*NodeForTopoOrdering
	IncomingEdges map[int]*NodeForTopoOrdering
}

func GraphOfAllConnectedComponents(graph propcheck.Pair[map[int]*Node, int]) []map[int]NodeForTopoOrdering {

	r1 := make(map[int]NodeForTopoOrdering)
	r2 := []map[int]NodeForTopoOrdering{r1}

	return r2
}
