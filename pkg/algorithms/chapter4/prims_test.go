package chapter4

import (
	"fmt"
	"github.com/greymatter-io/golangz/heap"
	"github.com/greymatter-io/golangz/sets"
	"testing"
)

var edgeEq = func(l, r *PrimsEdge) bool {
	if l.length == r.length && l.u == r.u && l.v == r.v {
		return true
	} else {
		return false
	}
}

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
	actual, totalCost := MinSpanningTree([]*PrimsNode{a, b, c, d, e, f, g}) //Total cost should be 24

	if totalCost != 24 {
		t.Errorf("Actual total cost:%v, expected total cost:%v", totalCost, 24)
	}
	if len(actual) != 6 {
		t.Errorf("Actual # Edges:%v, Expected # Edges:%v", len(actual), 6)
	}

	expected := []*PrimsEdge{ab, ad, ac, ce, cf, fg}
	if !sets.SetEquality(actual, expected, edgeEq) {
		t.Errorf("Actual Edges:%v, Expected Edges:%v", actual, expected)

	}

}

func TestMinBottleneckSpanningTree(t *testing.T) {

	a := &PrimsNode{id: "a", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	b := &PrimsNode{id: "b", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	c := &PrimsNode{id: "c", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	d := &PrimsNode{id: "d", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	ab := &PrimsEdge{
		u:      a.id,
		v:      b.id,
		length: 4,
	}
	ac := &PrimsEdge{
		u:      a.id,
		v:      c.id,
		length: 2,
	}

	bd := &PrimsEdge{
		u:      b.id,
		v:      d.id,
		length: 30,
	}

	cd := &PrimsEdge{
		u:      c.id,
		v:      d.id,
		length: 31,
	}
	cb := &PrimsEdge{
		u:      c.id,
		v:      b.id,
		length: 3,
	}

	a.connectionsTo = heap.HeapInsert[PrimsEdge, string](a.connectionsTo, ab, primsEdgeLt)
	a.connectionsTo = heap.HeapInsert[PrimsEdge, string](a.connectionsTo, ac, primsEdgeLt)
	b.connectionsTo = heap.HeapInsert[PrimsEdge, string](b.connectionsTo, bd, primsEdgeLt)
	c.connectionsTo = heap.HeapInsert[PrimsEdge, string](c.connectionsTo, cd, primsEdgeLt)
	c.connectionsTo = heap.HeapInsert[PrimsEdge, string](c.connectionsTo, cb, primsEdgeLt)
	actual, totalCost := MinSpanningTree([]*PrimsNode{a, b, c, d}) //Total cost should be 24

	if totalCost != 35 {
		t.Errorf("Actual total cost:%v, expected total cost:%v", totalCost, 35)
	}
	if len(actual) != 3 {
		t.Errorf("Actual # Edges:%v, Expected # Edges:%v", len(actual), 3)
	}

	expected := []*PrimsEdge{ac, bd, cb}
	if !sets.SetEquality(actual, expected, edgeEq) {
		t.Errorf("Actual Edges:%v, Expected Edges:%v", actual, expected)

	}
}

func TestAddSmallerEdge(t *testing.T) {

	a := &PrimsNode{id: "a", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	b := &PrimsNode{id: "b", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	c := &PrimsNode{id: "c", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	d := &PrimsNode{id: "d", connectionsTo: heap.New[PrimsEdge, string](extractor)}
	ab := &PrimsEdge{
		u:      a.id,
		v:      b.id,
		length: 2,
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

	dc := &PrimsEdge{
		u:      d.id,
		v:      c.id,
		length: 2,
	}

	a.connectionsTo = heap.HeapInsert[PrimsEdge, string](a.connectionsTo, ab, primsEdgeLt)
	a.connectionsTo = heap.HeapInsert[PrimsEdge, string](a.connectionsTo, ad, primsEdgeLt)
	b.connectionsTo = heap.HeapInsert[PrimsEdge, string](b.connectionsTo, bc, primsEdgeLt)
	d.connectionsTo = heap.HeapInsert[PrimsEdge, string](d.connectionsTo, dc, primsEdgeLt)

	//Changing the algorithm to return a map instead of an array allows me to O(1) lookup whether
	//or not the new edge ending in v would be in the minimum spanning tree. See the test below and
	//remove the commented-out section and comment out the additional edge db below. Then verify
	//this manually.  This is part (a) of question 10.
	//g := []*PrimsNode{a, b, c, d}
	//	actual := MinSpanningTreeReturningMap(g)
	//	fmt.Println(actual)
	db := &PrimsEdge{
		u:      d.id,
		v:      b.id,
		length: 1,
	}
	d.connectionsTo = heap.HeapInsert[PrimsEdge, string](d.connectionsTo, db, primsEdgeLt)
	g := []*PrimsNode{a, b, c, d}

	actual := MinSpanningTreeReturningMap(g)
	fmt.Println(actual)
	if actual[db.v].length < db.length {
		fmt.Println("new edge would not be in tree")
	}
}
