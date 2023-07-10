package chapter2

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
func HeapifyUp[A any](heap []A, i int, lt func(l, r A) bool, zeroVal A) []A {
	if len(heap) == 0 {
		return []A{}
	}
	r := make([]A, len(heap))
	copy(r, heap)

	if i > 0 {
		j := ParentIdx(i)
		if lt(r[i], r[j]) {
			//Swap elements
			temp := r[i]
			temp2 := r[j]
			r[j] = temp
			r[i] = temp2
			r = HeapifyUp(r, j, lt, zeroVal)
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
func HeapifyDown[A any](heap []A, zeroI int, lt func(l, r A) bool) []A {
	n := len(heap)
	if n == 0 {
		return []A{}
	}
	r := make([]A, len(heap))
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
//	zeroVal A - the  zero value for an element A in the heap.
//
// Returns - A new heap (as a slice) of size n that has every element initialized to the zero value
// Performance - O(N)
func StartHeap[A any](n int, zeroVal A) []A {
	var x = make([]A, n)
	for i := range x {
		x[i] = zeroVal
	}
	return x
}

// This is a pure function.
// Parameters:
//
//	heap []A - the slice that is holding the heap
//
// Returns -the minimum element in the given heap without removing it. O(1)
// Performance - O(1)
func FindMin[A any](heap []A) A {
	if len(heap) == 0 {
		panic("heap is empty. FindMin is therefore irrelevant.")
	}
	return heap[0]
}

//	h []A - the slice that is holding the heap
//	isZeroVal func(a A) bool - A predicate function that returns whether or not any A value in the heap is the zero value set in StartHeap
//
// Returns - The index of the first empty slot in the heap
// Performance - O(N)
func findFirstEmptySlotInHeap[A any](h []A, isZeroVal func(a A) bool) int {
	for i, x := range h {
		if isZeroVal(x) {
			return i
		}
	}
	return -1
	//No non-empty elements so return the index of the last element
	//return len(h) - 1
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
//	isZeroVal func(a A) bool - A predicate function that returns whether or not any A value in the heap is the zero value set in StartHeap
//
// Returns - A new heap (as a slice) that has the given element in its proper position in the heap
// Performance - O(log N)
func HeapInsert[A any](heap []A, a A, lt func(l, r A) bool, isZeroVal func(a A) bool, zeroVal A) []A {
	if len(heap) == 0 {
		return []A{}
	}
	r := make([]A, len(heap))
	copy(r, heap)

	l := findFirstEmptySlotInHeap(r, isZeroVal)
	r[l] = a
	return HeapifyUp(r, l, lt, zeroVal)
}

// Deletes an element from the given heap and returns a new heap. This is a pure function.
// Parameters:
//
//	heap []A - the slice that is holding the heap
//	i int - the index into the heap of the element you want to delete. Array indices start with the number zero.
//	lt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//	isZeroVal func(a A) bool - A predicate function that returns whether or not any A value in the heap is the zero value set in StartHeap
//	zeroVal A - the  zero value for an element A in the heap. The operation sets the i'th element to the zeroVal
//
// Returns - A new heap that has the given element in its proper position in the heap
// Performance - O(log N)
func HeapDelete[A any](heap []A, i int, lt func(l, r A) bool, isZeroVal func(a A) bool, zeroVal A) []A {
	n := len(heap)
	if n == 0 {
		return []A{}
	}
	r := make([]A, len(heap))
	copy(r, heap)

	//Zero out the given element
	r[i] = zeroVal
	parent := ParentIdx(i)
	if parent > 0 && (lt(r[i], r[parent]) || !isZeroVal(r[i])) {
		r = HeapifyUp(r, i, lt, zeroVal)
	} else {
		r = HeapifyDown(r, i, lt)
	}
	return r
}
