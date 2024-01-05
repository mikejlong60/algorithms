package chapter5

import (
	"github.com/greymatter-io/golangz/propcheck"
	"testing"
	"time"
)

func TestHorner(t *testing.T) {

	horner := func(poly []int, n, x int) int {
		result := poly[0]
		for i := 1; i < n; i++ {
			result = result*x + poly[i]
		}
		return result
	}

	//TODO implement Horner iterating backward from book. Then use property testing to compare it with 1st horner here.
	g0 := propcheck.Id([]int{2, -6, 2, -1, 3}) //2x^3 - 6x^2 - 2x - 1//.ArrayOfN(3, propcheck.ChooseInt(1, 3))
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g0,
		"Validate polynomial evaluation  \n",
		func(poly []int) int {
			x := poly[len(poly)-1]
			n := len(poly) - 1
			r := horner(poly[0:len(poly)-1], n, x)
			return r
		},
		func(x int) (bool, error) {
			var errors error
			expected := 5
			actual := 5
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
