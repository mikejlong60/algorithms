package skiena_3

import (
	"bitbucket.org/pcastools/hash"
	"fmt"
	"github.com/greymatter-io/golangz/option"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func TestGetAndSetForProbablyNoHashCollisions(t *testing.T) {
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	f := func(a []int, b []string) []KeyValuePair[int32, string] {
		var r = make([]KeyValuePair[int32, string], len(a))
		for c, d := range a {
			r[c] = KeyValuePair[int32, string]{int32(d), b[c]}
		}
		return r
	}

	ge := propcheck.ChooseArray(500, 500, propcheck.ChooseInt(-100000, 100000))
	gf := propcheck.ChooseArray(500, 500, propcheck.String(40))
	gg := propcheck.Map2(ge, gf, f)

	prop := propcheck.ForAll(gg,
		"Test get and set when you have few hash collisions  \n",
		func(xs []KeyValuePair[int32, string]) []KeyValuePair[int32, string] {
			return xs
		},
		func(xss []KeyValuePair[int32, string]) (bool, error) {
			fFNV32a := func(x int32) uint32 {
				return hash.Int32(x)
			}
			eq := func(k1, k2 KeyValuePair[int32, string]) bool {
				if k1.key == k2.key {
					return true
				} else {
					return false
				}
			}
			m := New2[int32, string](eq, fFNV32a, 51)
			start := time.Now()
			for _, b := range xss {
				p := func(s KeyValuePair[int32, string]) bool {
					if s.key == b.key { //Close around the variable in the loop
						return true
					} else {
						return false
					}
				}
				m = Set2(m, b, p)
			}
			fmt.Printf("Inserting %v values into hashmap took:%v\n", len(xss), time.Since(start))

			var errors error

			//The value should be in the map
			start = time.Now()
			for _, b := range xss {
				p := func(s KeyValuePair[int32, string]) bool {
					if s.key == b.key { //Close around the variable in the loop
						return true
					} else {
						return false
					}
				}
				err := option.GetOrElse(Get2(m, b.key, p), KeyValuePair[int32, string]{b.key, fmt.Sprintf("Should have found:%v in HashMap", 188)}) //fmt.Sprintf("Should not have found:%v in HashMap", 188))
				if err.key != b.key {
					errors = multierror.Append(errors, fmt.Errorf("Should have found:%v in HashMap", b.key))
				}
			}
			fmt.Printf("Getting %v values from hashmap took:%v\n", len(xss), time.Since(start))
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]KeyValuePair[int32, string]](t, result)
}

func TestGetAndSetForManyHashCollisions(t *testing.T) {
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	f := func(a []int, b []string) []KeyValuePair[int32, string] {
		var r = make([]KeyValuePair[int32, string], len(a))
		for c, d := range a {
			r[c] = KeyValuePair[int32, string]{int32(d), b[c]}
		}
		return r
	}

	ge := propcheck.ChooseArray(10, 10, propcheck.ChooseInt(0, 3))
	gf := propcheck.ChooseArray(10, 10, propcheck.String(40))
	gg := propcheck.Map2(ge, gf, f)

	prop := propcheck.ForAll(gg,
		"Test get and set when you have many hash collisions  \n",
		func(xs []KeyValuePair[int32, string]) []KeyValuePair[int32, string] {
			return xs
		},
		func(xss []KeyValuePair[int32, string]) (bool, error) {
			fFNV32a := func(x int32) uint32 {
				return hash.Int32(x)
			}
			eq := func(k1, k2 KeyValuePair[int32, string]) bool {
				if k1.key == k2.key {
					return true
				} else {
					return false
				}
			}
			m := New2[int32, string](eq, fFNV32a, 3)
			start := time.Now()
			for _, b := range xss {
				p := func(s KeyValuePair[int32, string]) bool {
					if s.key == b.key { //Close around the variable in the loop
						return true
					} else {
						return false
					}
				}
				m = Set2(m, b, p)
			}
			fmt.Printf("Inserting %v values into hashmap took:%v\n", len(xss), time.Since(start))

			var errors error

			//The value should be in the map
			start = time.Now()
			for _, b := range xss {
				p := func(s KeyValuePair[int32, string]) bool {
					if s.key == b.key { //Close around the variable in the loop
						return true
					} else {
						return false
					}
				}
				err := option.GetOrElse(Get2(m, b.key, p), KeyValuePair[int32, string]{b.key, fmt.Sprintf("Should have found:%v in HashMap", 188)}) //fmt.Sprintf("Should not have found:%v in HashMap", 188))
				if err.key != b.key {
					errors = multierror.Append(errors, fmt.Errorf("Should have found:%v in HashMap", b.key))
				}
			}
			fmt.Printf("Getting %v values from hashmap took:%v\n", len(xss), time.Since(start))
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]KeyValuePair[int32, string]](t, result)
}
