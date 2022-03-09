package chapter1

import (
	"fmt"
	"github.com/mikejlong60/golangz/pkg/linked_list"
)

type Man struct {
	Id          string
	Preferences *linked_list.ConsList[*Woman] //A stack of women I want in order of preferences. When a woman is missing from it he has already proposed to her.
}

type Woman struct {
	Id          string
	Preferences []*Man
	EngagedTo   *Man
}

func Match(freeMen *linked_list.ConsList[*Man]) []*Woman {
	fmt.Printf("Size of list:%v\n", linked_list.Len(freeMen))
	if linked_list.Len(freeMen) == 0 {
		return []*Woman{}
	}
	allWomen := linked_list.ToArray(linked_list.Head(freeMen).Preferences)

	mEq := func(m1 *Man, m2 *Man) bool {
		if m1.Id == m2.Id {
			return true
		} else {
			return false
		}
	}

	wPrefersMe := func(wp *Woman, me *Man) bool { //Does wp prefer m to whom she is currently engaged
		var result bool
		for _, m := range wp.Preferences {
			if mEq(wp.EngagedTo, me) || mEq(wp.EngagedTo, m) {
				if mEq(m, me) {
					result = true
					break
				} else if mEq(m, wp.EngagedTo) {
					result = false
					break
				}
			}
		}
		return result
	}

	for freeMen != nil {
		m := linked_list.Head(freeMen)
		for m.Preferences != nil {
			wp := linked_list.Head(m.Preferences)
			if wp.EngagedTo == nil {
				wp.EngagedTo = m
				break
			} else {
				//Does this woman prefer me to whom she is currently engaged? If so she
				//breaks her engagement to that guy and you add that guy to free men.
				//Otherwise just try the next woman in the current man's non-proposed-to(preferences) stack.
				if wPrefersMe(wp, m) {
					oldMan := wp.EngagedTo
					///Set up current man with this woman
					wp.EngagedTo = m
					freeMen = linked_list.AddLast(oldMan, freeMen)
					//fmt.Printf("Woman %v prefers me(%v) and just broke engagement to %v\n", wp.Id, m.Id, oldMan.Id)
				} else {
					//fmt.Printf("Woman %v prefers her current fiance %v over me (%v)\n", wp.Id, wp.EngagedTo.Id, m.Id)
				}
			}
			m.Preferences, _ = linked_list.Tail(m.Preferences)
		} //end woman for
		freeMen, _ = linked_list.Tail(freeMen)
	} // end man for
	return allWomen
}
