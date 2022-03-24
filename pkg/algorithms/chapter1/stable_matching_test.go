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

func shuffle[A any](toBeShuffled []*A) []*A {
	rr := make([]*A, len(toBeShuffled))
	copy(rr, toBeShuffled)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(rr), func(i, j int) {
		rr[i], rr[j] = rr[j], rr[i]
	})
	return rr
}

func TestStableMatching(t *testing.T) {
	var wPrefersMe = func(wp *Woman, me *Man) bool { //Does wp prefer m to whom she is currently engaged
		mEq := func(m1 *Man, m2 *Man) bool {
			if m1.Id == m2.Id {
				return true
			} else {
				return false
			}
		}

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
			r := Match(freeMen, wPrefersMe)
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
