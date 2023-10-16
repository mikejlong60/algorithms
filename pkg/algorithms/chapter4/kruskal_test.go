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
	actual := Kruskals([]*PrimsEdge{ab, ac, bd, cd, cb})
	expected := []*PrimsEdge{ac, cb, bd}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestKruskalsUsingUnionFind(t *testing.T) {
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
	actual := KruskalsUsingUnionFind([]*PrimsEdge{ab, ac, bd, cd, cb})
	expected := []*PrimsEdge{ac, cb, bd}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestKruskals1(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		length: 4.1,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		length: 4.2,
	}

	ad := &PrimsEdge{
		u:      "a",
		v:      "d",
		length: 4.3,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		length: 4.4,
	}

	cd := &PrimsEdge{
		u:      "c",
		v:      "d",
		length: 4.5,
	}
	cb := &PrimsEdge{
		u:      "c",
		v:      "b",
		length: 4.6,
	}

	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.length == r.length {
			return true
		} else {
			return false
		}
	}
	actual := Kruskals([]*PrimsEdge{ab, ac, ad, bd, cd, cb})
	expected := []*PrimsEdge{ac, ad, ab}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestKruskals1UsingUnionFind(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		length: 4.1,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		length: 4.2,
	}

	ad := &PrimsEdge{
		u:      "a",
		v:      "d",
		length: 4.3,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		length: 4.4,
	}

	cd := &PrimsEdge{
		u:      "c",
		v:      "d",
		length: 4.5,
	}
	cb := &PrimsEdge{
		u:      "c",
		v:      "b",
		length: 4.6,
	}

	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.length == r.length {
			return true
		} else {
			return false
		}
	}
	actual := KruskalsUsingUnionFind([]*PrimsEdge{ab, ac, ad, bd, cd, cb})
	expected := []*PrimsEdge{ac, ad, ab}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestKruskals2(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		length: 4.1,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		length: 4.2,
	}

	ad := &PrimsEdge{
		u:      "a",
		v:      "d",
		length: 1,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		length: 4.3,
	}

	cd := &PrimsEdge{
		u:      "c",
		v:      "d",
		length: 4.4,
	}
	cb := &PrimsEdge{
		u:      "c",
		v:      "b",
		length: 4.5,
	}

	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.length == r.length {
			return true
		} else {
			return false
		}
	}
	actual := Kruskals([]*PrimsEdge{ab, ac, ad, bd, cd, cb})
	expected := []*PrimsEdge{ac, ab, ad}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestKruskals2UsingUnionFind(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		length: 4.1,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		length: 4.2,
	}

	ad := &PrimsEdge{
		u:      "a",
		v:      "d",
		length: 1,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		length: 4.3,
	}

	cd := &PrimsEdge{
		u:      "c",
		v:      "d",
		length: 4.4,
	}
	cb := &PrimsEdge{
		u:      "c",
		v:      "b",
		length: 4.5,
	}

	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.length == r.length {
			return true
		} else {
			return false
		}
	}
	actual := KruskalsUsingUnionFind([]*PrimsEdge{ab, ac, ad, bd, cd, cb})
	expected := []*PrimsEdge{ac, ab, ad}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestKruskals3(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		length: 4,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		length: 8,
	}

	bc := &PrimsEdge{
		u:      "b",
		v:      "c",
		length: 11,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		length: 8.1,
	}

	cf := &PrimsEdge{
		u:      "c",
		v:      "f",
		length: 1,
	}

	ce := &PrimsEdge{
		u:      "c",
		v:      "e",
		length: 7.1,
	}
	de := &PrimsEdge{
		u:      "d",
		v:      "e",
		length: 2,
	}
	dh := &PrimsEdge{
		u:      "d",
		v:      "h",
		length: 4.1,
	}
	dg := &PrimsEdge{
		u:      "d",
		v:      "g",
		length: 7.2,
	}

	ef := &PrimsEdge{
		u:      "e",
		v:      "f",
		length: 6,
	}
	fh := &PrimsEdge{
		u:      "f",
		v:      "h",
		length: 2.1,
	}

	hi := &PrimsEdge{
		u:      "h",
		v:      "i",
		length: 10,
	}
	gh := &PrimsEdge{
		u:      "g",
		v:      "h",
		length: 14,
	}
	gi := &PrimsEdge{
		u:      "g",
		v:      "i",
		length: 9,
	}
	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.length == r.length {
			return true
		} else {
			return false
		}
	}
	actual := Kruskals([]*PrimsEdge{ab, ac, bc, bd, cf, ce, de, dh, dg, ef, fh, hi, gi, gh})
	expected := []*PrimsEdge{ab, ac, cf, de, dh, dg, gi, fh}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestKruskals3UsingUnionFind(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		length: 4,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		length: 8,
	}

	bc := &PrimsEdge{
		u:      "b",
		v:      "c",
		length: 11,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		length: 8.1,
	}

	cf := &PrimsEdge{
		u:      "c",
		v:      "f",
		length: 1,
	}

	ce := &PrimsEdge{
		u:      "c",
		v:      "e",
		length: 7.1,
	}
	de := &PrimsEdge{
		u:      "d",
		v:      "e",
		length: 2,
	}
	dh := &PrimsEdge{
		u:      "d",
		v:      "h",
		length: 4.1,
	}
	dg := &PrimsEdge{
		u:      "d",
		v:      "g",
		length: 7.2,
	}

	ef := &PrimsEdge{
		u:      "e",
		v:      "f",
		length: 6,
	}
	fh := &PrimsEdge{
		u:      "f",
		v:      "h",
		length: 2.1,
	}

	hi := &PrimsEdge{
		u:      "h",
		v:      "i",
		length: 10,
	}
	gh := &PrimsEdge{
		u:      "g",
		v:      "h",
		length: 14,
	}
	gi := &PrimsEdge{
		u:      "g",
		v:      "i",
		length: 9,
	}
	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.length == r.length {
			return true
		} else {
			return false
		}
	}
	actual := KruskalsUsingUnionFind([]*PrimsEdge{ab, ac, bc, bd, cf, ce, de, dh, dg, ef, fh, hi, gi, gh})
	expected := []*PrimsEdge{ab, ac, bc, cf, fh, dh, dg, gi}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}
