package chapter4

import (
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/mikejlong60/algorithms/pkg/algorithms/chapter3"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

var totalSteps2 int

func TestBFSTreeHasShortestPathFromRoot(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	prop := propcheck.ForAll(DirectedGraphGen(1, 100),
		"Generate a random graph and do a Tree search starting from some root.",
		func(graph propcheck.Pair[map[int]*chapter3.Node, int]) ([][]chapter3.Edge, int) {
			xs := DijkstraSearch(graph.A, graph.B)
			return xs, len(graph.A)
		},
		func(e [][]chapter3.Edge, x int) (bool, error) {
			var errors error
			if len(e[0]) != 1 { //First layer has only starting node
				t.Errorf("First layer should have had a single node:%v", e[0])
			}
			var totalEdges int
			for _, b := range e {
				totalEdges = totalEdges + len(b)
			}
			log.Infof("total steps:%v, number of Layers:%v, numberOfNodes:%v", totalSteps2, len(e), x)
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{50, rng})
	propcheck.ExpectSuccess[([][]chapter3.Edge, int](t, result)
}
