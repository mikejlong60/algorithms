package chapter3

import (
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
)

type Node struct {
	Id          int
	Connections []Node
}

type BFS = [][]Edge

type Edge struct {
	u int //the Id of the beginning node of the edge
	v int //the Id of the ending node of the edge
}

type NodeLayer struct {
	Id    int //The node Id
	layer int // Zero indexed array index indicating the layer(1st array index) in the BFS array the node lives
}

//Breadth-First search with cycle detection
//Params:
//  graph a hashmap of all the nodes in te graph. Facilitates n log n lookup
//  rootId the Node Id of the root node, the one at the top of the mobile from which all the other nodes hang
//Returns:
//   BFS  - the search tree represented as an array of layers, each layer constisting of an array of Edges(u, v pairs)
//   bool - whether or not the resulting search tree contained a cycle. A cycle is a relationship between two nodes that is farther than one layer apart.
func BFSearch(graph map[int]Node, rootId int) (BFS, bool) {
	hasCycle := func(nodeId int, currentLayer int, layers map[int]NodeLayer) bool {
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
	layersLookup := make(map[int]NodeLayer, len(graph))
	layersLookup[rootId] = NodeLayer{
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
					layersLookup[m.Id] = NodeLayer{Id: m.Id, layer: i + 1}
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
	return bfs, graphHasACycle
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

//TODO make a "real" that creates random graphs generator. You have alredy done this on your other machine but it's not in the repo yet because thr maching is getting replaced.
// Generates a graph of Node structures with A size in the indicated range using the given Gen
func GraphGen() func(propcheck.SimpleRNG) (map[int]Node, propcheck.SimpleRNG) {
	return func(rng propcheck.SimpleRNG) (map[int]Node, propcheck.SimpleRNG) {

		n0 := Node{
			Id:          0,
			Connections: nil,
		}
		n1 := Node{
			Id:          1,
			Connections: nil,
		}
		n2 := Node{
			Id:          2,
			Connections: nil,
		}
		n3 := Node{
			Id:          03,
			Connections: nil,
		}
		n4 := Node{
			Id:          4,
			Connections: nil,
		}
		n5 := Node{
			Id:          5,
			Connections: nil,
		}
		n6 := Node{
			Id:          6,
			Connections: nil,
		}
		n7 := Node{
			Id:          7,
			Connections: nil,
		}
		n0.Connections = []Node{n1, n2}
		n1.Connections = []Node{n4}
		n2.Connections = []Node{n3}
		n3.Connections = []Node{n5, n6, n7}
		n7.Connections = []Node{n7}
		n5.Connections = []Node{n7}
		n6.Connections = []Node{n4}
		graph := make(map[int]Node, 7) //First field of pair is the layer the node is in, -1 means it's never been seen before and is thus not in any layer
		graph[0] = n0
		graph[1] = n1
		graph[2] = n2
		graph[3] = n3
		graph[4] = n4
		graph[5] = n5
		graph[6] = n6
		graph[7] = n7
		return graph, rng
	}
}
