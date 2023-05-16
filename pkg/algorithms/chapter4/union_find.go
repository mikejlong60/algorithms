package chapter4

type UNode struct {
	Id  string
	Set *UNode
}

// Returns a union-find data structure on set S(make sure S is really a set) where all elements are in separate sets.
func MakeUnionFind(S []string) []*UNode {
	var r = make([]*UNode, len(S))
	for i, j := range S {
		r[i] = &UNode{j, nil}
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
	}
	return A
}
