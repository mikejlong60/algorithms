package chapter3

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func TestHalfConnectedNodesIsConnectedGraph(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	//lt := func(l, r int) bool {
	//	if l < r {
	//		return true
	//	} else {
	//		return false
	//	}
	//}
	//eq := func(l, r int) bool {
	//	if l == r {
	//		return true
	//	} else {
	//		return false
	//	}
	//}

	h := func(graph map[int]*Node) (bool, error) {
		//g := func(connectedComponent []Edge) []int {
		//	var r = []int{}
		//	for _, i := range connectedComponent {
		//		r = append(r, i.u)
		//		r = append(r, i.v)
		//	}
		//	return sets.ToSet(r, lt, eq)
		//}
		//allConnectedNodes := arrays.FlatMap(graph.A, g)
		var errors error
		if len(graph)%2 != 10 {
			errors = multierror.Append(errors, fmt.Errorf("not en even mnumber of nodes"))

		}
		//if len(allConnectedNodes) != len(graph.B) {
		//	errors = multierror.Append(errors, fmt.Errorf("Number of Nodes:%v in set of connected components should have equaled same number of Nodes:%v", len(allConnectedNodes), len(graph.B)))
		//}
		//if !sets.SetEquality(allConnectedNodes, graph.B, eq) {
		//	errors = multierror.Append(errors, fmt.Errorf("Not every node was connected"))
		//}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	prop := propcheck.ForAll(EvenNumberOfNodesGen(1, 1000),
		"Prove that a graph of n nodes where each is connected to at least half of the other nodes is connected",
		ConnectEveryNodeToAtLeastHalfOfTheOtherNodes,
		h,
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[map[int]*Node](t, result)
}
