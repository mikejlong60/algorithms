package chapter2

//Key for now is the array index, starting with zero.
func parentIdx(i int) int {
	//Odd number
	if i%2 > 0 {
		return i / 2
	} else {
		return (i / 2) - 1
	}
}

//Definition of almost-a-heap. Only one node in the tree has a value less than it's parent as per the lt function and that
//node is at the bottom rung of the heap.
//Definition of a heap.  Every node in the tree has a greater value than it's parent as per the lt function.
//This is a pure function because I copy the array each time.  r only gets mutated internally which is OK because its a copy.
//O(log N) assuming that the array is almost-a-heap with the key: heap(i) too small.
func HeapifyUp[A any](heap []A, i int, lt func(l, r A) bool) []A {
	if len(heap) == 0 {
		return []A{}
	}
	r := make([]A, len(heap))
	copy(r, heap)

	if i > 1 {
		j := parentIdx(i)
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

//This is a pure function because I copy the array each time.  r only gets mutated internally which is OK because its a copy.
//O(log N) assuming that the array is almost-a-heap with the key: heap(i) too big.
func HeapifyDown[A any](heap []A, i int, lt func(l, r A) bool) []A {
	if len(heap) == 0 {
		return []A{}
	}
	r := make([]A, len(heap))
	copy(r, heap)

	var j int
	n := len(r)
	if 2*i > n {
		return r
	} else if 2*i < n {
		leftIdx := 2 * i
		rightIdx := 2*i + 1
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
	val1 := r[j]
	val2 := r[i]
	if lt(val1, val2) {
		//Swap elements
		temp := r[i]
		temp2 := r[j]
		r[j] = temp
		r[i] = temp2
		r = HeapifyDown(r, j, lt)
	}
	return r
}

//Make an empty heap to store at most n elements. O(N)
func StartHeap[A any](n int) []A {
	return make([]A, n)
}

//Returns the minimum element in the given heap without removing it. O(1)
func FindMin[A any](heap []A) A {
	return heap[0]
}

//Inserts an element into the given heap and returns a new heap. This is a pure funciton. O(log n)
func HeapInsert[A any](heap []A, a A, lt func(l, r A) bool, isZeroVal func(a A) bool) []A {
	if len(heap) == 0 {
		return []A{}
	}
	r := make([]A, len(heap))
	copy(r, heap)

	findFirstEmptySlotInHeap := func(h []A) int {
		for i, x := range h {
			if isZeroVal(x) {
				return i
			}
		}
		panic("heap is already full. You must copy the old heap to a new larger heap and use that to HeapInsert your new element.")
	}
	l := findFirstEmptySlotInHeap(r)
	r[l] = a
	return HeapifyUp(r, l, lt)
}

//Deletes an element into the given heap and returns a new heap. This is a pure funciton. O(log n)
func HeapDelete[A any](heap []A, i int, lt func(l, r A) bool, zeroVal A) []A {
	if len(heap) == 0 {
		return []A{}
	}
	r := make([]A, len(heap))
	copy(r, heap)

	//Zero out the given element
	r[i] = zeroVal
	r = HeapifyUp(r, i, lt)
	return r
}
