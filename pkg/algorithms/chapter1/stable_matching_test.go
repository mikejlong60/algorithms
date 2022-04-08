package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/linked_list"
	"testing"
)

func TestStableMatchingWomanConflictsNoIndifference(t *testing.T) {
	w0 := &Woman{
		Id:          "0",
		Preferences: make(map[string]int),
		EngagedTo:   nil,
	}
	w1 := &Woman{
		Id:          "1",
		Preferences: make(map[string]int),
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

	w0.Preferences[m1.Id] = 0
	w0.Preferences[m0.Id] = 1
	w1.Preferences[m1.Id] = 0
	w1.Preferences[m0.Id] = 1

	var allMen *linked_list.LinkedList[*Man]
	var manPreferences *linked_list.LinkedList[*Woman]
	manPreferences = linked_list.Push(w1, manPreferences)
	manPreferences = linked_list.Push(w0, manPreferences)

	m0.Preferences = manPreferences
	m1.Preferences = manPreferences
	allMen = linked_list.Push(m1, allMen)
	allMen = linked_list.Push(m0, allMen)
	actual := Match(allMen, womanPrefersMe)
	if w0.EngagedTo.Id != m1.Id {
		t.Errorf("Expected woman 1 to be engaged to man 1")
	}
	if w1.EngagedTo.Id != m0.Id {
		t.Errorf("Expected woman 2 to be engaged to man 2")
	}
	fmt.Println(actual)
}

var womanPrefersMe = func(wp *Woman, courtier *Man) bool { //Does woman prefer this man to the one to which she is currently assigned?
	//This function assumes that the wp woman is already engaged.
	courtierRanking, courtierIsInPreferredList := wp.Preferences[courtier.Id]
	currentFianceeRanking, currentFianceeIsInPreferredList := wp.Preferences[wp.EngagedTo.Id]

	if !courtierIsInPreferredList && !currentFianceeIsInPreferredList { //Woman is indifferent to both men. So she will stick with her current fiancee.
		fmt.Printf("Woman:%v is indifferent to both men. So she will stick with her current fiancee%v\n", wp.Id, wp.EngagedTo.Id)
		return false
	} else if courtierIsInPreferredList && !currentFianceeIsInPreferredList { //Woman is indifferent to her current husband but not the courtier. So she chooses the courtier.
		fmt.Printf("Woman:%v is indifferent to her current husband:%v but not the courtier:%v. So she chooses the courtier.\n", wp.Id, wp.EngagedTo.Id, courtier.Id)
		return true
	} else if !courtierIsInPreferredList && currentFianceeIsInPreferredList { //Woman is indifferent to the courtier but she prefers her current fiancee is in her list of preferences.
		//So she sticks with her current fiancee
		fmt.Printf("Woman:%v is indifferent to the courtier:%v but her current fiancee:%v in in her preference list. So she sticks with her current fiancee\n", wp.Id, courtier.Id, wp.EngagedTo.Id)
		return false
	} else if courtierRanking < currentFianceeRanking {
		fmt.Printf("woman:%v prefers courtier:%v over current financee:%v\n", wp.Id, courtier.Id, wp.EngagedTo.Id)
		return true
	} else {
		fmt.Printf("woman:%v prefers current fiancee:%v over courtier:%v\n", wp.Id, wp.EngagedTo.Id, courtier.Id)
		return false
	}
}

func TestStableMatchingWomanConflicts2(t *testing.T) {
	w0 := &Woman{
		Id:          "0",
		Preferences: make(map[string]int),
		EngagedTo:   nil,
	}
	w1 := &Woman{
		Id:          "1",
		Preferences: make(map[string]int),
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

	w0.Preferences[m1.Id] = 1
	w0.Preferences[m0.Id] = 0
	w1.Preferences[m1.Id] = 0
	w1.Preferences[m0.Id] = 1
	actual := Match(allMen, womanPrefersMe)
	if w0.EngagedTo.Id != m0.Id {
		t.Errorf("Expected woman 0 to be engaged to man 0")
	}
	if w1.EngagedTo.Id != m1.Id {
		t.Errorf("Expected woman 1 to be engaged to man 1")
	}

	fmt.Println(actual)
}

func TestStableMatchingNoWomanPreferenceConflicts(t *testing.T) {
	w0 := &Woman{
		Id:          "0",
		Preferences: make(map[string]int),
		EngagedTo:   nil,
	}
	w1 := &Woman{
		Id:          "1",
		Preferences: make(map[string]int),
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

	w0.Preferences[m1.Id] = 0
	w0.Preferences[m0.Id] = 1
	w1.Preferences[m1.Id] = 1
	w1.Preferences[m0.Id] = 0
	actual := Match(allMen, womanPrefersMe)
	if w0.EngagedTo.Id != m0.Id {
		t.Errorf("Expected woman 0 to be engaged to man 0")
	}
	if w1.EngagedTo.Id != m1.Id {
		t.Errorf("Expected woman 1 to be engaged to man 1")
	}

	fmt.Println(actual)
}

func TestIndifferentStableMatching(t *testing.T) {
	w0 := &Woman{
		Id:          "0",
		Preferences: make(map[string]int),
		EngagedTo:   nil,
	}
	w1 := &Woman{
		Id:          "1",
		Preferences: make(map[string]int),
		EngagedTo:   nil,
	}

	w2 := &Woman{
		Id:          "2",
		Preferences: make(map[string]int),
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

	w0.Preferences[m1.Id] = 1
	w0.Preferences[m0.Id] = 0
	w1.Preferences[m1.Id] = 0
	w1.Preferences[m0.Id] = 1
	//w2 is indifferent to all men
	actual := Match(allMen, womanPrefersMe)
	if w0.EngagedTo.Id != m0.Id {
		t.Errorf("Expected woman 0 to be engaged to man 0")
	}
	if w1.EngagedTo.Id != m1.Id {
		t.Errorf("Expected woman 1 to be engaged to man 1")
	}
	if w2.EngagedTo.Id != m2.Id {
		t.Errorf("Expected woman 2 to be engaged to man 2")
	}

	fmt.Println(actual)
}
