package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/linked_list"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

//This test differs from TestStableMatching in that the woman may be indifferent to me, and in that case will not change husbands.
func TestIndifferentStableMatching2(t *testing.T) {
	var allTheMen []*Man
	var allTheWomen []*Woman
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	g0 := propcheck.ChooseInt(0, 3000)
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
				womenForMan := shuffle(allTheWomen)
				var allWomen *linked_list.LinkedList[*Woman]
				for _, s := range womenForMan {
					allWomen = linked_list.Push(s, allWomen)
				}
				allTheMen[i].Preferences = allWomen
			}

			for _, s := range allTheWomen {
				s.Preferences = shuffle(allTheMen)
			}

			return freeMen
		}
		ra := propcheck.ListOfN(a, propcheck.String(100))
		rb := propcheck.ListOfN(a, propcheck.String(100))
		return propcheck.Map2(ra, rb, mw)
	}

	g := propcheck.FlatMap(g0, fa)

	prop := propcheck.ForAll(g,
		"Make a bunch of men and women and match them up  \n",
		func(freeMen *linked_list.LinkedList[*Man]) []*Woman {
			len := linked_list.Len(freeMen)
			start := time.Now()
			r := Match(freeMen, indifferentWPrefersMe)
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
				fmt.Printf("%v men never got married\n", len(arrays.SetMinus(allMenIds, allHusbandIds, mEq)))
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

func TestIndifferentStableMatching1(t *testing.T) {
	w1 := &Woman{
		Id:          "1",
		Preferences: nil,
		EngagedTo:   nil,
	}
	w2 := &Woman{
		Id:          "2",
		Preferences: nil,
		EngagedTo:   nil,
	}

	m1 := &Man{
		Id:          "1",
		Preferences: nil,
	}
	m2 := &Man{
		Id:          "2",
		Preferences: nil,
	}

	w1.Preferences = []*Man{m2, m1}
	w2.Preferences = []*Man{m2, m1}

	var allMen *linked_list.LinkedList[*Man]
	var manPreferences *linked_list.LinkedList[*Woman]
	manPreferences = linked_list.Push(w2, manPreferences)
	manPreferences = linked_list.Push(w1, manPreferences)

	m1.Preferences = manPreferences
	m2.Preferences = manPreferences //TODO Fix this to be different
	allMen = linked_list.Push(m2, allMen)
	allMen = linked_list.Push(m1, allMen)
	actual := Match(allMen, indifferentWPrefersMe)
	if w1.EngagedTo.Id != m1.Id {
		t.Errorf("Expected woman 1 to be engaged to man 1")
	}
	if w2.EngagedTo.Id != m2.Id {
		t.Errorf("Expected woman 2 to be engaged to man 2")
	}
	fmt.Println(actual)
}

var indifferentWPrefersMe = func(wp *Woman, courtier *Man) bool { //Does woman prefer this man to the one to which she is currently assigned?
	mEq := func(m1 *Man, m2 *Man) bool {
		if m1.Id == m2.Id {
			return true
		} else {
			return false
		}
	}
	var courtierRanking = -1
	var currentFianceeRanking = -1
	for i, m := range wp.Preferences {
		if mEq(m, courtier) {
			courtierRanking = i
		} else if mEq(wp.EngagedTo, m) {
			currentFianceeRanking = i
		}
	}
	if courtierRanking == -1 && currentFianceeRanking == -1 { //Woman is indifferent to both men. So she will stick with her current fiancee.
		return false
	} else if courtierRanking >= 0 && currentFianceeRanking == -1 { //Woman is indifferent to her current husband but not the courtier. So she chooses the courtier.
		return true
	} else if courtierRanking == -1 && currentFianceeRanking >= 0 { //Woman is indifferent to the courter but she had a preference for her current fiancee.
		//So she sticks with her current fiance
		return false
	} else if courtierRanking < currentFianceeRanking {
		//fmt.Printf("woman:%v prefers this man:%v over current one:%v\n", wp.Id, counter.Id, wp.EngagedTo.Id)
		return true
	} else {
		//fmt.Printf("woman:%v prefers current man:%v over this one:%v\n", wp.Id, wp.EngagedTo.Id, counter.Id)
		return false
	}
}

func TestIndifferentStableMatching3(t *testing.T) {

	w1 := &Woman{
		Id:          "1",
		Preferences: nil,
		EngagedTo:   nil,
	}
	w2 := &Woman{
		Id:          "2",
		Preferences: nil,
		EngagedTo:   nil,
	}

	m1 := &Man{
		Id:          "1",
		Preferences: nil,
	}
	m2 := &Man{
		Id:          "2",
		Preferences: nil,
	}

	var allMen *linked_list.LinkedList[*Man]
	var manPreferences *linked_list.LinkedList[*Woman]
	manPreferences = linked_list.Push(w2, manPreferences)
	manPreferences = linked_list.Push(w1, manPreferences)

	m1.Preferences = manPreferences
	m2.Preferences = manPreferences //TODO Fix this to be different
	allMen = linked_list.Push(m2, allMen)
	allMen = linked_list.Push(m1, allMen)

	w1.Preferences = []*Man{m2, m1}
	w2.Preferences = []*Man{m2, m1}
	actual := Match(allMen, indifferentWPrefersMe)
	if w1.EngagedTo.Id != m2.Id {
		t.Errorf("Expected woman 1 to be engaged to man 2")
	}
	if w2.EngagedTo.Id != m1.Id {
		t.Errorf("Expected woman 2 to be engaged to man 1")
	}

	fmt.Println(actual)
}
