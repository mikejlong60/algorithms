package chapter1

import (
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

//TODO make all these tests a single properties test once you make a set generator in Golangz
func TestScheduleShips1(t *testing.T) {
	var s1 = Ship{Id: 1,
		ProposedSchedule: []int{1, 0, 2},
	}
	var s2 = Ship{Id: 2,
		ProposedSchedule: []int{1, 0, 2},
	}
	masterSchedule := []*Ship{&s1, &s2}

	actual := schedule(masterSchedule)

	p := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	if !arrays.SetEquality(actual[0].ActualSchedule, []int{}, p) {
		t.Errorf("Actual:%v, Expected:%v", actual[0].ActualSchedule, []int{1, 0})
	}
	if !arrays.SetEquality(actual[1].ActualSchedule, []int{1, 0, 2}, p) {
		t.Errorf("Actual:%v, Expected:%v", actual[1].ActualSchedule, []int{2, 0})
	}
}

func TestScheduleShips2(t *testing.T) {
	var s1 = Ship{Id: 1,
		ProposedSchedule: []int{1, 0, 2, 0},
	}
	var s2 = Ship{Id: 2,
		ProposedSchedule: []int{1, 0, 2, 0},
	}
	masterSchedule := []*Ship{&s1, &s2}

	actual := schedule(masterSchedule)

	p := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	if !arrays.SetEquality(actual[0].ActualSchedule, []int{}, p) {
		t.Errorf("Actual:%v, Expected:%v", actual[0].ActualSchedule, []int{1, 0})
	}
	if !arrays.SetEquality(actual[1].ActualSchedule, []int{1, 0, 2, 0}, p) {
		t.Errorf("Actual:%v, Expected:%v", actual[1].ActualSchedule, []int{2, 0})
	}
}

func TestScheduleShips3(t *testing.T) {
	var s1 = Ship{Id: 1,
		ProposedSchedule: []int{1, 0, 3, 0},
	}
	var s2 = Ship{Id: 2,
		ProposedSchedule: []int{1, 0, 2, 0},
	}
	masterSchedule := []*Ship{&s1, &s2}

	actual := schedule(masterSchedule)

	p := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	if !arrays.SetEquality(actual[0].ActualSchedule, []int{}, p) {
		t.Errorf("Actual:%v, Expected:%v", actual[0].ActualSchedule, []int{1, 0})
	}
	if !arrays.SetEquality(actual[1].ActualSchedule, []int{1, 0, 2, 0}, p) {
		t.Errorf("Actual:%v, Expected:%v", actual[1].ActualSchedule, []int{2, 0})
	}
}

func TestScheduleShips4(t *testing.T) {
	var s0 = Ship{Id: 1,
		ProposedSchedule: []int{3, 0, 2, 0},
	}
	var s1 = Ship{Id: 1,
		ProposedSchedule: []int{3, 0, 2, 0},
	}
	var s2 = Ship{Id: 2,
		ProposedSchedule: []int{1, 0, 2, 0},
	}
	masterSchedule := []*Ship{&s0, &s1, &s2}

	actual := schedule(masterSchedule)

	p := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	if !arrays.SetEquality(actual[0].ActualSchedule, []int{}, p) {
		t.Errorf("Actual:%v, Expected:%v", actual[0].ActualSchedule, []int{1, 0})
	}
	if !arrays.SetEquality(actual[1].ActualSchedule, []int{3, 0}, p) {
		t.Errorf("Actual:%v, Expected:%v", actual[0].ActualSchedule, []int{1, 0})
	}
	if !arrays.SetEquality(actual[2].ActualSchedule, []int{1, 0, 2, 0}, p) {
		t.Errorf("Actual:%v, Expected:%v", actual[1].ActualSchedule, []int{2, 0})
	}
}

func TestScheduleShips5(t *testing.T) {
	var s0 = Ship{Id: 1,
		ProposedSchedule: []int{4, 0, 2, 0},
	}
	var s1 = Ship{Id: 1,
		ProposedSchedule: []int{3, 0, 2, 0},
	}
	var s2 = Ship{Id: 2,
		ProposedSchedule: []int{1, 0, 2, 0},
	}
	masterSchedule := []*Ship{&s0, &s1, &s2}

	actual := schedule(masterSchedule)

	p := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	if !arrays.SetEquality(actual[0].ActualSchedule, []int{4, 0}, p) {
		t.Errorf("Actual:%v, Expected:%v", actual[0].ActualSchedule, []int{1, 0})
	}
	if !arrays.SetEquality(actual[1].ActualSchedule, []int{3, 0}, p) {
		t.Errorf("Actual:%v, Expected:%v", actual[0].ActualSchedule, []int{1, 0})
	}
	if !arrays.SetEquality(actual[2].ActualSchedule, []int{1, 0, 2, 0}, p) {
		t.Errorf("Actual:%v, Expected:%v", actual[1].ActualSchedule, []int{2, 0})
	}
}
