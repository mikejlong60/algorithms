package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/linked_list"
	"github.com/greymatter-io/golangz/propcheck"
)

//This algorithm accomodates both weak and strong instabilities.
type Man struct {
	Id          string
	Preferences *linked_list.LinkedList[*Woman] //A stack of women I want in order of preferences. When a woman is missing from it he has already proposed to her.
	EngagedTo   *Woman
}

type Woman struct {
	Id          string
	Preferences map[string]propcheck.Pair[int, *Man] // The key is the man's Id and the value is that man's ranking with 0 being the highest and a pointer to the complete Man.  No duplicate rankings are allowed.
	EngagedTo   *Man
}

var womanPrefersMe = func(wp *Woman, courtier *Man) bool { //Does woman prefer this man to the one to which she is currently assigned?
	//This function assumes that the wp woman is already engaged.
	courtierRanking, courtierIsInPreferredList := wp.Preferences[courtier.Id]
	currentFianceeRanking, currentFianceeIsInPreferredList := wp.Preferences[wp.EngagedTo.Id]

	if !courtierIsInPreferredList && !currentFianceeIsInPreferredList { //Woman is indifferent to both men. So she will stick with her current fiancee.
		//fmt.Printf("Woman:%v is indifferent to both men. So she will stick with her current fiancee%v\n", wp.Id, wp.EngagedTo.Id)
		return false
	} else if courtierIsInPreferredList && !currentFianceeIsInPreferredList { //Woman is indifferent to her current husband but not the courtier. So she chooses the courtier.
		//fmt.Printf("Woman:%v is indifferent to her current husband:%v but not the courtier:%v. So she chooses the courtier.\n", wp.Id, wp.EngagedTo.Id, courtier.Id)
		return true
	} else if !courtierIsInPreferredList && currentFianceeIsInPreferredList { //Woman is indifferent to the courtier but she prefers her current fiancee is in her list of preferences.
		//So she sticks with her current fiancee
		//fmt.Printf("Woman:%v is indifferent to the courtier:%v but her current fiancee:%v in in her preference list. So she sticks with her current fiancee\n", wp.Id, courtier.Id, wp.EngagedTo.Id)
		return false
	} else if courtierRanking.A < currentFianceeRanking.A {
		//fmt.Printf("woman:%v prefers courtier:%v over current financee:%v\n", wp.Id, courtier.Id, wp.EngagedTo.Id)
		return true
	} else {
		//fmt.Printf("woman:%v prefers current fiancee:%v over courtier:%v\n", wp.Id, wp.EngagedTo.Id, courtier.Id)
		return false
	}
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
					//fmt.Printf("Woman %v prefers me(%v) and just broke engagement to %v\n", wp.Id, m.Id, oldMan.Id)
					break
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
