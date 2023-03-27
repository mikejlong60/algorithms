package chapter4

import (
	"github.com/mikejlong60/algorithms/pkg/algorithms/chapter5"
)

func ScheduleAll(r []*TimeSlot) [][]*TimeSlot { //Each row of the returned array is the schedule for a single resource(say a thread)
	lt := func(l, r *TimeSlot) bool {
		if l.begin < r.begin {
			return true
		} else {
			return false
		}
	}
	rr := chapter5.MergeSort(r, lt)
	_, a := scheduleAll(rr, []*TimeSlot{})
	return a
}

func scheduleAll(remainingTimeSlots, scheduledThreads []*TimeSlot) ([]*TimeSlot, []*TimeSlot) {
	timeSlotsOverlap := func(x *TimeSlot, y *TimeSlot) bool {
		if x.begin < y.end {
			return true
		} else {
			return false //exclude b Timeslot because it overlaps with x
		}
	}

	if len(remainingTimeSlots) == 0 {
		return remainingTimeSlots, scheduledThreads
	} else {
		x := scheduledThreads[0]
		y := remainingTimeSlots[0]
		if !timeSlotsOverlap(x, y) { // Add x to scheduled thread
			scheduledThreads = append(scheduledThreads, y)
			remainingTimeSlots = remainingTimeSlots[1:]
		} else {
			remainingTimeSlots = remainingTimeSlots[x:]
			return scheduleAll(remainingTimeSlots, scheduledThreads)
		}
	}
}
