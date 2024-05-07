package skiena_3

import (
	"fmt"
	"github.com/greymatter-io/golangz/option"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func TestBalancedBinaryTreeFound(t *testing.T) {

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
	g0 := propcheck.ChooseArray(0, 62, propcheck.ChooseInt(-1000000, 1000000))

	fg := func(fx []int) func(rng propcheck.SimpleRNG) (propcheck.Pair[int, []int], propcheck.SimpleRNG) {
		a := propcheck.ChooseInt(0, len(fx))
		g := func(x int) propcheck.Pair[int, []int] {
			return propcheck.Pair[int, []int]{x, fx}
		}
		r := propcheck.Map(a, g)
		return r
	}

	g1 := propcheck.FlatMap(g0, fg)
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g1,
		"Make a binary tree",
		func(a propcheck.Pair[int, []int]) propcheck.Pair[int, []int] {
			return a
		},
		func(a propcheck.Pair[int, []int]) (bool, error) {

			btree := BinaryTree(a.B, lt)
			var errors error
			if len(a.B) > 0 {
				idx := a.A
				f := func(x int) int {
					if x != a.B[idx] {
						errors = multierror.Append(errors, fmt.Errorf("Not Found"))
					}
					return x
				}
				option.Map(Find(btree, a.B[idx], lt, eq), f)
			}

			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[propcheck.Pair[int, []int]](t, result)
}

func TestBalancedBinaryTreeNotFound(t *testing.T) {

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
	g0 := propcheck.ChooseArray(0, 62, propcheck.ChooseInt(0, 100))

	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Make a binary tree",
		func(a []int) []int {
			return a
		},
		func(a []int) (bool, error) {
			btree := BinaryTree(a, lt)
			var errors error
			if len(a) > 0 {
				f := func(x int) int {
					errors = fmt.Errorf("Should not have Found:%v", x)
					return x
				}
				option.Map(Find(btree, -10, lt, eq), f) //-10 is not in tree
			}

			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
