package chapter2

import "github.com/greymatter-io/golangz/arrays"

//Big O for this algorithm is O(2 times n to (N-1)th power)
func sum(source []int) [][]int64 {
	var result = make([][]int64, len(source))
	for i, j := range source {
		result[i] = make([]int64, 2)
		result[i][0] = int64(j)
		var sum int64
		for k := 0; k <= i; k++ {
			sum = sum + int64(source[k])
		}
		result[i][1] = sum
	}
	return result
}

//This algorithm has a constant running time of O(n)
func matrixSumWitoutInnerLoop(source []int) [][]int64 {
	var result = make([][]int64, len(source))
	for i, j := range source {
		result[i] = make([]int64, 2)
		result[i][0] = int64(j)
		if i == 0 { //Set first sum to first element value
			result[i][1] = result[i][0]
		} else { //Just grqab previous sum
			result[i][1] = result[i-1][1] + result[i][0]
		}
	}
	return result
}

//This algorithm has a constant running time of O(n) but is a bit slower(maybe 20%) than it's loop counterpart above, but much faster than the first sum above.
func matrixSumWithFoldRight(source []int) [][]int64 {
	var inner [][]int64
	var append = func(x int, xs [][]int64) [][]int64 {
		var result = make([]int64, 2)
		currentX := int64(x)
		result[0] = currentX
		currentAccumLen := len(xs)
		if currentAccumLen == 0 { //Set first sum to first element value
			result[1] = currentX
		} else { //Just grab previous sum
			result[1] = xs[currentAccumLen-1][1] + currentX
		}
		xs = append(xs, result)
		return xs
	}
	result := arrays.FoldRight(source, inner, append)
	return result
}
