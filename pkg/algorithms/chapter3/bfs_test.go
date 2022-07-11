package chapter3

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"testing"
	"time"
)

//..
type BFSNode struct {
	Id          int
	Connections []int //Ids of the nodes to which this Node has connections
}

func (w BFSNode) String() string {
	return fmt.Sprintf("BNFSNode{Id:%v, Connections:%v}", w.Id, w.Connections)
}

// Generates a graph of Node structures with a size in the indicated range
func GraphGen(min, maxExc int) func(propcheck.SimpleRNG) ([]BFSNode, propcheck.SimpleRNG) {
	return func(rng propcheck.SimpleRNG) ([]BFSNode, propcheck.SimpleRNG) {
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

		ge := propcheck.Int()
		ge2 := sets.ChooseSet(min, maxExc, ge, lt, eq)
		var i, rng2 = ge2(rng)
		var xs []int
		var xss [][]int
		//Generate a random list of indices into array i
		listIdxsGen := propcheck.ChooseArray(0, len(i), propcheck.ChooseInt(0, len(i)-1))
		for x := 0; x < len(i); x++ {
			xs, rng2 = listIdxsGen(rng2)
			xss = append(xss, xs)
		}
		var nodes []BFSNode
		//Create an array of nodes with actual connections by Id to other nodes
		for a, _ := range xss {
			var connections []int
			for ii := 0; ii < len(xss[a]); ii++ {
				connections = append(connections, i[xss[a][ii]])
			}
			node := BFSNode{
				Id:          i[a],
				Connections: connections,
			}
			nodes = append(nodes, node)
		}
		return nodes, rng2
	}
}

func TestBFS(t *testing.T) {
	g := GraphGen(0, 10)
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	graph, rng := g(rng)
	fmt.Println(graph)
	//fmt.Println(len(graph))

}
