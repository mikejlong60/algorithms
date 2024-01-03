package chapter5

import (
	"github.com/greymatter-io/golangz/propcheck"
	"testing"
	"time"
)

func TestHorner(t *testing.T) {

	var n int
	horner := func(a, x int) int {
		p := a
		for i := n; i >= 0; i-- {
			p = p*x + i
		}
		return p
	}

	g0 := propcheck.ArrayOfN(2, propcheck.ChooseInt(0, 30))
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g0,
		"Validate polynomial evaluation  \n",
		func(xs []int) []int {
			n = xs[1]
			r := horner(xs[0], xs[1])
			return append(xs, r)
		},
		func(xs []int) (bool, error) {
			var errors error
			expected := xs[0] * xs[1]
			actual := xs[3]
			if actual != expected {
				t.Errorf("Actual:%v Expected:%v", actual, expected)
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
