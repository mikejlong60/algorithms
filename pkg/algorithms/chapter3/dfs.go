package chapter3

//Depth-First search with cycle detection
//Params:
//  graph a hashmap of all the nodes in te graph. Facilitates n log n lookup
//  rootId the Node Id of the root node, the one at the top of the mobile from which all the other nodes hang
//Returns:
//   Tree  - the search tree represented as an array of layers, each layer constisting of an array of Edges(u, v)
//   bool - whether or not the resulting search tree contained a cycle. A cycle is a relationship between two nodes that is farther than one layer apart.
//   int - the number of nodes in the Tree

func DFSearch(u *Node, seen map[int]*Node, tree []Edge) (*Node, map[int]*Node, []Edge) {
	seen[u.Id] = u
	for _, connectedNode := range u.Connections {
		_, explored := seen[connectedNode.Id]
		if !explored {
			tree = append(tree, Edge{u.Id, connectedNode.Id})
			_, seen, tree = DFSearch(connectedNode, seen, tree)
		}
	}
	return u, seen, tree
}
