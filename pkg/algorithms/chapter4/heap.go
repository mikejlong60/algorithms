package chapter4

import log "github.com/sirupsen/logrus"

type Cache struct {
	ts   int
	data string
}

// i int - the index in the given heap of the parent of element i. Array indices start with the number zero.
// Performance - O(1)
func ParentIdx(i int) int {
	//Odd number
	if i%2 > 0 {
		return i / 2
	} else { // even number
		return (i / 2) - 1
	}
}

// Definition of almost-a-heap. Only one node in the tree has a value less than it's parent as per the lt function and that
// node is at the bottom rung of the heap.
// Definition of a heap.  Every node in the tree has a greater value than it's parent as per the lt function.
// This is a not pure function
// Parameters:
//
//	heap []A - the slice that is holding the heap
//	i int - the index into the heap of the element you want to move up. Array indices start with the number zero.
//	lt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//
// Returns - The modified heap (as a slice) that has the i'th element in its proper position in the heap
// Performance - O(log N) assuming that the array is almost-a-heap with the key: heap(i) too small.
func HeapifyUp(heap []*Cache, i int, lt func(l, r *Cache) bool) []*Cache {
	if len(heap) == 0 {
		return []*Cache{}
	}
	if i > 0 {
		j := ParentIdx(i)
		if lt(heap[i], heap[j]) {
			//Swap elements
			temp := heap[i]
			temp2 := heap[j]
			heap[j] = temp
			heap[i] = temp2
			heap = HeapifyUp(heap, j, lt)
		}
	}
	return heap
}

// This is not a pure function because it modified the array each time.
// Parameters:
//
//	heap []A - the slice that is holding the heap
//	zeroI int - the index into the heap of the element you want to move down. Array indices start with the number zero.TODO change
//	lt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//
// Returns - The original heap (as a slice) that has the i'th element in its proper position
// Performance - O(log N) assuming that the array is almost-a-heap with the key: heap(i) too big.
func HeapifyDown(heap []*Cache, zeroI int, lt func(l, r *Cache) bool) []*Cache {
	n := len(heap)
	if n == 0 {
		return []*Cache{}
	}

	i := zeroI + 1

	var j int
	if 2*i > n {
		return heap
	} else if 2*i < n {
		leftIdx := 2 * i
		rightIdx := (2 * i) + 1
		leftVal := heap[leftIdx]
		rightVal := heap[rightIdx]
		if lt(leftVal, rightVal) {
			j = leftIdx
		} else {
			j = rightIdx
		}
	} else if 2*i == n {
		j = 2 * i
	}
	if lt(heap[j], heap[i]) {
		//Swap elements
		temp := heap[i]
		temp2 := heap[j]
		heap[j] = temp
		heap[i] = temp2
		heap = HeapifyDown(heap, j, lt)
	}
	return heap
}

// Parameters:
//
//	n int - the size of the heap. This is fixed.
//
// Returns - A new heap (as a slice) of size n that has every element initialized to the zero value
// Performance - O(N)
func StartHeap(n int) []*Cache {
	var x = make([]*Cache, n)
	return x
}

// This is a pure function.
// Parameters:
//
//	heap []A - the slice that is holding the heap
//
// Returns -the minimum element in the given heap without removing it. O(1)
// Performance - O(1)
func FindMin(heap []*Cache) *Cache {
	if len(heap) == 0 {
		panic("heap is empty. FindMin is therefore irrelevant.")
	}
	return heap[0]
}

//	h []A - the slice that is holding the heap
//
// Returns - A pure function. The index of the first empty slot in the heap, or -1 if there are no empty slots
// Performance - O(N)
// TODO maybe you never have empty slots because you are using slices
func findFirstEmptySlotInHeap(h []*Cache) int {
	for i, x := range h {
		if x == nil {
			return i
		}
	}
	return -1
}

// Inserts the given element into the given heap and returns the modified heap.
//
// O(log n)
//
// Parameters:
//
//	heap []A - the slice that is holding the heap
//	a  A - the element you want to insert into the heap
//	lt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//
// Returns - The original heap (as a slice) that has the given element in its proper position
// Performance - O(log N)
func HeapInsert(heap []*Cache, a *Cache, lt func(l, r *Cache) bool) []*Cache {
	if len(heap) == 0 {
		return []*Cache{}
	}

	l := findFirstEmptySlotInHeap(heap)
	if l == -1 { //No empty slot so add a new spot at the end and put the new element there
		heap = append(heap, nil)
		l = len(heap) - 1
	}
	heap[l] = a
	return HeapifyUp(heap, l, lt)
}

// Deletes an element from the given heap. This is not a pure function.
// Parameters:
//
//	heap []A - the slice that is holding the heap
//	i int - the index into the heap of the element you want to delete. Array indices start with the number zero.
//	lt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//
// Returns - The original heap that has the given element in its proper position
// Performance - O(log N)
func HeapDelete(heap []*Cache, i int, lt func(l, r *Cache) bool) []*Cache {
	n := len(heap)
	if n == 0 {
		return []*Cache{}
	}

	if i > len(heap)-1 {
		log.Errorf("The element:%v you are trying to delete is longer than heap length: %v", i, len(heap)-1)
	}

	heap[i] = heap[len(heap)-1]  //Move last element into slot you are deleting
	heap = heap[0 : len(heap)-1] //Rip the last empty element off heap

	//If index you are trying to delete is one of the last two in the heap, then just return the heap(its underlying array)
	//with the last element missing because that cannot possibly violate the heap property that a child must be be greater
	//or equal to its parent.
	if i > len(heap)-2 {
		return heap
	}
	parent := ParentIdx(i)
	if parent > 0 && lt(heap[i], heap[parent]) {
		heap = HeapifyUp(heap, i, lt)
	} else {
		heap = HeapifyDown(heap, i, lt)
	}
	//heap = heap[0 : len(heap)-1] //Rip the last (the one you wanted to delete) empty element off heap
	return heap
}
