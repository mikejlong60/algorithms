package chapter2

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func TestMaxDropGivenBudget(t *testing.T) { //This reverses the order because foldRight does that. There is a Reverse function in Golangz if that matters to you.
	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}

	eq := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}

	g0 := propcheck.ChooseInt(0, 100)
	g1 := sets.ChooseSet(6, 20, g0, lt, eq) //This array comes back sorted

	f := func(xss []int) func(propcheck.SimpleRNG) (propcheck.Pair[[]int, int], propcheck.SimpleRNG) {
		g := propcheck.ChooseInt(4, len(xss)-1)
		i := func(x int) propcheck.Pair[[]int, int] {
			return propcheck.Pair[[]int, int]{
				xss,
				xss[x],
			}
		}
		h := propcheck.Map(g, i)
		return h
	}

	g2 := propcheck.FlatMap(g1, f)
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	fmt.Println(rng)
	budget := 2 //TODO make this a random value between 2 and the size of the generated array.
	prop := propcheck.ForAll(g2,
		"Exercise 2.8b, the max jar drop given a budget of not-to-exceed broken jars.",
		func(xs propcheck.Pair[[]int, int]) propcheck.Pair[int, propcheck.Pair[[]int, int]] {
			//A is the ladder, B is the breaking point on the ladder(the actual value in the array, not its index).
			r := HighestBreakingPoint(xs.A, xs.B, budget, 0)
			return propcheck.Pair[int, propcheck.Pair[[]int, int]]{r, xs}
		},
		func(highestWrungAEtAll propcheck.Pair[int, propcheck.Pair[[]int, int]]) (bool, error) {
			var errors error
			breakingPoint := highestWrungAEtAll.B.B
			highestWrung := highestWrungAEtAll.A
			if highestWrung >= breakingPoint {
				errors = multierror.Append(errors, fmt.Errorf("Expected highest non-breaking wrung withing budget to be:%v but was:%v", highestWrung, breakingPoint))
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[propcheck.Pair[[]int, int]](t, result)
}
