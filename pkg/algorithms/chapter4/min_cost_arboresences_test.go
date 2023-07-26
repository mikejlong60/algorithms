package chapter4

import (
	"github.com/greymatter-io/golangz/sets"
	"testing"
)

func TestBookExample(t *testing.T) {

	eq := func(l, r *ArbNodeEdge) bool {
		if l.u == r.u && l.v == r.v && l.length == r.length {
			return true
		} else {
			return false
		}
	}

	a := ArbNode{id: "a"}
	b := ArbNode{id: "b"}
	c := ArbNode{id: "c"}
	d := ArbNode{id: "d"}
	e := ArbNode{id: "e"}
	f := ArbNode{id: "f"}

	ab := ArbNodeEdge{&a, &b, 2}
	ac := ArbNodeEdge{&a, &c, 3}
	ad := ArbNodeEdge{&a, &d, 3}
	bd := ArbNodeEdge{&b, &d, 1}
	be := ArbNodeEdge{&b, &e, 2}
	cb := ArbNodeEdge{&c, &b, 1}
	cf := ArbNodeEdge{&c, &f, 3}
	de := ArbNodeEdge{&d, &e, 1}
	df := ArbNodeEdge{&d, &f, 3}
	ec := ArbNodeEdge{&e, &c, 1}

	s := []*ArbNodeEdge{&ab, &ac, &ad, &bd, &be, &cb, &cf, &de, &df, &ec}
	actual := MinCost(s, &a)
	expected := []*ArbNodeEdge{&ab, &bd, &de, &df, &ec}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}
