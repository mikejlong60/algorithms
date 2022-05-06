package chapter2

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
func betterMatrixSum(source []int) [][]int64 {
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
