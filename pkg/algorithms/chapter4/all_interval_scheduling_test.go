package chapter4

import (
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

func TestAllIntervalScheduling1(t *testing.T) {
	a := TimeSlot{
		id:    1,
		begin: 0,
		end:   3,
	}
	b := TimeSlot{
		id:    2,
		begin: 0,
		end:   8,
	}
	c := TimeSlot{
		id:    3,
		begin: 0,
		end:   3,
	}
	d := TimeSlot{
		id:    4,
		begin: 5,
		end:   8,
	}
	e := TimeSlot{
		id:    5,
		begin: 5,
		end:   14,
	}
	f := TimeSlot{
		id:    6,
		begin: 10,
		end:   15,
	}
	g := TimeSlot{
		id:    7,
		begin: 10,
		end:   15,
	}
	h := TimeSlot{
		id:    8,
		begin: 14,
		end:   20,
	}
	i := TimeSlot{
		id:    9,
		begin: 16,
		end:   20,
	}
	j := TimeSlot{
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
	r := []*TimeSlot{&a, &b, &c, &d, &e, &f, &g, &h, &i, &j}
	actual := ScheduleAll(r)
	expected := [][]*TimeSlot{{&a, &e, &h}, {&b, &g, &j}, {&c, &d, &f, &i}}
	for i, _ := range actual {
		if !arrays.ArrayEquality(actual[i], expected[i], eq) {
			t.Errorf("Actual:%v Expected:%v", actual, expected)
		}
	}
}
