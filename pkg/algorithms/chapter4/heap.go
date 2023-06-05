package chapter4

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
// This is a pure function because I copy the slice each time.  r only gets mutated internally which is OK because its a copy.
// Parameters:
//
//	heap []A - the slice that is holding the heap
//	i int - the index into the heap of the element you want to move up. Array indices start with the number zero.
//	lt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//
// Returns - A new heap (as a slice) that has the i'th element in its proper position in the heap
// Performance - O(log N) assuming that the array is almost-a-heap with the key: heap(i) too small.
func HeapifyUp(heap []*Cache, i int, lt func(l, r *Cache) bool) []*Cache {
	if len(heap) == 0 {
		return []*Cache{}
	}
	r := make([]*Cache, len(heap))
	copy(r, heap)

	if i > 0 {
		j := ParentIdx(i)
		if lt(r[i], r[j]) {
			//Swap elements
			temp := r[i]
			temp2 := r[j]
			r[j] = temp
			r[i] = temp2
			r = HeapifyUp(r, j, lt)
		}
	}
	return r
}

// This is a pure function because I copy the array each time.  r only gets mutated internally which is OK because its a copy.
// Parameters:
//
//	heap []A - the slice that is holding the heap
//	zeroI int - the index into the heap of the element you want to move down. Array indices start with the number zero.
//	lt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//
// Returns - A new heap (as a slice) that has the i'th element in its proper position in the heap
// Performance - O(log N) assuming that the array is almost-a-heap with the key: heap(i) too big.
func HeapifyDown(heap []*Cache, zeroI int, lt func(l, r *Cache) bool) []*Cache {
	n := len(heap)
	if n == 0 {
		return []*Cache{}
	}
	r := make([]*Cache, len(heap))
	copy(r, heap)

	i := zeroI + 1

	var j int
	if 2*i > n {
		return r
	} else if 2*i < n {
		leftIdx := 2 * i
		rightIdx := (2 * i) + 1
		leftVal := r[leftIdx]
		rightVal := r[rightIdx]
		if lt(leftVal, rightVal) {
			j = leftIdx
		} else {
			j = rightIdx
		}
	} else if 2*i == n {
		j = 2 * i
	}
	if lt(r[j], r[i]) {
		//Swap elements
		temp := r[i]
		temp2 := r[j]
		r[j] = temp
		r[i] = temp2
		r = HeapifyDown(r, j, lt)
	}
	return r
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
// Returns - The index of the first empty slot in the heap, or -1 if there are no empty slots
// Performance - O(N)
func findFirstEmptySlotInHeap(h []*Cache) int {
	for i, x := range h {
		if x == nil {
			return i
		}
	}
	return -1
}

// Inserts the given element into the given heap and returns a new heap. This is a pure funciton. O(log n)
//
//	This is a pure function because I copy the array each time.
//
// Parameters:
//
//	heap []A - the slice that is holding the heap
//	a  A - the element you want to insert into the heap
//	lt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//
// Returns - A new heap (as a slice) that has the given element in its proper position in the heap
// Performance - O(log N)
func HeapInsert(heap []*Cache, a *Cache, lt func(l, r *Cache) bool) []*Cache {
	if len(heap) == 0 {
		return []*Cache{}
	}
	var r = make([]*Cache, len(heap))
	copy(r, heap)

	l := findFirstEmptySlotInHeap(r)
	if l == -1 { //No empty slot so add a new spot at the end and put the new element there
		r = append(r, nil)
		l = len(r) - 1
	}
	r[l] = a
	return HeapifyUp(r, l, lt)
}

// Deletes an element from the given heap and returns a new heap. This is a pure function.
// Parameters:
//
//	heap []A - the slice that is holding the heap
//	i int - the index into the heap of the element you want to delete. Array indices start with the number zero.
//	lt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//
// Returns - A new heap that has the given element in its proper position in the heap
// Performance - O(log N)
func HeapDelete(heap []*Cache, i int, lt func(l, r *Cache) bool) []*Cache {
	n := len(heap)
	if n == 0 {
		return []*Cache{}
	}
	r := make([]*Cache, len(heap))
	copy(r, heap)

	//Find first empty slot and then go back one to get the last non-empty slot and call it l.
	//Then move element l to the hole and zero out the original l element.
	//var l = findFirstEmptySlotInHeap(r)
	//if l == -1 {
	//	r = append(r, nil) //Add an empty slot at the end of the heap if it is currently full
	//	l = len(r) - 1     //Make that slot the new empty slot
	//}
	//Zero out i'th element
	//	r[i] = r[l]
	r[i] = r[len(r)-1]  //Move last element into slot you are deleting
	r = r[0 : len(r)-1] //Rip the last empty element off heap

	parent := ParentIdx(i)
	if parent > 0 && lt(r[i], r[parent]) {
		r = HeapifyUp(r, i, lt)
	} else {
		r = HeapifyDown(r, i, lt)
	}
	return r
}

//
//////
//
////Find first empty slot and then go back one to get the last non-empty slot and call it l.
////Then move element l to the hole and zero out the original l element.
//l := findFirstEmptySlotInHeap(r, isZeroVal) - 1
////Zero out the given element
//r[i] = r[l]
//r[l] = zeroVal
//parent := ParentIdx(i)
//if parent > 0 && lt(r[i], r[parent]) {
//
//}
/////
