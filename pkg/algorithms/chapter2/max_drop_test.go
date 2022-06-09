package chapter2

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func findBreakingPointWithoutBreakingJar(ladder []int, breakingPoint int) int {
	breakingPointIsHigher := ladder[0] <= breakingPoint
	var wrungB4BreakingPoint = -1

	if breakingPointIsHigher {
		for i := 0; i < len(ladder); i++ {
			if ladder[i] >= breakingPoint {
				wrungB4BreakingPoint = ladder[i-1] //the wrung of the ladder right before the breaking point
				break
			}
		}
	} else { //breaking point is lower
		for i := len(ladder) - 1; i >= 0; i-- {
			if ladder[i] >= breakingPoint {
				wrungB4BreakingPoint = ladder[i-1] //the wrung of the ladder right before the breaking point
				break
			}
		}
	}
	return wrungB4BreakingPoint
}

func highestBreakingPoint(ladder []int, breakingPoint, budget, usedBudget int) int {
	if usedBudget+1 == budget {
		return findBreakingPointWithoutBreakingJar(ladder, breakingPoint)
	} else { //Divide in half again
		lowerHalf := ladder[0 : len(ladder)/2]
		upperHalf := ladder[len(ladder)/2:]
		if breakingPoint <= upperHalf[0] {
			return highestBreakingPoint(lowerHalf, breakingPoint, budget, usedBudget+1)
		} else {
			return highestBreakingPoint(upperHalf, breakingPoint, budget, usedBudget+1)
		}
	}
}

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
			r := highestBreakingPoint(xs.A, xs.B, budget, 0)
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
