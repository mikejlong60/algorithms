package chapter4

import (
	"github.com/greymatter-io/golangz/heap"
	"testing"
)

func TestMinSpanningTree(t *testing.T) {

	a := &PrimsNode{id: "a", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	b := &PrimsNode{id: "b", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	c := &PrimsNode{id: "c", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	d := &PrimsNode{id: "d", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	e := &PrimsNode{id: "e", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	f := &PrimsNode{id: "f", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	g := &PrimsNode{id: "g", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	ab := &PrimsEdge{
		u:      a.id,
		v:      b.id,
		length: 2,
	}
	ac := &PrimsEdge{
		u:      a.id,
		v:      c.id,
		length: 3,
	}
	ad := &PrimsEdge{
		u:      a.id,
		v:      d.id,
		length: 3,
	}

	bc := &PrimsEdge{
		u:      b.id,
		v:      c.id,
		length: 4,
	}
	be := &PrimsEdge{
		u:      b.id,
		v:      e.id,
		length: 3,
	}

	cd := &PrimsEdge{
		u:      c.id,
		v:      d.id,
		length: 5,
	}
	cf := &PrimsEdge{
		u:      c.id,
		v:      f.id,
		length: 6,
	}
	ce := &PrimsEdge{
		u:      c.id,
		v:      e.id,
		length: 1,
	}

	df := &PrimsEdge{
		u:      d.id,
		v:      f.id,
		length: 7,
	}

	ef := &PrimsEdge{
		u:      e.id,
		v:      f.id,
		length: 8,
	}

	fg := &PrimsEdge{
		u:      f.id,
		v:      g.id,
		length: 9,
	}

	a.connectionsTo = heap.HeapInsert[PrimsEdge, string](a.connectionsTo, ab, primsEdgeLt)
	a.connectionsTo = heap.HeapInsert[PrimsEdge, string](a.connectionsTo, ac, primsEdgeLt)
	a.connectionsTo = heap.HeapInsert[PrimsEdge, string](a.connectionsTo, ad, primsEdgeLt)
	b.connectionsTo = heap.HeapInsert[PrimsEdge, string](b.connectionsTo, bc, primsEdgeLt)
	b.connectionsTo = heap.HeapInsert[PrimsEdge, string](b.connectionsTo, be, primsEdgeLt)
	c.connectionsTo = heap.HeapInsert[PrimsEdge, string](c.connectionsTo, cd, primsEdgeLt)
	c.connectionsTo = heap.HeapInsert[PrimsEdge, string](c.connectionsTo, ce, primsEdgeLt)
	c.connectionsTo = heap.HeapInsert[PrimsEdge, string](c.connectionsTo, cf, primsEdgeLt)
	d.connectionsTo = heap.HeapInsert[PrimsEdge, string](d.connectionsTo, df, primsEdgeLt)
	e.connectionsTo = heap.HeapInsert[PrimsEdge, string](e.connectionsTo, ef, primsEdgeLt)
	f.connectionsTo = heap.HeapInsert[PrimsEdge, string](f.connectionsTo, fg, primsEdgeLt)
	_, totalCost := MinSpanningTree([]*PrimsNode{a, b, c, d, e, f, g}) //Total cost should be 24

	if totalCost != 24 {
		t.Errorf("Actual total cost:%v, expected total cost:%v", totalCost, 24)
	}

}
