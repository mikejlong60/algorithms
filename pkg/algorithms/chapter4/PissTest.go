package chapter4

import (
	//	"github.com/greymatter-io/golangz/arrays"
	"github.com/mikejlong60/algorithms/pkg/algorithms/chapter3"
	"testing"
)

type node struct {
	id                 int
	parent             *node
	distanceFromParent int
	connections        []*node
}

func computeDistanceFromRoot(current *node, length int) int {
	if current.parent != nil {
		computeDistanceFromRoot(current.parent, length+current.distanceFromParent)
	}
	return length + current.distanceFromParent
}

func piss(graph map[int]*node, rootId int) [][]chapter3.Edge {
	var tree = [][]chapter3.Edge{}
	l0 := []chapter3.Edge{{U: -1, V: rootId}}

	//A lookup map so you can look up whether or not a Node has been seen and if so what layer it is in.
	layersLookup := make(map[int]chapter3.NodeLayerTuple, len(graph))
	layersLookup[rootId] = chapter3.NodeLayerTuple{
		Id:               rootId,
		Layer:            0,
		DistanceFromRoot: 0,
	}
	tree = append(tree, l0)

	var i = 0 //current layer you are finding edges for
	for {
		var pendingLayer []chapter3.Edge
		for _, k := range tree[i] {
			node, _ := graph[k.V]

			for _, m := range node.connections {
				//Lookup the node in the layer to see if it has been seen before. If not or the current node is closer to root then add it to pending layer.
				currentNodeInLayer, alreadySeen := layersLookup[m.id]
				totalSteps2 = totalSteps2 + 1
				distanceFromRoot := computeDistanceFromRoot(m, m.distanceFromParent)
				if !alreadySeen || currentNodeInLayer.DistanceFromRoot > distanceFromRoot {
					pendingLayer = append(pendingLayer, chapter3.Edge{U: k.V, V: m.id})
					layersLookup[m.id] = chapter3.NodeLayerTuple{Id: m.id, Layer: i + 1, DistanceFromRoot: distanceFromRoot}
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

func TestPiss(t *testing.T) {

	//n0 := Node2{
	//	Id:          0,
	//	Connections: nil,
	//}
	//n1 := Node2{
	//	Id:          1,
	//	Connections: nil,
	//}
	//n2 := Node2{
	//	Id:          2,
	//	Connections: nil,
	//}
	//n3 := Node2{
	//	Id:          03,
	//	Connections: nil,
	//}
	//n4 := Node2{
	//	Id:          4,
	//	Connections: nil,
	//}
	//n0.Connections = []NodeDistanceTuple{{node: &n1, parent: &n0, distance: 3}, {parent: &n0, node: &n2, distance: 4}, {parent: &n0, node: &n3, distance: 5}, {parent: &n0, node: &n4, distance: 5}}
	//n1.Connections = []NodeDistanceTuple{{parent: &n1, node: &n3, distance: 12}, {parent: &n1, node: &n4, distance: 3}}
	//n2.Connections = []NodeDistanceTuple{{parent: &n2, node: &n3, distance: 1}, {parent: &n2, node: &n3, distance: 3}}
	//n3.Connections = []NodeDistanceTuple{{parent: &n3, node: &n4, distance: 10}, {parent: &n3, node: &n4, distance: 2}}
	//n4.Connections = []NodeDistanceTuple{{parent: &n4, node: &n0, distance: 13}}
	//graph := make(map[int]*Node2, 5) //First field of pair is the layer the node is in, -1 means it's never been seen before and is thus not in any layer
	//graph[0] = &n0
	//graph[1] = &n1
	//graph[2] = &n2
	//graph[3] = &n3
	//graph[4] = &n4
	//actual := DijkstraSearch2(graph, 0)
	//expected := [][]chapter3.Edge{{{-1, 0}}, {{0, 1}, {0, 2}, {0, 3}, {0, 4}}}
	//if !arrays.ArrayEquality(actual, expected, chapter3.TreeEquality) {
	//	t.Errorf("Actual:%child Expected:%child", actual, expected)
	//}
}
