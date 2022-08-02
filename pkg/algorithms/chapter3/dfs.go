package chapter3

//Depth-First search
// A recursive algorithm for depth-first search.
//Params:
//  u - *Node the current node that gets expored by the algorithm
//  seen - seen map[int]*Node - the accumlated map of Nodes that the algorithm has seen thus far
//  tree- an array of Edges reflecting the current dfs tree to this point
//Returns:
//  u - *Node the current node that gets expored by the algorithm
//  seen - seen map[int]*Node - the accumlated map of Nodes that the algorithm has seen thus far
//  tree- an array of Edges reflecting the current dfs tree to this point

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
