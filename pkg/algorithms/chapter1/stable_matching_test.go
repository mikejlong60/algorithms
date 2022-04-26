package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/linked_list"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/hashicorp/go-multierror"
	"math/rand"
	"testing"
	"time"
)

func shuffleAny[A any](toBeShuffled []*A) []*A {
	rr := make([]*A, len(toBeShuffled))
	copy(rr, toBeShuffled)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(rr), func(i, j int) {
		rr[i], rr[j] = rr[j], rr[i]
	})
	return rr
}

func TestStableMatchingPropertiesTest(t *testing.T) {
	var allTheMen []*Man
	var allTheWomen []*Woman
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	g0 := propcheck.ChooseInt(0, 100)
	fa := func(a int) func(propcheck.SimpleRNG) (*linked_list.LinkedList[*Man], propcheck.SimpleRNG) {
		allTheMen = []*Man{}
		allTheWomen = []*Woman{}
		mw := func(mIds []string, wIds []string) *linked_list.LinkedList[*Man] {
			if len(allTheWomen) != len(allTheMen) {
				t.Error("length of men and women arrays were not equal")
			}
			for _, s := range mIds {
				allTheMen = append(allTheMen, &Man{Id: s})
			}
			for _, s := range wIds {
				allTheWomen = append(allTheWomen, &Woman{Id: s})
			}

			//Make two arrays, one of shuffled men, one for each woman, and one of  shuffled women, one for each man.
			var freeMen *linked_list.LinkedList[*Man]
			for _, s := range allTheMen {
				freeMen = linked_list.Push(s, freeMen)
			}

			for i, _ := range allTheWomen {
				womenForMan := shuffleAny(allTheWomen)
				var allWomen *linked_list.LinkedList[*Woman]
				for _, s := range womenForMan {
					allWomen = linked_list.Push(s, allWomen)
				}
				allTheMen[i].HaveNotProposedTo = allWomen
				allTheMen[i].Preferences = linked_list.ToArray(allWomen)
			}

			for _, s := range allTheWomen {
				wpref := shuffleAny(allTheMen)
				var wprefMap = make(map[string]propcheck.Pair[int, *Man], len(wpref))
				for i, m := range wpref {
					wprefMap[m.Id] = propcheck.Pair[int, *Man]{i, m}
				}
				s.Preferences = wprefMap
			}

			return freeMen
		}
		ra := propcheck.ArrayOfN(a, propcheck.String(100))
		rb := propcheck.ArrayOfN(a, propcheck.String(100))
		return propcheck.Map2(ra, rb, mw)
	}

	g := propcheck.FlatMap(g0, fa)

	prop := propcheck.ForAll(g,
		"Make a bunch of men and women and match them up and see if all matches are stable \n",
		func(freeMen *linked_list.LinkedList[*Man]) []*Woman {
			len := linked_list.Len(freeMen)
			start := time.Now()
			r := Match(freeMen, womanPrefersMe)
			fmt.Printf("Match took %v for %v couples\n", time.Since(start), len)
			return r
		},
		func(allWomen []*Woman) (bool, error) {
			var errors error

			var allHusbandIds []string
			for _, j := range allWomen {
				if j.EngagedTo == nil {
					errors = multierror.Append(errors, fmt.Errorf("Woman:%v was not married ", j.Id))
				}
				allHusbandIds = append(allHusbandIds, j.EngagedTo.Id)
			}

			var allMenIds []string
			for _, man := range allTheMen {
				allMenIds = append(allMenIds, man.Id)
			}

			mEq := func(m1 string, m2 string) bool {
				if m1 == m2 {
					return true
				} else {
					return false
				}
			}
			if !arrays.SetEquality(allMenIds, allHusbandIds, mEq) {
				errors = multierror.Append(errors, fmt.Errorf("All men were not married"))
				fmt.Printf("These men never got married:%v\n", arrays.SetMinus(allMenIds, allHusbandIds, mEq))
			}
			unstableMatchings := unstableMatchings(allWomen)
			if len(unstableMatchings) > 0 {
				errors = multierror.Append(errors, fmt.Errorf("There were unstable matchings as follow:%v", unstableMatchings))
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{10, rng})
	propcheck.ExpectSuccess[*linked_list.LinkedList[*Man]](t, result)
	fmt.Println(rng)
}

func TestStableMatchingWithWomanWhoAreIndifferentAboutSomeMen(t *testing.T) {
	var allTheMen []*Man
	var allTheWomen []*Woman
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	g0 := propcheck.ChooseInt(0, 100)
	fa := func(a int) func(propcheck.SimpleRNG) (*linked_list.LinkedList[*Man], propcheck.SimpleRNG) {
		allTheMen = []*Man{}
		allTheWomen = []*Woman{}
		mw := func(mIds []string, wIds []string) *linked_list.LinkedList[*Man] {
			if len(allTheWomen) != len(allTheMen) {
				t.Error("length of men and women arrays were not equal")
			}
			for _, s := range mIds {
				allTheMen = append(allTheMen, &Man{Id: s})
			}
			for _, s := range wIds {
				allTheWomen = append(allTheWomen, &Woman{Id: s})
			}

			//Make two arrays, one of shuffled men, one for each woman, and one of  shuffled women, one for each man.
			var freeMen *linked_list.LinkedList[*Man]
			for _, s := range allTheMen {
				freeMen = linked_list.Push(s, freeMen)
			}

			for i, _ := range allTheWomen {
				womenForMan := shuffleAny(allTheWomen)
				var allWomen *linked_list.LinkedList[*Woman]
				for _, s := range womenForMan {
					allWomen = linked_list.Push(s, allWomen)
				}
				allTheMen[i].HaveNotProposedTo = allWomen
				allTheMen[i].Preferences = linked_list.ToArray(allWomen)
			}

			for _, s := range allTheWomen {
				wpref := shuffleAny(allTheMen)
				var wprefMap = make(map[string]propcheck.Pair[int, *Man], len(wpref)/2)
				for i, m := range wpref {
					wprefMap[m.Id] = propcheck.Pair[int, *Man]{i, m}
					if i > len(wpref)/2 { //Make every woman indifferent to half the men by excluding them from her preference list
						break
					}
				}
				s.Preferences = wprefMap
			}

			return freeMen
		}
		ra := propcheck.ArrayOfN(a, propcheck.String(100))
		rb := propcheck.ArrayOfN(a, propcheck.String(100))
		return propcheck.Map2(ra, rb, mw)
	}

	g := propcheck.FlatMap(g0, fa)

	prop := propcheck.ForAll(g,
		"Make a bunch of men and women where the women are undecided about some men and match them up and verify that all matches are stable \n",
		func(freeMen *linked_list.LinkedList[*Man]) []*Woman {
			len := linked_list.Len(freeMen)
			start := time.Now()
			r := Match(freeMen, womanPrefersMe)
			fmt.Printf("Match took %v for %v couples\n", time.Since(start), len)
			return r
		},
		func(allWomen []*Woman) (bool, error) {
			var errors error

			var allHusbandIds []string
			for _, j := range allWomen {
				if j.EngagedTo == nil {
					errors = multierror.Append(errors, fmt.Errorf("Woman:%v was not married ", j.Id))
				}
				allHusbandIds = append(allHusbandIds, j.EngagedTo.Id)
			}

			var allMenIds []string
			for _, man := range allTheMen {
				allMenIds = append(allMenIds, man.Id)
			}

			mEq := func(m1 string, m2 string) bool {
				if m1 == m2 {
					return true
				} else {
					return false
				}
			}
			if !arrays.SetEquality(allMenIds, allHusbandIds, mEq) {
				errors = multierror.Append(errors, fmt.Errorf("All men were not married"))
				fmt.Printf("These men never got married:%v\n", arrays.SetMinus(allMenIds, allHusbandIds, mEq))
			}
			unstableMatchings := unstableMatchings(allWomen)
			if len(unstableMatchings) > 0 {
				errors = multierror.Append(errors, fmt.Errorf("There were unstable matchings as follow:%v", unstableMatchings))
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{10, rng})
	propcheck.ExpectSuccess[*linked_list.LinkedList[*Man]](t, result)
	fmt.Println(rng)
}

//Algorithm for determining unstable matchings
//   Make a new list r of strings that you will return as the potential list of Unstable Matchings
//   For1 each woman w from all women
//        grab w's husband as m
//        make a new list ipw of potential instabilities
//        For2 each man m2 in woman w's preferences
//            if m2 ranks w above his current woman w2?
//                add m2 to ipw as a candidate for instability for woman w
//            end if
//        End For2
//        For3 each ipw as m3
//            if m3 ranks above w's current husband
//                 log that notation that w prefers m3 and m3 prefers w. This is in comparison to m.  Log all three of these values
//            end if
//        End For3
//    End For1

//If len(r) > 1 you have unstable matching and you should print list to console

func unstableMatchings(allWomen []*Woman) []string {
	var unstableMatchings []string
	var womanRanking = func(w *Woman, currentWomanPreferences []*Woman, manId string) int {
		for k, w2 := range currentWomanPreferences {
			if w.Id == w2.Id {
				return k
			}
		}
		return len(currentWomanPreferences) - 1 //Return lowest possible ranking if man has no preference for his current woman
	}
	for _, w := range allWomen { //for1
		m := w.EngagedTo
		var ipw []*Man
		var wPrefs []string
		for _, k := range w.Preferences {
			wPrefs = append(wPrefs, k.B.Id)
		}
		for _, m2 := range w.Preferences { //for2
			//Get man for m2 id
			//Then determine for each man if new woman w ranks above his current woman in his preferences list versus the woman to whom he is currently engaged and add that
			//man to the list of instability candidates for evaluation in the next for loop.
			if womanRanking(w, m2.B.Preferences, m2.B.Id) < womanRanking(m2.B.EngagedTo, m2.B.Preferences, m2.B.Id) { //lower is a higher preference
				ipw = append(ipw, m2.B)
			}
		}
		//end for2
		for _, m3 := range ipw { //for3
			if w.Preferences[m3.Id].A < w.Preferences[m.Id].A { //lower on the perference list is better
				unstableMatchings = append(unstableMatchings, fmt.Sprintf("Woman:%v prefers Man:%v over her current husband:%v and this is an instablity", w.Id, m3.Id, m.Id))
			}
		} //end for3
	} //end for1
	return unstableMatchings
}

//func TestStableMatchingWomanConflictsNoIndifference(t *testing.T) {
//	w0 := &Woman{
//		Id:          "0",
//		Preferences: make(map[string]propcheck.Pair[int, *Man]),
//		EngagedTo:   nil,
//	}
//	w1 := &Woman{
//		Id:          "1",
//		Preferences: make(map[string]propcheck.Pair[int, *Man]),
//		EngagedTo:   nil,
//	}
//
//	m0 := &Man{
//		Id:                "0",
//		HaveNotProposedTo: nil,
//	}
//	m1 := &Man{
//		Id:                "1",
//		HaveNotProposedTo: nil,
//	}
//
//	w0.Preferences[m1.Id] = propcheck.Pair[int, *Man]{0, m1}
//	w0.Preferences[m0.Id] = propcheck.Pair[int, *Man]{1, m0}
//	w1.Preferences[m1.Id] = propcheck.Pair[int, *Man]{0, m1}
//	w1.Preferences[m0.Id] = propcheck.Pair[int, *Man]{1, m0}
//
//	var allMen *linked_list.LinkedList[*Man]
//	var manPreferences *linked_list.LinkedList[*Woman]
//	manPreferences = linked_list.Push(w1, manPreferences)
//	manPreferences = linked_list.Push(w0, manPreferences)
//
//	m0.HaveNotProposedTo = manPreferences
//	m1.HaveNotProposedTo = manPreferences
//	m0.Preferences = linked_list.ToArray(manPreferences)
//	m1.Preferences = linked_list.ToArray(manPreferences)
//	allMen = linked_list.Push(m1, allMen)
//	allMen = linked_list.Push(m0, allMen)
//	Match(allMen, womanPrefersMe)
//	if w0.EngagedTo.Id != m1.Id {
//		t.Errorf("Expected woman 0 to be engaged to man 1")
//	}
//	if w1.EngagedTo.Id != m0.Id {
//		t.Errorf("Expected woman 1 to be engaged to man 0")
//	}
//	unstableMatchings := unstableMatchings([]*Woman{w0, w1})
//	if len(unstableMatchings) > 0 {
//		t.Errorf("There were unstable matchings as follow:%v", unstableMatchings)
//	}
//}
//
//func TestStableMatchingWomanConflicts2(t *testing.T) {
//	w0 := &Woman{
//		Id:          "0",
//		Preferences: make(map[string]propcheck.Pair[int, *Man]),
//		EngagedTo:   nil,
//	}
//	w1 := &Woman{
//		Id:          "1",
//		Preferences: make(map[string]propcheck.Pair[int, *Man]),
//		EngagedTo:   nil,
//	}
//
//	m0 := &Man{
//		Id:                "0",
//		HaveNotProposedTo: nil,
//	}
//	m1 := &Man{
//		Id:                "1",
//		HaveNotProposedTo: nil,
//	}
//
//	var allMen *linked_list.LinkedList[*Man]
//	var manPreferences *linked_list.LinkedList[*Woman]
//	manPreferences = linked_list.Push(w1, manPreferences)
//	manPreferences = linked_list.Push(w0, manPreferences)
//
//	m0.HaveNotProposedTo = manPreferences
//	m1.HaveNotProposedTo = manPreferences
//	m0.Preferences = linked_list.ToArray(manPreferences)
//	m1.Preferences = linked_list.ToArray(manPreferences)
//
//	allMen = linked_list.Push(m1, allMen)
//	allMen = linked_list.Push(m0, allMen)
//
//	w0.Preferences[m1.Id] = propcheck.Pair[int, *Man]{1, m1}
//	w0.Preferences[m0.Id] = propcheck.Pair[int, *Man]{0, m0}
//	w1.Preferences[m1.Id] = propcheck.Pair[int, *Man]{0, m1}
//	w1.Preferences[m0.Id] = propcheck.Pair[int, *Man]{1, m0}
//	Match(allMen, womanPrefersMe)
//	if w0.EngagedTo.Id != m0.Id {
//		t.Errorf("Expected woman 0 to be engaged to man 0")
//	}
//	if w1.EngagedTo.Id != m1.Id {
//		t.Errorf("Expected woman 1 to be engaged to man 1")
//	}
//
//	unstableMatchings := unstableMatchings([]*Woman{w0, w1})
//	if len(unstableMatchings) > 0 {
//		t.Errorf("There were unstable matchings as follow:%v", unstableMatchings)
//	}
//}
//
//func TestStableMatchingNoWomanPreferenceConflicts(t *testing.T) {
//	w0 := &Woman{
//		Id:          "0",
//		Preferences: make(map[string]propcheck.Pair[int, *Man]),
//		EngagedTo:   nil,
//	}
//	w1 := &Woman{
//		Id:          "1",
//		Preferences: make(map[string]propcheck.Pair[int, *Man]),
//		EngagedTo:   nil,
//	}
//
//	m0 := &Man{
//		Id:                "0",
//		HaveNotProposedTo: nil,
//	}
//	m1 := &Man{
//		Id:                "1",
//		HaveNotProposedTo: nil,
//	}
//
//	var allMen *linked_list.LinkedList[*Man]
//	var man0Preferences *linked_list.LinkedList[*Woman]
//	man0Preferences = linked_list.Push(w1, man0Preferences)
//	man0Preferences = linked_list.Push(w0, man0Preferences)
//	var man1Preferences *linked_list.LinkedList[*Woman]
//	man1Preferences = linked_list.Push(w0, man1Preferences)
//	man1Preferences = linked_list.Push(w1, man1Preferences)
//
//	m0.HaveNotProposedTo = man0Preferences
//	m1.HaveNotProposedTo = man1Preferences
//	m0.Preferences = linked_list.ToArray(man0Preferences)
//	m1.Preferences = linked_list.ToArray(man1Preferences)
//
//	allMen = linked_list.Push(m1, allMen)
//	allMen = linked_list.Push(m0, allMen)
//
//	w0.Preferences[m1.Id] = propcheck.Pair[int, *Man]{0, m1}
//	w0.Preferences[m0.Id] = propcheck.Pair[int, *Man]{1, m0}
//	w1.Preferences[m1.Id] = propcheck.Pair[int, *Man]{1, m1}
//	w1.Preferences[m0.Id] = propcheck.Pair[int, *Man]{0, m0}
//	Match(allMen, womanPrefersMe)
//	if w0.EngagedTo.Id != m0.Id {
//		t.Errorf("Expected woman 0 to be engaged to man 0")
//	}
//	if w1.EngagedTo.Id != m1.Id {
//		t.Errorf("Expected woman 1 to be engaged to man 1")
//	}
//	unstableMatchings := unstableMatchings([]*Woman{w0, w1})
//	if len(unstableMatchings) > 0 {
//		t.Errorf("There were unstable matchings as follow:%v", unstableMatchings)
//	}
//}
//
//func TestIndifferentStableMatching(t *testing.T) {
//	w0 := &Woman{
//		Id:          "0",
//		Preferences: make(map[string]propcheck.Pair[int, *Man]),
//		EngagedTo:   nil,
//	}
//	w1 := &Woman{
//		Id:          "1",
//		Preferences: make(map[string]propcheck.Pair[int, *Man]),
//		EngagedTo:   nil,
//	}
//
//	w2 := &Woman{
//		Id:          "2",
//		Preferences: make(map[string]propcheck.Pair[int, *Man]),
//		EngagedTo:   nil,
//	}
//
//	m0 := &Man{
//		Id:                "0",
//		HaveNotProposedTo: nil,
//	}
//	m1 := &Man{
//		Id:                "1",
//		HaveNotProposedTo: nil,
//	}
//	m2 := &Man{
//		Id:                "2",
//		HaveNotProposedTo: nil,
//	}
//
//	var allMen *linked_list.LinkedList[*Man]
//	var manPreferences *linked_list.LinkedList[*Woman]
//	manPreferences = linked_list.Push(w2, manPreferences)
//	manPreferences = linked_list.Push(w1, manPreferences)
//	manPreferences = linked_list.Push(w0, manPreferences)
//
//	m0.HaveNotProposedTo = manPreferences
//	m1.HaveNotProposedTo = manPreferences
//	m2.HaveNotProposedTo = manPreferences
//	m0.Preferences = linked_list.ToArray(manPreferences)
//	m1.Preferences = linked_list.ToArray(manPreferences)
//	m2.Preferences = linked_list.ToArray(manPreferences)
//
//	allMen = linked_list.Push(m2, allMen)
//	allMen = linked_list.Push(m1, allMen)
//	allMen = linked_list.Push(m0, allMen)
//
//	w0.Preferences[m1.Id] = propcheck.Pair[int, *Man]{1, m1}
//	w0.Preferences[m0.Id] = propcheck.Pair[int, *Man]{0, m0}
//	w1.Preferences[m1.Id] = propcheck.Pair[int, *Man]{0, m1}
//	w1.Preferences[m0.Id] = propcheck.Pair[int, *Man]{1, m0}
//	//w2 is indifferent to all men
//	Match(allMen, womanPrefersMe)
//	if w0.EngagedTo.Id != m0.Id {
//		t.Errorf("Expected woman 0 to be engaged to man 0")
//	}
//	if w1.EngagedTo.Id != m1.Id {
//		t.Errorf("Expected woman 1 to be engaged to man 1")
//	}
//	if w2.EngagedTo.Id != m2.Id {
//		t.Errorf("Expected woman 2 to be engaged to man 2")
//	}
//	unstableMatchings := unstableMatchings([]*Woman{w0, w1, w2})
//	if len(unstableMatchings) > 0 {
//		t.Errorf("There were unstable matchings as follow:%v", unstableMatchings)
//	}
//}
