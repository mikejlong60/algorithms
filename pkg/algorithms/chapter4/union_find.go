package chapter4

import (
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type UNode struct {
	Id       string   //The is the label of the DIT object
	Set      *UNode   //This is the parent set of this node, empty if top of DIT and non-empty otherwise
	Children []*UNode //This is empty for a leaf node, non-empty otherwise
}

// Returns a union-find data structure on set S(make sure S is really a set) where all elements are in separate sets.
func MakeUnionFind(S []string) []*UNode {
	var r = make([]*UNode, len(S))
	for i, j := range S {
		r[i] = &UNode{Id: j, Set: nil, Children: []*UNode{}}
	}
	return r
}

// For an element u that is a member of some set S, Find(u) returns the name of the set containing u
func Find(u *UNode) string {
	if u.Set == nil { //You have reached the set that is housing u
		return u.Id
	} else { //Keep looking up until you reach the containing set
		return Find(u.Set)
	}
}

// For two sets A and B Union(A, B) merges the set B into the set A
func Union(A, B *UNode) *UNode {
	if Find(A) != Find(B) { //Make B a member of set A because it is not a member of set A
		B.Set = A
		A.Children = append(A.Children, B)
	}
	return A
}

// Creates a set of RDNs from the whole list of User DNs(i.e. cn=test tester10,ou=people,ou=fred,ou=bigfoot,o=u.s. government,c=us)
func MakeSetOfRDNs(users []string) []string {
	var rdns = []string{}
	for _, j := range users {
		a := strings.Split(j, ",")
		for _, k := range a {
			rdns = append(rdns, k)
		}
	}
	ff := map[string]struct{}{}

	for _, b := range rdns {
		ff[b] = struct{}{}
	}
	r := []string{}
	for aa, _ := range ff {
		r = append(r, aa)
	}
	return r
}

// Makes a DIT from a list of User DNs(i.e. cn=test tester10,ou=people,ou=fred,ou=bigfoot,o=u.s. government,c=us)
// Growth of algorithm is linear, O(n) where n is the number of users.
func ToDirectoryInformationTree(users []string) map[string]*UNode {
	start := time.Now()
	a := MakeSetOfRDNs(users) //Splits up big RDN for each user into a set of strings, meaning no duplicates
	s := MakeUnionFind(a)     //Makes a UNode for every member of set a above.

	var i = make(map[string]*UNode, len(s))
	//Turns set s into a map so you can lookup tokens in O(1)
	for _, k := range s {
		i[k.Id] = k
	}

	for _, j := range users { //Builds the whole DIT by unioning each object in the User DN.
		aa := strings.Split(j, ",")
		for index := range aa {
			if index+1 < len(aa) {
				Union(i[aa[index+1]], i[aa[index]])
			}
		}
	}
	log.Infof("ToDirectoryInformationTree for %v userDNs  took:%v", len(users), time.Since(start))
	return i
}

type Edge struct {
	U string //the Id of the beginning node of the edge
	V string //the Id of the ending node of the edge
}

//Depth-First search
// A recursive algorithm for depth-first search.
//Params:
//  U - *Node the current node that gets explored by the algorithm
//  seen - seen map[int]*Node - the accumulated map of Nodes that the algorithm has seen thus far
//  tree- an array of Edges reflecting the current dfs tree to this point
//Returns:
//  U - *Node the current node that gets explored by the algorithm
//  seen - seen map[int]*Node - the accumulated map of Nodes that the algorithm has seen thus far
//  tree- an array of Edges reflecting the current dfs tree to this point

func DFSearch(u *UNode, seen map[string]*UNode, tree []Edge) (*UNode, map[string]*UNode, []Edge) {
	seen[u.Id] = u
	for _, connectedNode := range u.Children {
		_, explored := seen[connectedNode.Id]
		if !explored {
			tree = append(tree, Edge{u.Id, connectedNode.Id})
			_, seen, tree = DFSearch(connectedNode, seen, tree)
		}
	}
	return u, seen, tree
}

// Given a DIT produces the complete list of strings that produced the DIT.
// This is isomorphic with the ToDirectoryInformationTree function above
func FromDirectoryInformationTree(dit map[string]*UNode, rootId string) (*UNode, map[string]*UNode, []Edge) {
	start := time.Now()
	//Start at the top of the tree
	root := dit[rootId]
	//Maybe do a Depth-first-search starting there and stop at leaf and add the whole path as the complete DN

	var tree []Edge
	//r, s, t :=
	dit2, allNodes, edges := DFSearch(root, make(map[string]*UNode), tree)
	log.Infof("FromDirectoryInformationTree took:%v", time.Since(start))
	return dit2, allNodes, edges
	//log.Info(r)
	//log.Info(s)
	//log.Info(t)

	//return []string{}
}
