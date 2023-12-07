package chapter4

import (
	"testing"
)

// Algorithm  //TODO
// 1. Build a balanced binary tree(even number of leaf nodes) using Huffman with equal probabilities for each letter.  Assign a ramdom integer
// that serves as a weight to each edge in tree as you build it.
// 2. Add a new distance attribute to each Frequency type.  As you build the binary tree keep a weight sum at each node,  the
// sum of the weights leading to a given node.
// 3. Sort the leaf nodes(l and r attributes are nil for a leaf) in an array xs by highest weight sum descending.
// 4. Iterate over xs and increase the edge above each leaf node so that its sum equals the first leaf node(the most expensive).
func TestZeroSkew(t *testing.T) {
	a := Frequency{
		.20, "a", nil, nil,
	}
	b := Frequency{
		.20, "b", nil, nil,
	}
	c := Frequency{
		.20, "c", nil, nil,
	}
	d := Frequency{
		.20, "d", nil, nil,
	}
	//e := Frequency{
	//	.20, "e", nil, nil,
	//}
	f := []*Frequency{&a, &b, &c, &d} //, &e}
	insertIntoHeap := func(xss []*Frequency) []*Frequency {
		var r = StartHeapF(4)
		for _, x := range xss {
			r = HeapInsertF(r, x, frequencyLt)
		}
		return r
	}
	var freq = insertIntoHeap(f)

	freq = Huffman(freq, frequencyLt)
	if len(freq) != 1 {
		t.Errorf("Expected freq to be len 1 but was:%v", len(freq))
	}

	expected := Frequency{probability: 1, letter: "((c:(e:d)):(b:a))"}
	if !(freq[0].letter == expected.letter && freq[0].probability == 1) {
		t.Errorf("Expected freq to be a tree with the combined letters ((c:(e:d)):(b:a)) but was:%v", freq)
	}

}
