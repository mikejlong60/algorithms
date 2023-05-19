package chapter4

import (
	"github.com/greymatter-io/golangz/linked_list"
	"math"
)

type Node3 struct {
	Id       string
	Distance int
}

type Node4 struct {
	Id          string
	Connections []*Node3
}

type Pq struct {
	Distance int
	Id       string
}

//def shortest_path(graph, start):
//distances = {node: float('inf') for node in graph}
//distances[start] = 0

func shortest_path(graph map[string]*Node3, start string) {

	var distances map[string]*Node3

	for i, j := range graph {
		j.Distance = math.MaxInt64
		distances[i] = j
	}

	distances[start].Distance = 0

	var pq *linked_list.LinkedList[*Pq]
	pq = linked_list.Push(&Pq{0, distances[start].Id}, pq)
	for linked_list.Len(pq) > 0 {
		current := linked_list.Head(pq)
		pq = linked_list.Drop(pq, 1) //Tail of the list is now current list
		if current.Distance > distances[current.Id].Distance {
			//Do something equivalent to Continue. Go to the next Head of Pq
		} else {
			//		for i, x := range dis
		}

	}
} //Continued below

/////////////////////////////////////////////
/////////////////////////////////////////

//	for pq != nil {
//		current_distance, current_node = linked_list.heappop(pq)
//if current_distance > distances[current_node]:
//continue
//for neighbor, weight in graph[current_node].items():
//distance = current_distance + weight
//if distance < distances[neighbor]:
//distances[neighbor] = distance
//heapq.heappush(pq, (distance, neighbor))
//
//return distances
//}
//graph = {
//'A': {'B': 1, 'C': 4},
//'B': {'C': 2, 'D': 5},
//'C': {'D': 1},
//'D': {}
//}
//
//shortest_path(graph, 'B')
//
