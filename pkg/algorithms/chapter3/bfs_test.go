package chapter3

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"testing"
	"time"
)

type BFSNode struct {
	Id          int
	Connections []int //Ids of the nodes to which this Node has connections
}

func (w BFSNode) String() string {
	return fmt.Sprintf("BNFSNode{Id:%v, Connections:%v}", w.Id, w.Connections)
}

// Generates a graph of Node structures with A size in the indicated range using the given Gen
func GraphGen() func(propcheck.SimpleRNG) ([]BFSNode, propcheck.SimpleRNG) {
	return func(rng propcheck.SimpleRNG) ([]BFSNode, propcheck.SimpleRNG) {

		var r []BFSNode
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

		maxListSize := 10
		minListSize := 0
		ge := propcheck.Int()
		ge2 := sets.ChooseSet(minListSize, maxListSize, ge, lt, eq)
		////w := propcheck.(start, stopExclusive, l) //Generates a list of ACMs with a stopExclusive in the range startInclusive - StopInclusive
		i, rng2 := ge2(rng)
		gen2 := propcheck.ChooseArray(0, len(i)-1, ge2) //Generator to produce an index to a random element that you later use to choose from the list you just generated.
		n, rng3 := gen2(rng2)                           //Index into the array of ACMs

		rr := propcheck.Id(12)
		g := func(x int) func(propcheck.SimpleRNG) (int, propcheck.SimpleRNG) {
			return propcheck.Id(x + 1)
		}
		res := propcheck.FlatMap(rr, g)
		res(rng)

		//gen2 := propcheck.ArrayOfN(n, ge2) //Generator to produce an index to a random element that you later use to choose from the list you just generated.
		//n, rng3 := gen2(rng3)                    //Index into the array of ACMs
		fmt.Println(n)
		gdd := func(nodes []int) func(propcheck.SimpleRNG) ([]int, propcheck.SimpleRNG) {
			//var r []BFSNode
			//	for _, k := range nodes {
			//		r = append(r, BFSNode{
			//			Id:          k,
			//			Connections: nil,
			//		})
			//
			//	return r, rng
			//}
			//sets.SetEquality()
			ass := sets.ChooseSet(0, len(nodes), propcheck.Id(nodes), lt, eq)
			//return propcheck.ChooseArray(0, len(nodes), propcheck.Id(3))
		}
		//dfd := propcheck.Map2(ge2, )

		piss := propcheck.FlatMap(ge2, gdd)
		//fmt.Println(piss)
		//	gg := func(nodes []int) func(propcheck.SimpleRNG) ([]BFSNode, propcheck.SimpleRNG) {
		//		return
		//		for _, k := range nodes {
		//			r = append(r, BFSNode{
		//				Id:          k,
		//				Connections: nil,
		//			})
		//		}
		//	}
		//}
		//	fff := propcheck.FlatMap(ge2, gg)
		//	for _, k := range i {
		//		r = append(r, BFSNode{
		//			Id:          k,
		//			Connections: nil,
		//		})
		//	}
		crap, df := piss(rng2)
		return crap, df
	}
}

func TestBFS(t *testing.T) {
	g := GraphGen()
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	graph, rng := g(rng)
	fmt.Println(graph)
	fmt.Println(len(graph))

}
