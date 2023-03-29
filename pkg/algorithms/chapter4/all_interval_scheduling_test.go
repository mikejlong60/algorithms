package chapter4

import (
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

func TestAllIntervalScheduling1(t *testing.T) {
	a1 := TimeSlot{
		id:    1,
		begin: 0,
		end:   3,
	}
	a2 := TimeSlot{
		id:    2,
		begin: 0,
		end:   8,
	}
	a3 := TimeSlot{
		id:    3,
		begin: 0,
		end:   3,
	}
	a4 := TimeSlot{
		id:    4,
		begin: 5,
		end:   8,
	}
	a5 := TimeSlot{
		id:    5,
		begin: 5,
		end:   14,
	}
	a6 := TimeSlot{
		id:    6,
		begin: 10,
		end:   15,
	}
	a7 := TimeSlot{
		id:    7,
		begin: 10,
		end:   15,
	}
	a8 := TimeSlot{
		id:    8,
		begin: 14,
		end:   20,
	}
	a9 := TimeSlot{
		id:    9,
		begin: 16,
		end:   20,
	}
	a10 := TimeSlot{
		id:    10,
		begin: 16,
		end:   20,
	}

	eq := func(l, r *TimeSlot) bool {
		if l.begin == r.begin && l.end == r.end && l.id == r.id {
			return true
		} else {
			return false
		}
	}
	r := []*TimeSlot{&a1, &a2, &a3, &a4, &a5, &a6, &a7, &a8, &a9, &a10}
	actual := ScheduleAll(r)
	expected := [][]*TimeSlot{{&a3, &a5, &a10}, {&a2, &a7, &a9}, {&a1, &a4, &a6}, {&a8}}
	//TODO verify correctness of above answer
	for i, _ := range actual {
		if !arrays.ArrayEquality(actual[i], expected[i], eq) {
			t.Errorf("Actual:%v Expected:%v", actual, expected)
		}
	}
}
