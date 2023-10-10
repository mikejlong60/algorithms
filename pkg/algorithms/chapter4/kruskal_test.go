package chapter4

import (
	"github.com/greymatter-io/golangz/sets"
	"testing"
)

func TestKruskals(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		length: 4,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		length: 2,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		length: 30,
	}

	cd := &PrimsEdge{
		u:      "c",
		v:      "d",
		length: 31,
	}
	cb := &PrimsEdge{
		u:      "c",
		v:      "b",
		length: 3,
	}

	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.length == r.length {
			return true
		} else {
			return false
		}
	}
	actual := Kruskals([]*PrimsEdge{ab, ac, bd, cd, cb}, 3)
	expected := []*PrimsEdge{ac, cb, bd}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}
