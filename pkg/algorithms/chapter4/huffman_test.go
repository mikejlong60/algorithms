package chapter4

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

func frequencyLt(l, r *Frequency) bool {
	if l.probability < r.probability {
		return true
	} else {
		return false
	}
}

func frequencyEq(l, r *Frequency) bool {
	if l.probability == r.probability {
		return true
	} else {
		return false
	}
}

func TestHuffmanHeapFromBook(t *testing.T) {
	a := Frequency{
		.32, "a",
	}
	b := Frequency{
		.25, "b",
	}
	c := Frequency{
		.20, "c",
	}
	d := Frequency{
		.18, "d",
	}
	e := Frequency{
		.05, "e",
	}
	f := []*Frequency{&a, &b, &c, &d, &e}
	insertIntoHeap := func(xss []*Frequency) []*Frequency {
		var r = StartHeapF(5)
		for _, x := range xss {
			r = HeapInsertF(r, x, frequencyLt)
		}
		return r
	}
	//var actual *[]*Frequency
	var actual = insertIntoHeap(f)
	var expected = []*Frequency{&e, &d, &b, &a, &c}
	if !arrays.ArrayEquality(actual, expected, frequencyEq) {
		t.Errorf("Actual:%v, \nexpected:%v", actual, expected)
	}
	actual, _ = HeapDeleteF(actual, 0, frequencyLt)
	expected = []*Frequency{&d, &c, &b, &a}
	if !arrays.ArrayEquality(actual, expected, frequencyEq) {
		t.Errorf("Actual:%v, \n              expected:%v", actual, expected)
	}
	actual, _ = HeapDeleteF(actual, 0, frequencyLt)
	expected = []*Frequency{&c, &b, &a}
	if !arrays.ArrayEquality(actual, expected, frequencyEq) {
		t.Errorf("Actual:%v, \n              expected:%v", actual, expected)
	}
	actual, _ = HeapDeleteF(actual, 0, frequencyLt)
	expected = []*Frequency{&b, &a}
	if !arrays.ArrayEquality(actual, expected, frequencyEq) {
		t.Errorf("Actual:%v, \n              expected:%v", actual, expected)
	}
	actual, _ = HeapDeleteF(actual, 0, frequencyLt)
	expected = []*Frequency{&a}
	if !arrays.ArrayEquality(actual, expected, frequencyEq) {
		t.Errorf("Actual:%v, \n              expected:%v", actual, expected)
	}
	actual, _ = HeapDeleteF(actual, 0, frequencyLt)
	expected = []*Frequency{}
	if !arrays.ArrayEquality(actual, expected, frequencyEq) {
		t.Errorf("Actual:%v, \n              expected:%v", actual, expected)
	}
}

func TestHuffmanHeapFromBook2(t *testing.T) {
	a := Frequency{
		.32, "a",
	}
	b := Frequency{
		.25, "b",
	}
	c := Frequency{
		.20, "c",
	}
	d := Frequency{
		.18, "d",
	}
	e := Frequency{
		.05, "e",
	}
	f := []*Frequency{&a, &b, &c, &d, &e}
	insertIntoHeap := func(xss []*Frequency) []*Frequency {
		var r = StartHeapF(5)
		for _, x := range xss {
			r = HeapInsertF(r, x, frequencyLt)
		}
		return r
	}
	var freq = insertIntoHeap(f)
	//var expected = []*Frequency{&e, &d, &b, &a, &c}

	freq, enc := Huffman(freq, []*Frequency{}, frequencyLt)
	if len(freq) != 0 {
		t.Errorf("Expected freq to be len 0 but was:%v", len(freq))
	}
	fmt.Println(enc)
}
