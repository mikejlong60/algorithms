package chapter2

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func TestSumMatrix(t *testing.T) {

	g0 := propcheck.ChooseInt(1, 3000)
	g1 := propcheck.ChooseArray(0, 1000, g0)
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g1,
		"Validate exercise 2.6, a sort-of sum of a matrix  \n",
		func(xs []int) [][]int64 {
			return sum(xs)
		},
		func(xss [][]int64) (bool, error) {
			var errors error
			for i := 1; i < len(xss); i++ {
				last := xss[i-1][1]
				if xss[i][1] < last {
					errors = multierror.Append(errors, fmt.Errorf("Array element sum[%v] should not have been less than previous accumulated value", xss[i][1]))
				}
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
