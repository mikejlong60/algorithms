package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

func TestScheduleShips(t *testing.T) {

	//ship1
	////w2 is indifferent to all men
	//actual := Match(allMen, womanPrefersMe)
	//if w0.EngagedTo.Id != m0.Id {
	//	t.Errorf("Expected woman 0 to be engaged to man 0")
	//}
	//if w1.EngagedTo.Id != m1.Id {
	//	t.Errorf("Expected woman 1 to be engaged to man 1")
	//}
	//if w2.EngagedTo.Id != m2.Id {
	//	t.Errorf("Expected woman 2 to be engaged to man 2")
	//}
	var s1 = Ship{Id: 1,
		ProposedSchedule: []int{1, 0},
	}
	var s2 = Ship{Id: 2,
		ProposedSchedule: []int{2, 0},
	}
	masterSchedule := []*Ship{&s1, &s2}

	actual := schedule(masterSchedule)

	fmt.Printf("actual[0]:%v\n", actual[0].ActualSchedule)
	fmt.Printf("actual[1]:%v\n", actual[1].ActualSchedule)
	p := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	if !arrays.SetEquality(actual[0].ActualSchedule, []int{1, 0}, p) {
		t.Errorf("Actual:%v, Expected:%v", actual[0].ActualSchedule, []int{1, 0})
	}
	if !arrays.SetEquality(actual[1].ActualSchedule, []int{2, 0}, p) {
		t.Errorf("Actual:%v, Expected:%v", actual[1].ActualSchedule, []int{2, 0})
	}

}
