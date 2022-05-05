package chapter2

//Big O for this algorithm is O(2(N-1))
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
