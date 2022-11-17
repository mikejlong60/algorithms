package chapter4

var numberOfSteps int

func peakOfAList(xs []int) int {
	numberOfSteps = numberOfSteps + 1
	if len(xs) == 3 {
		if xs[2] > xs[0] && xs[2] > xs[1] { //peak is last
			return xs[2]
		} else if xs[1] > xs[0] && xs[1] > xs[2] { //peak is second
			return xs[1]
		} else if xs[0] > xs[1] && xs[0] > xs[2] { //peak is first
			return xs[0]
		} else { //they are all equal so just return first
			return xs[0]
		}
	} else if len(xs) == 2 {
		if xs[0] > xs[1] {
			return xs[0]
		} else {
			return xs[1]
		}
	} else {
		//slice the array in half and send it off recursively
		i := len(xs) / 2
		return peakOfAList([]int{peakOfAList(xs[0:i]), peakOfAList(xs[i:])})
	}
}
