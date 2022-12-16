package chapter4

import (
	"github.com/greymatter-io/golangz/arrays"
	"sort"
	"testing"
)

func TestMergeSort1(t *testing.T) {

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

	xs := []int{4, 45, 12, 23, 35, 56, 78, 1}

	actual := MergeSort(xs, lt)

	expected := make([]int, len(xs))
	copy(expected, xs)
	sort.Ints(expected)

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestMergeSort2(t *testing.T) {

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

	xs := []int{4, 45, 12, 23, 35, 56, 78, 1, 23, 1, 5, 6, 7, 8, 9}

	actual := MergeSort(xs, lt)
	expected := make([]int, len(xs))
	copy(expected, xs)
	sort.Ints(expected)

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}
