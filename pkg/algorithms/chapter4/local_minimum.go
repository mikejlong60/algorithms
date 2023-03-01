package chapter4

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  int
}

var totalSteps int

func localMinimum(currentMin int, guess *Node) int {
	totalSteps = totalSteps + 1
	lesserValue := func(currentMin, guess int) int {
		if currentMin < guess {
			return currentMin
		} else {
			return guess
		}
	}
	var x int
	if guess.Parent != nil {
		x = lesserValue(currentMin, guess.Parent.Value)
	} else {
		x = currentMin
	}

	if guess.Left == nil || guess.Right == nil { //At a leaf of the complete binary tree
		return lesserValue(x, guess.Value)
	} else { //Divide and Conquer until you get to a leaf
		l := localMinimum(x, guess.Left)
		r := localMinimum(x, guess.Right)
		return lesserValue(l, r)
	}
}

func LocalMinimum(guess *Node) int {
	return localMinimum(guess.Value, guess)
}
