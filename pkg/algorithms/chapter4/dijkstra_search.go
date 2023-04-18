package chapter4

import (
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"github.com/mikejlong60/algorithms/pkg/algorithms/chapter3"
)

// The ncluding the total distance from starting node to the current node to deternine the shortest
// path to it if you have seen it before that means that its in an earlier layer.  Therfore it is by definition
// the shortest distance from the root.

// Breadth-First search with shortest length from root node to every other node. The assumption is that
// every layer is a distance of 1 from the previous layer. So tbe fewer layers between the root and a given
// node, the shorter the distance. If the algorithm has seen the node before that means that its in an earlier layer.
// Therfore it is by definition the shortest distance from the root.
// Params:
//
//	graph a hashmap of all the nodes in the graph. Facilitates log n lookup
//	rootId the Node id of the root node, the one at the top of the mobile from which all the other nodes hang
//
// Returns:
//
//	Tree  - the search tree represented as an array of layers, each layer consisting of an array of Edges(u, v)
func DijkstraSearch(graph map[int]*chapter3.Node, rootId int) [][]chapter3.Edge {
	var tree = [][]chapter3.Edge{}
	l0 := []chapter3.Edge{{U: -1, V: rootId}}

	//A lookup map so you can look up whether or not a Node has been seen and if so what layer it is in.
	layersLookup := make(map[int]chapter3.NodeLayerTuple, len(graph))
	layersLookup[rootId] = chapter3.NodeLayerTuple{
		Id:    rootId,
		Layer: 0,
	}
	tree = append(tree, l0)

	var i = 0 //current layer you are finding edges for
	for {
		var pendingLayer []chapter3.Edge
		for _, k := range tree[i] {
			node, _ := graph[k.V]
			for _, m := range node.Connections {
				//totalSteps2 = totalSteps2 + 1
				//Lookup tail(v) of every edge in the layer to see if it has been seen before. If not add it to pending layer.
				//If it has been seen, it is already the shortest path from the root.
				_, alreadySeen := layersLookup[m.Id]
				if !alreadySeen {
					pendingLayer = append(pendingLayer, chapter3.Edge{U: k.V, V: m.Id})
					layersLookup[m.Id] = chapter3.NodeLayerTuple{Id: m.Id, Layer: i + 1}
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
	return tree
}

func DirectedGraphGen(lower, upperExc int) func(propcheck.SimpleRNG) (propcheck.Pair[map[int]*chapter3.Node, int], propcheck.SimpleRNG) {
	return func(rng propcheck.SimpleRNG) (propcheck.Pair[map[int]*chapter3.Node, int], propcheck.SimpleRNG) {
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

		var nodeIds []int
		nodeIds, rng = sets.ChooseSet(lower, upperExc, propcheck.ChooseInt(0, 1000000), lt, eq)(rng)
		graph := make(map[int]*chapter3.Node, len(nodeIds))
		for _, j := range nodeIds {
			graph[j] = &chapter3.Node{Id: j}
		}

		var connectionIds []int
		for _, node := range graph {
			var connections []*chapter3.Node
			connectedNodeSize := len(nodeIds)
			connectionIds, rng = sets.ChooseSet(0, int(connectedNodeSize), propcheck.ChooseInt(0, int(connectedNodeSize)), lt, eq)(rng)
			for _, connectedNodeId := range connectionIds {
				if node.Id != graph[nodeIds[connectedNodeId]].Id {
					connections = append(connections, graph[nodeIds[connectedNodeId]])
				}
			}
			node.Connections = connections
		}
		var root int
		root, rng = propcheck.ChooseInt(0, len(graph))(rng)
		return propcheck.Pair[map[int]*chapter3.Node, int]{graph, nodeIds[root]}, rng
	}
}
