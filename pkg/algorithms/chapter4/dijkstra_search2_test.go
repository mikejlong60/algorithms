package chapter4

import (
	"github.com/greymatter-io/golangz/arrays"
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

func Dijkstra2(graph map[int]*node, rootId int) [][]chapter3.Edge {
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

	n0 := node{
		id: 0,
	}
	n1 := node{
		id:                 1,
		distanceFromParent: 3,
		parent:             &n0,
	}
	n2 := node{
		id:                 2,
		distanceFromParent: 4,
		parent:             &n0,
	}
	n3 := node{
		id:                 3,
		distanceFromParent: 5,
		parent:             &n1,
	}
	n4 := node{
		id:                 4,
		distanceFromParent: 6,
		parent:             &n1,
	}
	n0.connections = []*node{&n1, &n2, &n3, &n4}
	n1.connections = []*node{&n3, &n4}
	graph := make(map[int]*node, 5) //First field of pair is the layer the node is in, -1 means it's never been seen before and is thus not in any layer
	graph[0] = &n0
	graph[1] = &n1
	graph[2] = &n2
	graph[3] = &n3
	graph[4] = &n4
	actual := Dijkstra2(graph, 0)
	expected := [][]chapter3.Edge{{{-1, 0}}, {{0, 1}, {0, 2}, {0, 3}, {0, 4}}}
	if !arrays.ArrayEquality(actual, expected, chapter3.TreeEquality) {
		t.Errorf("Actual:%child Expected:%child", actual, expected)
	}
}
