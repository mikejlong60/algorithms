package skiena_1

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"testing"
	"time"
)

// Given an integer array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....
//
// You may assume the input array always has a valid answer.
//
// Input: nums = [1,5,1,1,6,4]
// Output: [1,6,1,5,1,4]
// Explanation: [1,4,1,5,1,6] is also accepted.
//
// Input: nums = [1,3,2,2,3,1]
// Output: [2,3,1,3,1,2]

// Not a pure function
func wiggleSort(xs []int) {
	//lastIdx := len(xs) - 1
	for i := 0; i < len(xs); i++ {
		//if xs[i] < xs[i+1] {
		//
		//		}
	}
	return
}

func TestWiggleSort(t *testing.T) {
	//g0 := propcheck.ArrayOfN(3, propcheck.Id(1))
	//g1 := propcheck.ArrayOfN(3, propcheck.Id(2))
	//g3 := propcheck.Map2(g0, g1, func(a, b []int) []int {
	//	return append(a, b...)
	//})
	g4 := propcheck.Id([]int{1, 2, 1, 2, 1, 2})
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}

	verify := func(xs []int) bool {
		left_lt_middle_middle_gt_right := func(l, m, r int) bool {
			if l < m && m > r {
				return true
			} else {
				return false
			}
		}
		left_gt_middle_middle_lt_right := func(l, m, r int) bool {
			if l > m && m < r {
				return true
			} else {
				return false
			}
		}
		var i = 0
		var r bool

		for { //Assume length of xs is divisible by three for simplicity
			if i+2 < len(xs) {
				if i == 0 || (i+1)%2 != 0 { //an odd element number
					r = left_lt_middle_middle_gt_right(xs[i], xs[i+1], xs[i+2])
					if !r {
						break
					}
				} else { // an even element number
					r = left_gt_middle_middle_lt_right(xs[i], xs[i+1], xs[i+2])
					if !r {
						break
					}
				}
			} else {
				break
			}
			i = i + 1
		}
		return r
	}
	prop := propcheck.ForAll(g4,
		"Verify wiggle sort  \n",
		func(xs []int) []int {
			wiggleSort(xs)
			return xs
		},
		func(xs []int) (bool, error) {
			var errors error

			if verify(xs) {
				fmt.Println("Correct!!!")
			} else {
				errors = fmt.Errorf("Actual:%v", xs)
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
