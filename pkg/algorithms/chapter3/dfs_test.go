package chapter3

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"testing"
	"time"
)

func TestDFS(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	prop := propcheck.ForAll(UndirectedGraphGen(1, 100),
		"Generate a random graph and do a Tree search starting from some root.",
		func(graph propcheck.Pair[map[int]*Node, int]) propcheck.Pair[*Node, map[int]*Node] {
			var tree []Edge
			start := time.Now()
			fmt.Println("starting DFS search")
			node, seenNodes, dfsTree := DFSearch(graph.A[graph.B], make(map[int]*Node), tree)
			fmt.Printf("DFS on a graph of size:%v took %v\n", len(graph.A), time.Since(start))
			//fmt.Println(graph)
			fmt.Println(dfsTree)
			a := propcheck.Pair[*Node, map[int]*Node]{node, seenNodes}
			return a
		},
		func(p propcheck.Pair[*Node, map[int]*Node]) (bool, error) {
			var errors error
			//if !p.A.B {
			//	t.Errorf("Rule 3 failure:%v", p.B)
			//}
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

func TestDumb(t *testing.T) {
	n1 := Node{
		Id:          1,
		Connections: nil,
	}
	n2 := Node{
		Id:          2,
		Connections: nil,
	}
	n3 := Node{
		Id:          3,
		Connections: nil,
	}
	n4 := Node{
		Id:          4,
		Connections: nil,
	}
	n5 := Node{
		Id:          5,
		Connections: nil,
	}
	n6 := Node{
		Id:          6,
		Connections: nil,
	}
	n7 := Node{
		Id:          7,
		Connections: nil,
	}
	n8 := Node{
		Id:          8,
		Connections: nil,
	}
	n9 := Node{
		Id:          9,
		Connections: nil,
	}
	n10 := Node{
		Id:          10,
		Connections: nil,
	}
	n11 := Node{
		Id:          11,
		Connections: nil,
	}
	n12 := Node{
		Id:          12,
		Connections: nil,
	}
	n13 := Node{
		Id:          13,
		Connections: nil,
	}

	n1.Connections = []*Node{&n2, &n3}
	n2.Connections = []*Node{&n1, &n3, &n4, &n5}
	n3.Connections = []*Node{&n1, &n2, &n5, &n7, &n8}
	n4.Connections = []*Node{&n5, &n2}
	n5.Connections = []*Node{&n3, &n4, &n6, &n2}
	n7.Connections = []*Node{&n3, &n8}
	n8.Connections = []*Node{&n3, &n7}
	n9.Connections = []*Node{&n10}
	n10.Connections = []*Node{&n9}

	n11.Connections = []*Node{&n12}
	n12.Connections = []*Node{&n13, &n11}
	n13.Connections = []*Node{&n12}

	var tree []Edge
	node, seenNodes, actual := DFSearch(&n1, make(map[int]*Node), tree)
	expected := []Edge{{1, 2}, {2, 3}, {3, 5}, {5, 4}, {5, 6}, {3, 7}, {7, 8}}

	edgeEq := func(a, b Edge) bool {
		if a.u == b.u && a.v == b.v {
			return true
		} else {
			return false
		}
	}

	if !arrays.ArrayEquality(actual, expected, edgeEq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}

	fmt.Println(node)
	fmt.Println(seenNodes)
}
