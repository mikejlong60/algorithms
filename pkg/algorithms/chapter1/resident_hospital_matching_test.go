package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/linked_list"
	"testing"
)

func TestResidentHospitalMatching(t *testing.T) {
	res1 := &Resident{
		Id:          "1",
		Preferences: nil,
		Hospital:    nil,
	}
	res2 := &Resident{
		Id:          "2",
		Preferences: nil,
		Hospital:    nil,
	}
	res3 := &Resident{
		Id:          "3",
		Preferences: nil,
		Hospital:    nil,
	}
	res4 := &Resident{
		Id:          "4",
		Preferences: nil,
		Hospital:    nil,
	}
	res5 := &Resident{
		Id:          "5",
		Preferences: nil,
		Hospital:    nil,
	}

	hosp1 := &Hospital{
		Id:                  "1",
		ResidentCapacity:    2,
		Residents:           nil,
		ResidentPreferences: nil,
	}
	hosp2 := &Hospital{
		Id:                  "2",
		ResidentCapacity:    2,
		Residents:           nil,
		ResidentPreferences: nil,
	}

	res1.Preferences = []*Hospital{hosp1, hosp2}
	res2.Preferences = []*Hospital{hosp1, hosp2}
	res3.Preferences = []*Hospital{hosp2, hosp1}
	res4.Preferences = []*Hospital{hosp1, hosp2}
	res5.Preferences = []*Hospital{hosp2, hosp1}

	var hospitalsWithResidentOpenings *linked_list.LinkedList[*Hospital]
	var hosp1ResPrefs *linked_list.LinkedList[*Resident]
	hosp1ResPrefs = linked_list.Push(res1, hosp1ResPrefs)
	hosp1ResPrefs = linked_list.Push(res2, hosp1ResPrefs)
	hosp1ResPrefs = linked_list.Push(res3, hosp1ResPrefs)
	hosp1ResPrefs = linked_list.Push(res4, hosp1ResPrefs)
	hosp1ResPrefs = linked_list.Push(res5, hosp1ResPrefs)

	hosp1.ResidentPreferences = hosp1ResPrefs
	hosp2.ResidentPreferences = hosp1ResPrefs //TODO Fix this to be different
	hospitalsWithResidentOpenings = linked_list.Push(hosp1, hospitalsWithResidentOpenings)
	hospitalsWithResidentOpenings = linked_list.Push(hosp2, hospitalsWithResidentOpenings)
	match := MatchResidentToHospitals(hospitalsWithResidentOpenings)
	fmt.Println(match)
	t.Errorf("fail")
}
