package chapter4

import (
	"github.com/greymatter-io/golangz/sets"
	"testing"
)

func NearTreeMinSpanningTree(g []PrimsEdge) []PrimsEdge {
	type vCost struct {
		u      string
		weight float32
	}

	var seen = map[string]vCost{} //The key of this map is the v NodeId because each node can have only one incoming edge.  But a node can have > 1 outgoing edges.
	for _, j := range g {
		e, seenBefore := seen[j.v]
		if seenBefore {
			//new edge from u is cheaper way to get to v node
			if e.weight > j.weight {
				seen[j.v] = vCost{j.u, j.weight}
			}
		} else {
			seen[j.v] = vCost{j.u, j.weight}
		}

	}

	var r = []PrimsEdge{}
	for v, j := range seen {
		r = append(r, PrimsEdge{
			u:      j.u,
			v:      v,
			weight: j.weight,
		})
	}
	return r
}

var undirectedEq = func(l, r PrimsEdge) bool { //This equality function is for undirected edges.
	if l.u == r.u && l.v == r.v && l.weight == r.weight {
		return true
	} else if l.u == r.v && l.v == r.u && l.weight == r.weight {
		return true
	} else {
		return false
	}
}

func TestNearTreeBigOn1(t *testing.T) {
	ab := PrimsEdge{
		u:      "a",
		v:      "b",
		weight: 4,
	}
	ac := PrimsEdge{
		u:      "a",
		v:      "c",
		weight: 5,
	}
	bc := PrimsEdge{
		u:      "b",
		v:      "c",
		weight: 3,
	}

	var actual = NearTreeMinSpanningTree([]PrimsEdge{ab, ac, bc})
	var expected = []PrimsEdge{ab, bc}
	if !sets.SetEquality(actual, expected, undirectedEq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestNearTreeBigOn2(t *testing.T) {
	ab := PrimsEdge{
		u:      "a",
		v:      "b",
		weight: 4,
	}
	ac := PrimsEdge{
		u:      "a",
		v:      "c",
		weight: 5,
	}
	bc := PrimsEdge{
		u:      "b",
		v:      "c",
		weight: 10,
	}

	var actual = NearTreeMinSpanningTree([]PrimsEdge{ab, ac, bc})
	var expected = []PrimsEdge{ab, ac}
	if !sets.SetEquality(actual, expected, undirectedEq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestNearTreeBigOn3(t *testing.T) {
	ab := PrimsEdge{
		u:      "a",
		v:      "b",
		weight: 1,
	}
	ac := PrimsEdge{
		u:      "a",
		v:      "c",
		weight: 5,
	}
	bc := PrimsEdge{
		u:      "b",
		v:      "c",
		weight: 2,
	}

	ad := PrimsEdge{
		u:      "a",
		v:      "d",
		weight: 6,
	}
	bd := PrimsEdge{
		u:      "b",
		v:      "d",
		weight: 3,
	}
	cd := PrimsEdge{
		u:      "c",
		v:      "d",
		weight: 4,
	}
	var actual = NearTreeMinSpanningTree([]PrimsEdge{ab, ac, bc, ad, bd, cd})
	var expected = []PrimsEdge{ab, bc, bd}
	if !sets.SetEquality(actual, expected, undirectedEq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}
