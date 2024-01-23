package skiena_1

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"math"
	"testing"
	"time"
)

type point struct {
	x float64
	y float64
}

// Finds the closest point (p1) to p and return that point and the original array with the closest point removed.
func findAndRemoveClosestPoint(p point, points []point) (point, []point) {
	distance := func(p1, p2 point) float64 {
		return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x)) + ((p2.y - p1.y) * (p2.y - p1.y))
	}
	var closestDistance = distance(p, points[0])
	var currentDistance float64
	var closestPointIdx int = 0
	for l, m := range points {
		currentDistance = distance(p, m)
		if currentDistance < closestDistance {
			closestDistance = currentDistance
			closestPointIdx = l
		}
	}
	closestPointToP := points[closestPointIdx]
	points = append(points[:closestPointIdx], points[closestPointIdx+1:]...)
	//returns closest point to p and original points array minus the closest point.
	return closestPointToP, points
}

// points is starting array, closestPoints is result array with last element being the next point from which you find the closest point c.
// starting array is missing the last element on every recursion.
func nearestNeighbor(points, closestPoints []point) ([]point, []point) {
	if len(points) == 0 {
		return points, closestPoints
	} else {
		p := points[len(points)-1] //Last element is the one we use to derive next closest point
		closestPointToP, points := findAndRemoveClosestPoint(p, points)
		closestPoints = append(closestPoints, closestPointToP)
		return nearestNeighbor(points, closestPoints)
	}
}

func TestNearestNeighbor(t *testing.T) {
	g0 := propcheck.ArrayOfN(10, propcheck.ChooseInt(-50, 51))
	g1 := propcheck.ArrayOfN(10, propcheck.ChooseInt(-50, 51))
	g2 := propcheck.Map2(g0, g1, func(xs, ys []int) []point {
		r := []point{}
		for i := range xs {
			p := point{float64(xs[i]), float64(ys[i])}
			r = append(r, p)
		}
		return r
	})
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g2,
		"Validate traveling salesman algo  \n",
		func(xs []point) []point {

			return xs
		},
		func(xs []point) (bool, error) {
			var errors error
			//expected := xs[0] * xs[1]
			//actual := xs[3]
			//if actual != expected {
			//	t.Errorf("Actual:%v Expected:%v", actual, expected)
			//}
			_, actual := nearestNeighbor(xs, []point{})
			fmt.Printf("Origin:%v\n", xs)
			fmt.Printf("Actual%v\n", actual)

			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]point](t, result)
}
