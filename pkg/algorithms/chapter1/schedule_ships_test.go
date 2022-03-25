package chapter1

import (
	"fmt"
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
		ProposedSchedule: []interface{}{Port{Id: 1}, AtSea{}},
	}
	var s2 = Ship{Id: 2,
		ProposedSchedule: []interface{}{Port{Id: 2}, AtSea{}},
	}
	r := []Ship{s1, s2}

	schedule(r)
	fmt.Println(schedule(r))
}
