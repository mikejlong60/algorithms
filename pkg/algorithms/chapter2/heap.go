package chapter2

//Key for now is the array index, starting with zero.
func ParentIdx(i int) int {
	//Odd number
	if i%2 > 0 {
		return i / 2
	} else {
		return (i / 2) - 1
	}
}

//This is a pure function.  But it's slower than the impure version because I copy the heap array every time.
func HeapifyUp(heap []int, i int) []int {
	if len(heap) == 0 {
		return []int{}
	}
	r := make([]int, len(heap))
	copy(r, heap)

	if i > 1 {
		j := ParentIdx(i)
		if r[i] < r[j] {
			//Swap elements
			temp := r[i]
			temp2 := r[j]
			r[j] = temp
			r[i] = temp2
			HeapifyUp(r, j)
		}
	}
	return r
}

func HeapifyDown(heap []int, i int) []int {
	if len(heap) == 0 {
		return []int{}
	}
	r := make([]int, len(heap))
	copy(r, heap)

	var j int
	n := len(heap)
	if 2*i > n {
		return r
	} else if 2*i < n {
		leftIdx := 2 * i
		rightIdx := 2*i + 1
		leftVal := r[leftIdx]
		rightVal := r[rightIdx]
		if leftVal < rightVal {
			j = leftIdx
		} else {
			j = rightIdx
		}
	} else if 2*i == n {
		j = 2 * i
	}
	val1 := r[j]
	val2 := r[i]
	if val1 < val2 {
		//Swap elements
		temp := r[i]
		temp2 := r[j]
		r[j] = temp
		r[i] = temp2
		HeapifyDown(r, j)
	}
	return r
}
