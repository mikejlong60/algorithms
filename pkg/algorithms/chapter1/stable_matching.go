package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/linked_list"
)

//This algorithm accomodates both weak and strong instabilities.
type Man struct {
	Id          string
	Preferences *linked_list.LinkedList[*Woman] //A stack of women I want in order of preferences. When a woman is missing from it he has already proposed to her.
	EngagedTo   *Woman
}

type Woman struct {
	Id           string
	pPreferences []*Man
	Preferences  map[string]int
	EngagedTo    *Man
}

func Match(freeMen *linked_list.LinkedList[*Man], wPrefersMe func(wp *Woman, me *Man) bool) []*Woman {
	fmt.Printf("Size of list:%v\n", linked_list.Len(freeMen))
	if linked_list.Len(freeMen) == 0 {
		return []*Woman{}
	}
	allWomen := linked_list.ToArray(linked_list.Head(freeMen).Preferences) //Every man must have every woman in his list of preferences

	for freeMen != nil {
		m := linked_list.Head(freeMen)
		for m.Preferences != nil {
			wp := linked_list.Head(m.Preferences)
			if wp.EngagedTo == nil {
				wp.EngagedTo = m
				m.EngagedTo = wp
				break
			} else {
				//Does this woman prefer me to whom she is currently engaged? If so she
				//breaks her engagement to that guy and you add that guy to free men.
				//Otherwise just try the next woman in the current man's non-proposed-to(preferences) stack.
				if wPrefersMe(wp, m) {
					oldMan := wp.EngagedTo
					oldMan.EngagedTo = nil
					///Set up current man with this woman
					wp.EngagedTo = m
					m.EngagedTo = wp
					freeMen = linked_list.AddLast(oldMan, freeMen)
					fmt.Printf("Woman %v prefers me(%v) and just broke engagement to %v\n", wp.Id, m.Id, oldMan.Id)
					break
				} else {
					fmt.Printf("Woman %v prefers her current fiance %v over me (%v)\n", wp.Id, wp.EngagedTo.Id, m.Id)
				}
			}
			m.Preferences, _ = linked_list.Tail(m.Preferences)
		} //end woman for
		freeMen, _ = linked_list.Tail(freeMen)
	} // end man for
	return allWomen
}
