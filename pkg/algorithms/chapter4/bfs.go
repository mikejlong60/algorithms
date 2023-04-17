package chapter4

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
)

// TODO add mechanism for including the total distance from starting node to the current node to deternine the shortest
// path to it if you have seen it before. And if you have seen it before and it's DistanceFromStartingNode is less than the
// DistanceFromStartingNode you recorded previously then move that node to the current layer in the tree by changing the Layer.Id
// value in the NodeLayerTuple.
type Node struct {
	Id                       int
	Connections              []*Node
	DistanceFromStartingNode int //This stores the current distance from the starting node.
}

type Edge struct {
	u      int //the Id of the beginning node of the edge
	v      int //the Id of the ending node of the edge
	length int //the length of the edge
}

type NodeLayerTuple struct {
	Id    int //The node Id
	layer int // Zero indexed array index indicating the layer(array index) in the Tree array the node lives
}

// Breadth-First search with cycle detection of a directed graph
// Params:
//
//	graph a hashmap of all the nodes in the graph. Facilitates n log n lookup
//	rootId the Node Id of the root node, the one at the top of the mobile from which all the other nodes hang
//
// Returns:
//
//	Tree  - the search tree represented as an array of layers, each layer constisting of an array of Edges(u, v)
//	bool - whether or not the resulting search tree contained a cycle. A cycle is a relationship between two nodes that is farther than one layer apart.
//	int - the number of nodes in the Tree
func DijkstraSearch(graph map[int]*Node, rootId int) ([][]Edge, bool, int) {
	hasCycle := func(nodeId int, currentLayer int, layers map[int]NodeLayerTuple) bool {
		l := layers[nodeId]
		if currentLayer-2 >= l.layer { //there is a cycle
			return true
		} else {
			return false
		}
	}

	var tree = [][]Edge{}
	l0 := []Edge{{u: -1, v: rootId}}

	//A lookup map so you can look up whether or not a Node has been seen and if so what layer it is in.
	layersLookup := make(map[int]NodeLayerTuple, len(graph))
	layersLookup[rootId] = NodeLayerTuple{
		Id:    rootId,
		layer: 0,
	}
	tree = append(tree, l0)

	var graphHasACycle = false
	var i = 0 //current layer you are finding edges for
	for {
		var pendingLayer []Edge
		for _, k := range tree[i] {
			node, _ := graph[k.v]
			for _, m := range node.Connections {
				//Lookup tail(v) of every edge in the layer to see if it has been seen before. If not add it to pending layer.
				_, alreadySeen := layersLookup[m.Id]
				if !alreadySeen { //TODO make this a function that changes the layer to this layer if the total length is less
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
			tree = append(tree, pendingLayer)
			i = i + 1
		} else {
			break
		}
	}
	return tree, graphHasACycle, len(layersLookup)
}

func TreeEquality(a, b []Edge) bool {
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
	bfsTree, hasCycle, numNodes := DijkstraSearch(graph, rootNode)
	numEdgesInTree := func(tree [][]Edge) int {
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

func DirectedGraphGen(lower, upperExc int) func(propcheck.SimpleRNG) (propcheck.Pair[map[int]*Node, int], propcheck.SimpleRNG) {
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
		for _, node := range graph {
			var connections []*Node
			connectedNodeSize := len(nodeIds)
			connectionIds, rng3 = sets.ChooseSet(0, int(connectedNodeSize), propcheck.ChooseInt(0, int(connectedNodeSize)), lt, eq)(rng3)
			for _, connectedNodeId := range connectionIds {
				if node.Id != graph[nodeIds[connectedNodeId]].Id {
					connections = append(connections, graph[nodeIds[connectedNodeId]])
				}
			}
			node.Connections = connections
		}
		root, rng4 := propcheck.ChooseInt(0, len(graph))(rng3)
		return propcheck.Pair[map[int]*Node, int]{graph, nodeIds[root]}, rng4
	}
}
