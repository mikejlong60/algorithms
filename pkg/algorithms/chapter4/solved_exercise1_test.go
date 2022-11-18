package chapter4

import (
	"fmt"
	"testing"
	"time"
)

var gt = func(a, b int) bool {
	if a > b {
		return true
	} else {
		return false
	}
}

func TestSolvedExercise1(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 9, 3, 2}
	fmt.Printf("len of a:%v\n", len(a))
	b := PeakOfAList(a, 1, gt)
	if b != 9 {
		t.Errorf("Actual:%v Expected:%v", b, 9)
	}
}

func TestMax(t *testing.T) {
	a := []int{0, 1, 20, 200, 3, 4, 5, 6, 7, 201, 3, 2}
	fmt.Printf("len of a:%v\n", len(a))
	b := PeakOfAList(a, 1, gt)
	if b != 201 {
		t.Errorf("Actual:%v Expected:%v", b, 201)
	}
}

func TestSolvedExercise1bug(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 3}
	fmt.Printf("len of a:%v\n", len(a))
	b := PeakOfAList(a, 1, gt)
	if b != 9 {
		t.Errorf("Actual:%v Expected:%v", b, 9)
	}
}

func TestSolvedExercise1a(t *testing.T) {
	a := []int{0, 1, 2}
	fmt.Printf("len of a:%v\n", len(a))
	b := PeakOfAList(a, 1, gt)
	if b != 2 {
		t.Errorf("Actual:%v Expected:%v", b, 2)
	}
}

func TestSolvedExercise1b(t *testing.T) {
	aaaa := []int{-30, -31, -32, -33, -30, -31, 200020, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, 3000010, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -200, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -30, -31, -32, -33, -28, -29, -27, -26, -25, -24, -23, -22, -21, -20, -19, -18, -17, -16, -15, -14, -13, -12, -11, -10, -9, -8, -7, -6, -5, -4. - 3. - 2. - 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1770, 1771, 191, 190, 18, 17, 16, 15, 14, 13, 12, 11, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4}
	aaa := append(append(aaaa, aaaa...), append(aaaa, aaaa...)...)
	aa := append(aaa, aaa...)
	a := append(aa, aa...)
	fmt.Printf("len of a:%v\n", len(a))
	start := time.Now()
	b := PeakOfAList(a, 10, gt)
	fmt.Printf("Took:%v", time.Since(start))
	if b != 3000010 {
		t.Errorf("Actual:%v Expected:%v", b, 3000010)
	}
}
