package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/linked_list"
)

//Perform the outer loop until every hospital has reached it's resident capacity.
//Inside that loop at first level make another loop that fills hospital resident capacity until complete.
type Hospital struct { //man
	Id                  string
	ResidentCapacity    int
	Residents           map[string]*Resident
	ResidentPreferences *linked_list.LinkedList[*Resident] //A stack of residents I want in order of preferences. When a resident is missing from it a hospital has already proposed to him.
}

type Resident struct { //woman
	Id          string
	Preferences []*Hospital //An array of hospital preferences for a resident. All the hospitals are in this array.
	Hospital    *Hospital   //The hospital the resident is currently assigned to
}

func MatchResidentToHospitals(hospitalsWithResidentOpenings *linked_list.LinkedList[*Hospital]) []*Resident {
	fmt.Printf("Size of list:%v\n", linked_list.Len(hospitalsWithResidentOpenings))

	//Invariants
	// 1. The sum of all hospital resident openings cannot exceed the total number of residents
	// 2. Due to preceding there can be more residents than there are resident openings at hospitals
	validateInvariants := func(hospitalsWithResidentOpenings *linked_list.LinkedList[*Hospital]) {
		xs := linked_list.ToArray(hospitalsWithResidentOpenings)
		var z int
		for _, y := range xs {
			z = z + y.ResidentCapacity //+ 1
		}
		if z > linked_list.Len(xs[0].ResidentPreferences) {
			panic("Cannot have more hospital resident openings than you have residents")
		}
	}
	if linked_list.Len(hospitalsWithResidentOpenings) == 0 {
		return []*Resident{}
	}

	validateInvariants(hospitalsWithResidentOpenings)

	allResidents := linked_list.ToArray(linked_list.Head(hospitalsWithResidentOpenings).ResidentPreferences)

	mEq := func(m1 *Hospital, m2 *Hospital) bool {
		if m1.Id == m2.Id {
			return true
		} else {
			return false
		}
	}

	residentPrefersThisHospital := func(res *Resident, thisHospital *Hospital) bool { //Does resident prefer this hospital to the one to which he is currently assigned?
		var result bool
		for _, m := range res.Preferences {
			if mEq(res.Hospital, thisHospital) || mEq(res.Hospital, m) {
				if mEq(m, thisHospital) {
					result = true
					break
				} else if mEq(m, res.Hospital) {
					result = false
					break
				}
			}
		}
		return result
	}

	for hospitalsWithResidentOpenings != nil {
		hospital := linked_list.Head(hospitalsWithResidentOpenings)
		var hospitalResidentPreferences = hospital.ResidentPreferences
		var emptyStackError error
		for hospitalResidentPreferences != nil && hospital.ResidentCapacity < len(hospital.Residents) { //Loop over the hospital's resident preferences until the hospital has reached it' capacity
			resident := linked_list.Head(hospitalResidentPreferences)
			if resident.Hospital == nil {
				resident.Hospital = hospital
				hospitalResidentPreferences, emptyStackError = linked_list.Tail(hospitalResidentPreferences)
				if emptyStackError != nil {
					fmt.Println(emptyStackError)
				}
				hospital.Residents[resident.Id] = resident
				//				break
			} else {
				//Does this resident prefer this hospital to the one to whom he is currently assigned? If so he
				//breaks his agreement to that hospital and you make hospital have an additional resident opening.
				//Otherwise just try the next hospital in the current resident's non-proposed-to(preferences) stack.
				if residentPrefersThisHospital(resident, hospital) {
					oldHospital := resident.Hospital
					delete(oldHospital.Residents, resident.Id)
					///Set up current resident with this hospital
					resident.Hospital = hospital
					hospitalsWithResidentOpenings = linked_list.AddLast(oldHospital, hospitalsWithResidentOpenings) //TODO Might have to check to see if this hospital is already in the hospitalsWithResidentOpenings list
					//fmt.Printf("Woman %v prefers me(%v) and just broke engagement to %v\n", resident.Id, hospital.Id, oldMan.Id)
				} else {
					//fmt.Printf("Woman %v prefers her current fiance %v over me (%v)\n", resident.Id, resident.EngagedTo.Id, hospital.Id)
				}
			}
			hospital.ResidentPreferences, emptyStackError = linked_list.Tail(hospital.ResidentPreferences)
			if emptyStackError != nil {
				fmt.Println(emptyStackError)
			}

		} //end resident for
		hospitalsWithResidentOpenings, emptyStackError = linked_list.Tail(hospitalsWithResidentOpenings)
		if emptyStackError != nil {
			fmt.Println(emptyStackError)
		}
	} // end hospital with resident openings for
	return allResidents
}
