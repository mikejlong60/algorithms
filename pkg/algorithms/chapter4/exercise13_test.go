package chapter4

import (
	"github.com/go-test/deep"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/sets"
	"testing"
)

func TestShortestPath(t *testing.T) {

	na := Node4{
		Id:          "A",
		Connections: nil,
	}
	nb := Node4{
		Id:          "B",
		Connections: nil,
	}
	aWeight := 10
	bWeight := 2
	na.Connections = map[string]Node4{"A": {Id: "B", Distance: 1 * bWeight}}
	nb.Connections = map[string]Node4{"B": {Id: "A", Distance: 3 * aWeight}}

	graph := map[string]Node4{"A": na, "B": nb}
	actual := DijkstraSearch(graph, "A")
	expected := map[string]*Pq{"A": &Pq{Id: "A", Distance: 0}, "B": &Pq{Id: "B", Distance: 2}}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}

func TestPiss(t *testing.T) {
	na := Node4{
		Id:          "A",
		Connections: nil,
	}
	nb := Node4{
		Id:          "B",
		Connections: nil,
	}
	nc := Node4{
		Id:          "C",
		Connections: nil,
	}
	nd := Node4{
		Id:          "D",
		Connections: nil,
	}
	na.Connections = map[string]Node4{"B": {Id: "B", Distance: 1}, "C": {Id: "C", Distance: 4}}
	nb.Connections = map[string]Node4{"C": {Id: "C", Distance: 2}, "D": {Id: "D", Distance: 5}}
	nc.Connections = map[string]Node4{"D": {Id: "D", Distance: 1}}
	nd.Connections = map[string]Node4{}

	graph := map[string]Node4{"A": na, "B": nb, "C": nc, "D": nd}
	actual := DijkstraSearch(graph, "A")
	expected := map[string]*Pq{"A": {Id: "A", Distance: 0}, "B": {Id: "B", Distance: 1}, "D": {Id: "D", Distance: 4}, "C": {Id: "C", Distance: 3}}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}

func TestShortestCompletionTime(t *testing.T) {
	job1 := Process{
		id:       1,
		length:   10 * 1,
		deadline: 200,
	}
	job2 := Process{
		id:       2,
		length:   2 * 3,
		deadline: 200,
	}

	eq := func(l, r *Process) bool {
		if l.id == r.id {
			return true
		} else {
			return false
		}
	}
	r := []*Process{&job1, &job2}
	actual, maxLate := MinimizeLateness(r)
	expected := []*Process{&job1, &job2}
	if !(arrays.ArrayEquality(actual, expected, eq) && maxLate == &job2 && maxLate.finishTime == 11 && maxLate.finishTime-maxLate.deadline == 1) {
		t.Errorf("Actual Schedule:%v Expected Schedule:%v, Max Late:=%v", actual, expected, maxLate)
	}
}

func TestMinCostArborescence(t *testing.T) {

	eq := func(l, r Edge) bool {
		if l.u == r.u && l.v == r.v {
			return true
		} else {
			return false
		}
	}

	a := Node{id: "a"}
	b := Node{id: "b"}

	ab := Edge{&a, &b, 10}
	ba := Edge{&b, &a, 40}

	b.nodesEntering = []Edge{ab}
	b.nodesEntering = []Edge{ba}

	s := []*Node{&a, &b}
	actual := MinCost(s, &a)
	expected := []Edge{ab, ba}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}
