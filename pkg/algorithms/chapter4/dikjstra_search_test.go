package chapter4

import (
	"github.com/go-test/deep"
	"math"
	"testing"
)

func TestDikjstra1(t *testing.T) {
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
	actual := DikjstraSearch(graph, "A")
	expected := map[string]*Pq{"A": &Pq{Id: "A", Distance: 0}, "B": &Pq{Id: "B", Distance: 1}, "C": &Pq{Id: "C", Distance: 3}, "D": &Pq{Id: "D", Distance: 4}}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}

func TestDikjstra2(t *testing.T) {
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
	actual := DikjstraSearch(graph, "B")
	expected := map[string]*Pq{"A": &Pq{Id: "A", Distance: math.MaxInt64}, "B": &Pq{Id: "B", Distance: 0}, "C": &Pq{Id: "C", Distance: 2}, "D": &Pq{Id: "D", Distance: 3}}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Error(diff)
	}
}
