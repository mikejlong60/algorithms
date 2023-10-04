package chapter4

func MinSpanningTreeReturningMap(xs []*PrimsNode) map[string]*PrimsEdge {
	_, r := minSpanningTree(xs, []*PrimsEdge{})

	rr := make(map[string]*PrimsEdge)
	for _, b := range r {
		rr[b.v] = b
	}
	return rr
}
