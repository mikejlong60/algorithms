package chapter4

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  int
}

var totalSteps int

func LocalMinimum(currentMin int, guess *Node) int {
	totalSteps = totalSteps + 1
	lesserValue := func(currentMin int, guess int) int {
		if currentMin < guess {
			return currentMin
		} else {
			return guess
		}
	}

	x := lesserValue(currentMin, guess.Parent.Value)
	if guess.Left == nil || guess.Right == nil { //At a leaf of the complete binary tree
		return lesserValue(x, guess.Value)
	} else { //Divide and Conquer until you get to a leaf
		l := LocalMinimum(x, guess.Left)
		r := LocalMinimum(x, guess.Right)
		return lesserValue(l, r)
	}
}
