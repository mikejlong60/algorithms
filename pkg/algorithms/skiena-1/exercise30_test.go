package chapter5

import (
	"github.com/greymatter-io/golangz/propcheck"
	"math"
	"testing"
	"time"
)

type point struct {
	x int
	y int
}

// Finds the closest point (p1) to p and the points that remain after removing p1. p is not in the passed array of points
func findAndRemoveClosestPoint(p point, points []point) (point, []point) {
	var closestDistance = math.Abs(float64((p.x + p.y) - (points[0].x + points[0].y)))
	var currentDistance float64
	var closestPointIdx int = 0
	for l, m := range points {
		currentDistance = math.Abs(float64((p.x + p.y) - (m.x + m.y)))
		if currentDistance < closestDistance {
			closestDistance = currentDistance
			closestPointIdx = l
		}
	}

	//returns original points array minus the closest point.
	return points[closestPointIdx], append(points[0:closestPointIdx], points[closestPointIdx+1])
}

// points is starting array, closestPoints is result array with last element being the next point from which you find the closest point c.
// starting array is missing the last element on every recursion.
func nearestNeighbor(points, closestPoints []point) ([]point, []point) {
	if len(points) == 0 {
		return points, closestPoints
	} else {
		p := points[0]
		closestPointToP, yyy := findAndRemoveClosestPoint(p, points)
		closestPoints = append(closestPoints, closestPointToP)
		points = yyy
		return nearestNeighbor(points, closestPoints)
	}
}

// TODO Fix this to work with above
func TestNearestNeighbor(t *testing.T) {
	g0 := propcheck.ArrayOfN(3, propcheck.ChooseInt(3, 300000))
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g0,
		"Validate weird multiplication algorithm  \n",
		func(xs []int) []int {
			c = xs[2]
			r := mult(xs[0], xs[1])
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
