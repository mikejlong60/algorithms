package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/linked_list"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"github.com/hashicorp/go-multierror"
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

	g0 := propcheck.ChooseInt(1, 300)
	g1 := sets.ChooseSet(0, 60, g0, lt, eq)
	rng := propcheck.SimpleRNG{Seed: 42934000} //propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	//	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
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
					OutputWirePreferences: outputWires, //shuffleRoutes(outputWires),//TODO Shuffling makes the bug show up less often. But its still there.
				}
				r = append(r, &s)
			}
			return r
		},
		func(inputWires []*InputWire) (bool, error) {

			eq := func(l, r *InputWire) bool {
				if l.Id == r.Id {
					return true
				} else {
					return false
				}
			}

			lt := func(l, r *InputWire) bool {
				if l.Id < r.Id {
					return true
				} else {
					return false
				}
			}
			var errors error
			ll := linked_list.ToList(inputWires)

			start := time.Now()
			r := MakeSwitches(ll)
			fmt.Printf("Scheduling an array of %v inputWires took %v\n", len(inputWires), time.Since(start))
			var liw []*InputWire
			l := len(r)
			for _, ow := range r {
				liw = append(liw, ow.InputJunctions[l-1])
			}
			liwAsSet := sets.ToSet(liw, lt, eq)
			if len(liwAsSet) != len(r) {
				errors = multierror.Append(errors, fmt.Errorf("Expected the length:%v of the set of last input junctions for all output wires to equal the size of the set resulting from the switching operation:%v", len(liwAsSet), len(r)))
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	fmt.Println(rng)
	result := prop.Run(propcheck.RunParms{1, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
