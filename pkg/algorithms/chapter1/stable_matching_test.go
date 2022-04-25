package chapter1

import (
	"github.com/greymatter-io/golangz/linked_list"
	"github.com/greymatter-io/golangz/propcheck"
	"testing"
)

func TestStableMatchingWomanConflictsNoIndifference(t *testing.T) {
	w0 := &Woman{
		Id:          "0",
		Preferences: make(map[string]propcheck.Pair[int, *Man]),
		EngagedTo:   nil,
	}
	w1 := &Woman{
		Id:          "1",
		Preferences: make(map[string]propcheck.Pair[int, *Man]),
		EngagedTo:   nil,
	}

	m0 := &Man{
		Id:          "0",
		Preferences: nil,
	}
	m1 := &Man{
		Id:          "1",
		Preferences: nil,
	}

	w0.Preferences[m1.Id] = propcheck.Pair[int, *Man]{0, m1}
	w0.Preferences[m0.Id] = propcheck.Pair[int, *Man]{1, m0}
	w1.Preferences[m1.Id] = propcheck.Pair[int, *Man]{0, m1}
	w1.Preferences[m0.Id] = propcheck.Pair[int, *Man]{1, m0}

	var allMen *linked_list.LinkedList[*Man]
	var manPreferences *linked_list.LinkedList[*Woman]
	manPreferences = linked_list.Push(w1, manPreferences)
	manPreferences = linked_list.Push(w0, manPreferences)

	m0.Preferences = manPreferences
	m1.Preferences = manPreferences
	allMen = linked_list.Push(m1, allMen)
	allMen = linked_list.Push(m0, allMen)
	Match(allMen, womanPrefersMe)
	if w0.EngagedTo.Id != m1.Id {
		t.Errorf("Expected woman 1 to be engaged to man 1")
	}
	if w1.EngagedTo.Id != m0.Id {
		t.Errorf("Expected woman 2 to be engaged to man 2")
	}
	unstableMatchings := unstableMatchings([]*Woman{w0, w1})
	if len(unstableMatchings) > 0 {
		t.Errorf("There were unstable matchings as follow:%v", unstableMatchings)
	}
}

func TestStableMatchingWomanConflicts2(t *testing.T) {
	w0 := &Woman{
		Id:          "0",
		Preferences: make(map[string]propcheck.Pair[int, *Man]),
		EngagedTo:   nil,
	}
	w1 := &Woman{
		Id:          "1",
		Preferences: make(map[string]propcheck.Pair[int, *Man]),
		EngagedTo:   nil,
	}

	m0 := &Man{
		Id:          "0",
		Preferences: nil,
	}
	m1 := &Man{
		Id:          "1",
		Preferences: nil,
	}

	var allMen *linked_list.LinkedList[*Man]
	var manPreferences *linked_list.LinkedList[*Woman]
	manPreferences = linked_list.Push(w1, manPreferences)
	manPreferences = linked_list.Push(w0, manPreferences)

	m0.Preferences = manPreferences
	m1.Preferences = manPreferences
	allMen = linked_list.Push(m1, allMen)
	allMen = linked_list.Push(m0, allMen)

	w0.Preferences[m1.Id] = propcheck.Pair[int, *Man]{1, m1}
	w0.Preferences[m0.Id] = propcheck.Pair[int, *Man]{0, m0}
	w1.Preferences[m1.Id] = propcheck.Pair[int, *Man]{0, m1}
	w1.Preferences[m0.Id] = propcheck.Pair[int, *Man]{1, m0}
	Match(allMen, womanPrefersMe)
	if w0.EngagedTo.Id != m0.Id {
		t.Errorf("Expected woman 0 to be engaged to man 0")
	}
	if w1.EngagedTo.Id != m1.Id {
		t.Errorf("Expected woman 1 to be engaged to man 1")
	}

	unstableMatchings := unstableMatchings([]*Woman{w0, w1})
	if len(unstableMatchings) > 0 {
		t.Errorf("There were unstable matchings as follow:%v", unstableMatchings)
	}
}

func TestStableMatchingNoWomanPreferenceConflicts(t *testing.T) {
	w0 := &Woman{
		Id:          "0",
		Preferences: make(map[string]propcheck.Pair[int, *Man]),
		EngagedTo:   nil,
	}
	w1 := &Woman{
		Id:          "1",
		Preferences: make(map[string]propcheck.Pair[int, *Man]),
		EngagedTo:   nil,
	}

	m0 := &Man{
		Id:          "0",
		Preferences: nil,
	}
	m1 := &Man{
		Id:          "1",
		Preferences: nil,
	}

	var allMen *linked_list.LinkedList[*Man]
	var man0Preferences *linked_list.LinkedList[*Woman]
	man0Preferences = linked_list.Push(w1, man0Preferences)
	man0Preferences = linked_list.Push(w0, man0Preferences)
	var man1Preferences *linked_list.LinkedList[*Woman]
	man1Preferences = linked_list.Push(w0, man1Preferences)
	man1Preferences = linked_list.Push(w1, man1Preferences)

	m0.Preferences = man0Preferences
	m1.Preferences = man1Preferences
	allMen = linked_list.Push(m1, allMen)
	allMen = linked_list.Push(m0, allMen)

	w0.Preferences[m1.Id] = propcheck.Pair[int, *Man]{0, m1}
	w0.Preferences[m0.Id] = propcheck.Pair[int, *Man]{1, m0}
	w1.Preferences[m1.Id] = propcheck.Pair[int, *Man]{1, m1}
	w1.Preferences[m0.Id] = propcheck.Pair[int, *Man]{0, m0}
	Match(allMen, womanPrefersMe)
	if w0.EngagedTo.Id != m0.Id {
		t.Errorf("Expected woman 0 to be engaged to man 0")
	}
	if w1.EngagedTo.Id != m1.Id {
		t.Errorf("Expected woman 1 to be engaged to man 1")
	}
	unstableMatchings := unstableMatchings([]*Woman{w0, w1})
	if len(unstableMatchings) > 0 {
		t.Errorf("There were unstable matchings as follow:%v", unstableMatchings)
	}
}

func TestIndifferentStableMatching(t *testing.T) {
	w0 := &Woman{
		Id:          "0",
		Preferences: make(map[string]propcheck.Pair[int, *Man]),
		EngagedTo:   nil,
	}
	w1 := &Woman{
		Id:          "1",
		Preferences: make(map[string]propcheck.Pair[int, *Man]),
		EngagedTo:   nil,
	}

	w2 := &Woman{
		Id:          "2",
		Preferences: make(map[string]propcheck.Pair[int, *Man]),
		EngagedTo:   nil,
	}

	m0 := &Man{
		Id:          "0",
		Preferences: nil,
	}
	m1 := &Man{
		Id:          "1",
		Preferences: nil,
	}
	m2 := &Man{
		Id:          "2",
		Preferences: nil,
	}

	var allMen *linked_list.LinkedList[*Man]
	var manPreferences *linked_list.LinkedList[*Woman]
	manPreferences = linked_list.Push(w2, manPreferences)
	manPreferences = linked_list.Push(w1, manPreferences)
	manPreferences = linked_list.Push(w0, manPreferences)

	m0.Preferences = manPreferences
	m1.Preferences = manPreferences
	m2.Preferences = manPreferences //TODO Fix this to be different
	allMen = linked_list.Push(m2, allMen)
	allMen = linked_list.Push(m1, allMen)
	allMen = linked_list.Push(m0, allMen)

	w0.Preferences[m1.Id] = propcheck.Pair[int, *Man]{1, m1}
	w0.Preferences[m0.Id] = propcheck.Pair[int, *Man]{0, m0}
	w1.Preferences[m1.Id] = propcheck.Pair[int, *Man]{0, m1}
	w1.Preferences[m0.Id] = propcheck.Pair[int, *Man]{1, m0}
	//w2 is indifferent to all men
	Match(allMen, womanPrefersMe)
	if w0.EngagedTo.Id != m0.Id {
		t.Errorf("Expected woman 0 to be engaged to man 0")
	}
	if w1.EngagedTo.Id != m1.Id {
		t.Errorf("Expected woman 1 to be engaged to man 1")
	}
	if w2.EngagedTo.Id != m2.Id {
		t.Errorf("Expected woman 2 to be engaged to man 2")
	}
	unstableMatchings := unstableMatchings([]*Woman{w0, w1, w2})
	if len(unstableMatchings) > 0 {
		t.Errorf("There were unstable matchings as follow:%v", unstableMatchings)
	}
}
