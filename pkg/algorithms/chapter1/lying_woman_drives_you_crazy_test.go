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

func TestStableMatchingTODOWithLyingWoman(t *testing.T) {
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
				womenForMan := shuffleAny(allTheWomen)
				var allWomen *linked_list.LinkedList[*Woman]
				for _, s := range womenForMan {
					allWomen = linked_list.Push(s, allWomen)
				}
				allTheMen[i].Preferences = allWomen
			}

			for _, s := range allTheWomen {
				wpref := shuffleAny(allTheMen)
				var wprefMap = make(map[string]int, len(wpref))
				for i, m := range wpref {
					wprefMap[m.Id] = i
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
		"Make a bunch of men and women and match them up  \n",
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
//   Make a new list r of strings of UnstableMatchings
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
