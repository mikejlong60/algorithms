package chapter4

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	set := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	b := MakeUnionFind(set)

	Union(b[0], b[1])
	if Find(b[1]) != "A" {
		t.Errorf("Actual:%v Expected:%v", Find(b[1]), "A")
	}
	if b[0].Set != nil {
		t.Errorf("Actual:%v Expected:%v", b[0].Set, nil)
	}

	Union(b[0], b[2])

	if Find(b[2]) != "A" {
		t.Errorf("Actual:%v Expected:%v", Find(b[2]), "A")
	}

	//Now make a set out of elements D - H
	Union(b[3], b[4])
	Union(b[3], b[5])
	Union(b[3], b[6])
	Union(b[3], b[7])
	if Find(b[4]) != "D" {
		t.Errorf("Actual:%v Expected:%v", Find(b[4]), "D")
	}
	if Find(b[5]) != "D" {
		t.Errorf("Actual:%v Expected:%v", Find(b[5]), "D")
	}
	if Find(b[6]) != "D" {
		t.Errorf("Actual:%v Expected:%v", Find(b[6]), "D")
	}
	if Find(b[7]) != "D" {
		t.Errorf("Actual:%v Expected:%v", Find(b[7]), "D")
	}

	//Now merge the sets A and D into set A
	Union(b[0], b[3])

	if b[0].Set != nil {
		t.Errorf("Actual:%v Expected:%v", b[0].Set, nil)
	}
	if Find(b[1]) != "A" {
		t.Errorf("Actual:%v Expected:%v", Find(b[1]), "A")
	}
	if Find(b[2]) != "A" {
		t.Errorf("Actual:%v Expected:%v", Find(b[2]), "A")
	}
	if Find(b[3]) != "A" {
		t.Errorf("Actual:%v Expected:%v", Find(b[3]), "A")
	}
	if Find(b[4]) != "A" {
		t.Errorf("Actual:%v Expected:%v", Find(b[4]), "A")
	}
	if Find(b[5]) != "A" {
		t.Errorf("Actual:%v Expected:%v", Find(b[5]), "A")
	}
	if Find(b[6]) != "A" {
		t.Errorf("Actual:%v Expected:%v", Find(b[6]), "A")
	}
	if Find(b[7]) != "A" {
		t.Errorf("Actual:%v Expected:%v", Find(b[7]), "A")
	}
}
