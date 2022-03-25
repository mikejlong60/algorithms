package chapter1

//This algorithm accomodates both weak and strong instabilities.
type Ship struct {
	Id               int
	ProposedSchedule []interface{}
	ActualSchedule   []interface{}
}

type ShipState interface{ AtSea | Port }

type AtSea struct{}
type Port struct {
	Id int
}

var s1 = Ship{Id: 1,
	ProposedSchedule: []interface{}{Port{Id: 1}, AtSea{}},
}

//Algorithm
//1. range loop over array of ships, index i
//2. For current ship range loop ov all other ships and make sure there is not a port conflict with element i in any of their proposed schedules
//2.1  If there no port  conflicts, add schedule to end of ActualSchedule. A port conflict is same port at same element. At sea is no conflict
//2.2  Otherwise skip to next element in array of ships
//3 Return array of ships with ActualSchedule
func schedule(ships []Ship) []Ship {

	return ships

}

//
//import (
//"fmt"
//"constraints"
//)

//func max[T Ordered](x, y T) T {
//	if x > y {
//		return x
//	}
//
//	return y
//}
//
//func main() {
//	fmt.Println(max(99, 11)) // 99
//}

//func MakeMasterSchedule(ships []Ship) {}

//, wPrefersMe func(wp *Woman, me *Man) bool) []*Woman {
//	fmt.Printf("Size of list:%v\n", linked_list.Len(freeMen))
//	if linked_list.Len(freeMen) == 0 {
//		return []*Woman{}
//	}
//	allWomen := linked_list.ToArray(linked_list.Head(freeMen).Preferences) //Every man must have every woman in his list of preferences
//
//	for freeMen != nil {
//		m := linked_list.Head(freeMen)
//		for m.Preferences != nil {
//			wp := linked_list.Head(m.Preferences)
//			if wp.EngagedTo == nil {
//				wp.EngagedTo = m
//				m.EngagedTo = wp
//				break
//			} else {
//				//Does this woman prefer me to whom she is currently engaged? If so she
//				//breaks her engagement to that guy and you add that guy to free men.
//				//Otherwise just try the next woman in the current man's non-proposed-to(preferences) stack.
//				if wPrefersMe(wp, m) {
//					oldMan := wp.EngagedTo
//					oldMan.EngagedTo = nil
//					///Set up current man with this woman
//					wp.EngagedTo = m
//					m.EngagedTo = wp
//					freeMen = linked_list.AddLast(oldMan, freeMen)
//					fmt.Printf("Woman %v prefers me(%v) and just broke engagement to %v\n", wp.Id, m.Id, oldMan.Id)
//					break
//				} else {
//					fmt.Printf("Woman %v prefers her current fiance %v over me (%v)\n", wp.Id, wp.EngagedTo.Id, m.Id)
//				}
//			}
//			m.Preferences, _ = linked_list.Tail(m.Preferences)
//		} //end woman for
//		freeMen, _ = linked_list.Tail(freeMen)
//	} // end man for
//	return allWomen
//}
