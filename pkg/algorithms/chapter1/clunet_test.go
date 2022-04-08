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

	g0 := propcheck.ChooseInt(1, 3000)
	g1 := sets.ChooseSet(0, 50, g0, lt, eq)
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	prop := propcheck.ForAll(g1,
		"Validate Clunet switch algorithm  \n",
		func(xs []int) []*InputWire {
			var outputWires []*OutputWire
			for _, y := range xs {
				ow := OutputWire{
					Id:             y,
					InputJunctions: make([]*InputWire, len(xs), len(xs)), //[5]int
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
			r := MakeSwitches(ll)
			fmt.Printf("Scheduling an array of %v inputWires took %v\n", len(inputWires), time.Since(start))
			//for i, ow := range r { //Range loop over array of all inputWires
			//	currentOutputRoute := ow..OutputRoute
			//	for j := i + 1; j < len(inputWires); j++ { //For current ow iterate over all inputWires later in array and truncate ow's Proposed Schedule at earliest conflict
			//		otherOutputRoute := inputWires[j]//.OutputRoute
			//		for k, _ := range otherOutputRoute {
			//			if k < len(currentOutputRoute)-1 { //Current ow not at sea and is not truncated before otherShip
			//				if currentOutputRoute[k] == otherOutputRoute[k] { //Not at same port on same day
			//					errors = multierror.Append(errors, fmt.Errorf("Ship:%v scheduled port:%v conflicted with ow%v on day:%v", ow.Id, currentOutputRoute[k], inputWires[j].Id, k))
			//
			//				}
			//			}
			//		}
			//	}
			//}
			fmt.Println(r)
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{200, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
