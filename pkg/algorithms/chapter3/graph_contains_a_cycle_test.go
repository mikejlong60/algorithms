package chapter3

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/sets"
	"testing"
)

type Node struct {
	Id            int
	IncomingEdges []int //Ids of the nodes that point to this node
	OutgoingEdges []int //Ids of the nodes that this node points to
}

func (w Node) String() string {
	return fmt.Sprintf("Node{Id:%v, IncomingEdges:%v, OutGoingEdges:%v}", w.Id, w.IncomingEdges, w.OutgoingEdges)
}

func topologicalOrdering(graph []Node, topo []Node) ([]Node, []Node) { //1st return value is G, second is topological ordering
	eq := func(a, b Node) bool {
		if a.Id == b.Id {
			return true
		} else {
			return false
		}
	}
	eqInt := func(a, b int) bool {
		if a == b {
			return true
		} else {
			return false
		}
	}
	removeIFromOandI := func(g []Node, nodeId int) []Node {
		var r []Node
		for _, j := range g {
			out := sets.SetMinus(j.OutgoingEdges, []int{nodeId}, eqInt)
			in := sets.SetMinus(j.IncomingEdges, []int{nodeId}, eqInt)
			n := Node{Id: j.Id,
				OutgoingEdges: out,
				IncomingEdges: in,
			}
			r = append(r, n)
		}
		return r
	}

	var x = -1
	var r = make([]Node, len(graph))
	copy(r, graph)
	for _, j := range r {
		if len(j.IncomingEdges) == 0 {
			x = j.Id
			break
		}
	}
	if x != -1 {
		//remove i from all incoming and outgoing edge arrays for every Node
		r = removeIFromOandI(r, x)
		return topologicalOrdering(sets.SetMinus(r, []Node{r[x]}, eq), append(topo, r[x]))
	} else {
		return r, topo
	}
}

func TestGraphContainsACycle(t *testing.T) {
	v1 := Node{Id: 1,
		IncomingEdges: []int{},
		OutgoingEdges: []int{7, 5, 4},
	}
	v2 := Node{Id: 2,
		IncomingEdges: []int{},
		OutgoingEdges: []int{6, 5, 3},
	}
	v3 := Node{Id: 3,
		IncomingEdges: []int{2},
		OutgoingEdges: []int{5, 4},
	}

	v4 := Node{Id: 4,
		IncomingEdges: []int{3, 1},
		OutgoingEdges: []int{5},
	}

	v5 := Node{Id: 5,
		IncomingEdges: []int{1, 2, 3, 4},
		OutgoingEdges: []int{6, 7},
	}

	v6 := Node{Id: 6,
		IncomingEdges: []int{2, 5},
		OutgoingEdges: []int{7},
	}

	v7 := Node{Id: 7,
		IncomingEdges: []int{6, 5, 1},
		OutgoingEdges: []int{},
	}

	graph := []Node{v7, v6, v5, v4, v3, v2, v1}
	expected := []Node{v1, v2, v3, v4, v5, v6, v7}
	//graph := []Node{v1, v2, v3, v4, v5, v6, v7}
	r, topo := topologicalOrdering(graph, []Node{})
	fmt.Printf("graph:%v\n", r)
	fmt.Printf("topological ordering:%v\n", topo)
	eq := func(a, b Node) bool {
		if a.Id == b.Id {
			return true
		} else {
			return false
		}
	}

	if !arrays.ArrayEquality(topo, expected, eq) {
		t.Errorf("\nActual:  %v\nExpected:%v", topo, expected)
	}

	//graph:[Node{Id:7, IncomingEdges:[6 5 1], OutGoingEdges:[]} Node{Id:6, IncomingEdges:[2 5], OutGoingEdges:[7]} Node{Id:5, IncomingEdges:[1 2 3 4], OutGoingEdges:[6 7]} Node{Id:4, IncomingEdges:[3 1], OutGoingEdges:[5]} Node{Id:3, IncomingEdges:[2], OutGoingEdges:[5 4]}]
	//	topological ordering:[Node{Id:2, IncomingEdges:[], OutGoingEdges:[6 5 3]} Node{Id:1, IncomingEdges:[], OutGoingEdges:[7 5 4]}]

}
