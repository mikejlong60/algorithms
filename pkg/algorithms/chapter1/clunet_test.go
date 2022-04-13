package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/linked_list"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"math/rand"
	"testing"
	"time"
)

func shuffleRoutes(toBeShuffled []*OutputWire) []*OutputWire {
	rr := make([]*OutputWire, len(toBeShuffled))
	copy(rr, toBeShuffled)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(rr), func(i, j int) {
		rr[i], rr[j] = rr[j], rr[i]
	})
	return rr
}

func TestClunetSwitch(t *testing.T) {
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

	g0 := propcheck.ChooseInt(1, 30000)
	g1 := sets.ChooseSet(0, 6000, g0, lt, eq)
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	prop := propcheck.ForAll(g1,
		"Validate Clunet switch algorithm  \n",
		func(xs []int) []*InputWire {
			var outputWires []*OutputWire
			for _, y := range xs {
				ow := OutputWire{
					Id:             y,
					InputJunctions: make([]*InputWire, len(xs), len(xs)),
				}
				outputWires = append(outputWires, &ow)
			}
			var r []*InputWire
			for _, x := range xs {
				s := InputWire{
					Id:                    x,
					OutputWirePreferences: shuffleRoutes(outputWires),
				}
				r = append(r, &s)
			}
			return r
		},
		func(inputWires []*InputWire) (bool, error) {
			var errors error
			ll := linked_list.ToList(inputWires)

			start := time.Now()
			MakeSwitches(ll)
			fmt.Printf("Scheduling an array of %v inputWires took %v\n", len(inputWires), time.Since(start))
			//TODO validate that last junction of every output wire is unique. To do that take the last element of every outout wire and use set to compare its set length to the length of the last element array.
			//TODO And then verify that you have the proper number of output wires == to the size of the XS set.
			//TODO And then make the ChooseInt smaller after you show Jeff.
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
