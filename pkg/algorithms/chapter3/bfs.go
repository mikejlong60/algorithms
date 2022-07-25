package chapter3

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"time"
)

type Node struct {
	Id          int
	Connections []*Node
}

type BFS = [][]Edge

type Edge struct {
	u int //the Id of the beginning node of the edge
	v int //the Id of the ending node of the edge
}

type NodeLayerTuple struct {
	Id    int //The node Id
	layer int // Zero indexed array index indicating the layer(array index) in the BFS array the node lives
}

//Breadth-First search with cycle detection
//Params:
//  graph a hashmap of all the nodes in te graph. Facilitates n log n lookup
//  rootId the Node Id of the root node, the one at the top of the mobile from which all the other nodes hang
//Returns:
//   BFS  - the search tree represented as an array of layers, each layer constisting of an array of Edges(u, v)
//   bool - whether or not the resulting search tree contained a cycle. A cycle is a relationship between two nodes that is farther than one layer apart.
//   int - the number of nodes in the BFS
func BFSearch(graph map[int]*Node, rootId int) (BFS, bool, int) {
	hasCycle := func(nodeId int, currentLayer int, layers map[int]NodeLayerTuple) bool {
		l := layers[nodeId]
		if currentLayer-2 >= l.layer { //there is a cycle
			return true
		} else {
			return false
		}
	}

	var bfs = [][]Edge{}
	l0 := []Edge{{u: -1, v: rootId}}

	//A lookup map so you can look up whether or not a Node has been seen and if so what layer it is in.
	layersLookup := make(map[int]NodeLayerTuple, len(graph))
	layersLookup[rootId] = NodeLayerTuple{
		Id:    rootId,
		layer: 0,
	}
	bfs = append(bfs, l0)

	var graphHasACycle = false
	var i = 0 //current layer you are finding edges for
	for {
		var pendingLayer []Edge
		for _, k := range bfs[i] {
			node, _ := graph[k.v]
			for _, m := range node.Connections {
				//Lookup tail(v) of every edge in the layer to see if it has been seen before. If not add it to pending layer.
				_, alreadySeen := layersLookup[m.Id]
				if !alreadySeen {
					pendingLayer = append(pendingLayer, Edge{u: k.v, v: m.Id})
					layersLookup[m.Id] = NodeLayerTuple{Id: m.Id, layer: i + 1}
				} else { //Don't add it since we already know about this Node. But DO see if its a cycle.
					if !graphHasACycle { //Can only set this value to true one time
						graphHasACycle = hasCycle(m.Id, i, layersLookup)
					}
				}
			}
		}
		if len(pendingLayer) > 0 {
			bfs = append(bfs, pendingLayer)
			i = i + 1
		} else {
			break
		}
	}
	return bfs, graphHasACycle, len(layersLookup)
}

func BFSEquality(a, b []Edge) bool {
	edgeEq := func(a, b Edge) bool {
		if a.u == b.u && a.v == b.v {
			return true
		} else {
			return false
		}
	}
	if sets.SetEquality(a, b, edgeEq) {
		return true
	} else {
		return false
	}
}

func Rule3_2(graph map[int]*Node, rootNode int) (bool, bool, string) {
	start := time.Now()
	bfsTree, hasCycle, numNodes := BFSearch(graph, rootNode)
	fmt.Printf("BFS Search on a tree of %v nodes took: %v\n", len(graph), time.Since(start))
	numEdgesInTree := func(tree BFS) int {
		var edges int
		for _, node := range tree {
			edges = edges + len(node)
		}
		return edges - 1
	}

	numEdges := numEdgesInTree(bfsTree)
	isConnected := true
	hasN_1Edges := numNodes-1 == numEdges

	//hasCycle is based upon the original graph. The resulting bfsTree has no cycles
	return !hasCycle, isConnected && hasN_1Edges, fmt.Sprintf("Has Cycle:%v, isConnected: %v, has n-1 edges:%v\n:", hasCycle, isConnected, hasN_1Edges)
}

func GraphGenReal(lower, upperExc int) func(propcheck.SimpleRNG) (propcheck.Pair[map[int]*Node, int], propcheck.SimpleRNG) {
	return func(rng propcheck.SimpleRNG) (propcheck.Pair[map[int]*Node, int], propcheck.SimpleRNG) {
		eq := func(l, r int) bool {
			if l == r {
				return true
			} else {
				return false
			}
		}

		lt := func(l, r int) bool {
			if l < r {
				return true
			} else {
				return false
			}
		}

		nodeIds, rng2 := sets.ChooseSet(lower, upperExc, propcheck.ChooseInt(0, 1000000), lt, eq)(rng)

		graph := make(map[int]*Node, len(nodeIds))
		for _, j := range nodeIds {
			graph[j] = &Node{Id: j}
		}

		var rng3 = rng2
		var connectionIds []int
		for _, k := range graph {
			var connections []*Node
			connectionIds, rng3 = sets.ChooseSet(0, len(graph), propcheck.ChooseInt(0, len(graph)), lt, eq)(rng3)
			for _, l := range connectionIds {
				connections = append(connections, graph[nodeIds[l]])
			}
			k.Connections = connections
		}
		root, rng4 := propcheck.ChooseInt(0, len(graph))(rng3)
		return propcheck.Pair[map[int]*Node, int]{graph, nodeIds[root]}, rng4
	}
}
