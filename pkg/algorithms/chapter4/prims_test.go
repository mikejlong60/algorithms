package chapter4

import (
	"fmt"
	"github.com/greymatter-io/golangz/heap"
	"testing"
)

func TestMinSpanningTree(t *testing.T) {

	a := &PrimsNode{id: "a", connectionsTo: heap.New[PrimsEdge, *PrimsNode](extractor)}
	b := &PrimsNode{id: "b", connectionsTo: heap.New[PrimsEdge, *PrimsNode](extractor)}
	c := &PrimsNode{id: "c", connectionsTo: heap.New[PrimsEdge, *PrimsNode](extractor)}
	d := &PrimsNode{id: "d", connectionsTo: heap.New[PrimsEdge, *PrimsNode](extractor)}
	e := &PrimsNode{id: "e", connectionsTo: heap.New[PrimsEdge, *PrimsNode](extractor)}
	f := &PrimsNode{id: "f", connectionsTo: heap.New[PrimsEdge, *PrimsNode](extractor)}
	g := &PrimsNode{id: "g", connectionsTo: heap.New[PrimsEdge, *PrimsNode](extractor)}
	ab := &PrimsEdge{
		u:      a,
		v:      b,
		length: 2,
	}
	ac := &PrimsEdge{
		u:      a,
		v:      c,
		length: 3,
	}
	ad := &PrimsEdge{
		u:      a,
		v:      d,
		length: 3,
	}

	bc := &PrimsEdge{
		u:      b,
		v:      c,
		length: 4,
	}
	be := &PrimsEdge{
		u:      b,
		v:      e,
		length: 3,
	}

	cd := &PrimsEdge{
		u:      c,
		v:      d,
		length: 5,
	}
	cf := &PrimsEdge{
		u:      c,
		v:      f,
		length: 6,
	}
	ce := &PrimsEdge{
		u:      c,
		v:      e,
		length: 1,
	}

	df := &PrimsEdge{
		u:      d,
		v:      f,
		length: 7,
	}

	ef := &PrimsEdge{
		u:      e,
		v:      f,
		length: 8,
	}

	fg := &PrimsEdge{
		u:      f,
		v:      g,
		length: 9,
	}

	a.connectionsTo = heap.HeapInsert[PrimsEdge, *PrimsNode](a.connectionsTo, ab, primsLt)
	a.connectionsTo = heap.HeapInsert[PrimsEdge, *PrimsNode](a.connectionsTo, ac, primsLt)
	a.connectionsTo = heap.HeapInsert[PrimsEdge, *PrimsNode](a.connectionsTo, ad, primsLt)
	b.connectionsTo = heap.HeapInsert[PrimsEdge, *PrimsNode](b.connectionsTo, bc, primsLt)
	b.connectionsTo = heap.HeapInsert[PrimsEdge, *PrimsNode](b.connectionsTo, be, primsLt)
	c.connectionsTo = heap.HeapInsert[PrimsEdge, *PrimsNode](c.connectionsTo, cd, primsLt)
	c.connectionsTo = heap.HeapInsert[PrimsEdge, *PrimsNode](c.connectionsTo, ce, primsLt)
	c.connectionsTo = heap.HeapInsert[PrimsEdge, *PrimsNode](c.connectionsTo, cf, primsLt)
	d.connectionsTo = heap.HeapInsert[PrimsEdge, *PrimsNode](d.connectionsTo, df, primsLt)
	e.connectionsTo = heap.HeapInsert[PrimsEdge, *PrimsNode](e.connectionsTo, ef, primsLt)
	f.connectionsTo = heap.HeapInsert[PrimsEdge, *PrimsNode](f.connectionsTo, fg, primsLt)
	fmt.Println(heap.FindMin(f.connectionsTo))
	actual := minSpanningTree([]*PrimsNode{a, b, c, d, e, f, g})
	fmt.Println(actual)

}
