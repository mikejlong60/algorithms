package chapter4

import (
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/mikejlong60/algorithms/pkg/algorithms/chapter3"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestBFSTreeHasShortestPathFromRoot(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	prop := propcheck.ForAll(DirectedGraphGen(1, 500),
		"Generate a random graph and do a Tree search starting from some root.",
		func(graph propcheck.Pair[map[int]*chapter3.Node, int]) propcheck.Pair[[][]chapter3.Edge, int] {
			totalSteps2 = 0
			xs := DijkstraSearch(graph.A, graph.B)
			return propcheck.Pair[[][]chapter3.Edge, int]{xs, len(graph.A)}
		},
		func(e propcheck.Pair[[][]chapter3.Edge, int]) (bool, error) {
			var errors error
			if len(e.A[0]) != 1 { //First layer has only starting node
				t.Errorf("First layer should have had a single node:%v", e.A[0])
			}
			var totalEdges int
			for _, b := range e.A {
				totalEdges = totalEdges + len(b)
			}
			log.Infof("total steps:%v, number of Layers:%v, numberOfNodes:%v, edges:%v", totalSteps2, len(e.A), e.B, totalEdges)
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{50, rng})
	propcheck.ExpectSuccess[propcheck.Pair[map[int]*chapter3.Node, int]](t, result)
}

func TestTreeWithShortestPath(t *testing.T) {
	n0 := chapter3.Node{
		Id:          0,
		Connections: nil,
	}
	n1 := chapter3.Node{
		Id:          1,
		Connections: nil,
	}
	n2 := chapter3.Node{
		Id:          2,
		Connections: nil,
	}
	n3 := chapter3.Node{
		Id:          03,
		Connections: nil,
	}
	n4 := chapter3.Node{
		Id:          4,
		Connections: nil,
	}
	n0.Connections = []*chapter3.Node{&n1, &n2, &n3, &n4}
	n1.Connections = []*chapter3.Node{&n4, &n3}
	n2.Connections = []*chapter3.Node{&n3}
	n3.Connections = []*chapter3.Node{&n4}
	n4.Connections = []*chapter3.Node{&n0}
	graph := make(map[int]*chapter3.Node, 5) //First field of pair is the layer the node is in, -1 means it's never been seen before and is thus not in any layer
	graph[0] = &n0
	graph[1] = &n1
	graph[2] = &n2
	graph[3] = &n3
	graph[4] = &n4
	actual := DijkstraSearch(graph, 0)
	expected := [][]chapter3.Edge{{{-1, 0}}, {{0, 1}, {0, 2}, {0, 3}, {0, 4}}}
	if !arrays.ArrayEquality(actual, expected, chapter3.TreeEquality) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}
