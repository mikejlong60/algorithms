package chapter2

func findBreakingPointWithoutBreakingJar(ladder []int, breakingPoint int) int {
	breakingPointIsHigher := ladder[0] <= breakingPoint
	var wrungB4BreakingPoint = -1

	if breakingPointIsHigher {
		for i := 0; i < len(ladder); i++ {
			if ladder[i] >= breakingPoint {
				wrungB4BreakingPoint = ladder[i-1] //the wrung of the ladder right before the breaking point
				break
			}
		}
	} else { //breaking point is lower
		for i := len(ladder) - 1; i >= 0; i-- {
			if ladder[i] >= breakingPoint {
				wrungB4BreakingPoint = ladder[i-1] //the wrung of the ladder right before the breaking point
				break
			}
		}
	}
	return wrungB4BreakingPoint
}

//asymptotic lower bound is (n Log n)
//theta is (n + (n log n))
//asymptotic upper bound is (n)
//breaking point is assumed to be an ordered set of integers from low to high
func HighestBreakingPoint(ladder []int, breakingPoint, budget, usedBudget int) int {
	if usedBudget+1 == budget {
		return findBreakingPointWithoutBreakingJar(ladder, breakingPoint)
	} else { //Divide in half again
		lowerHalf := ladder[0 : len(ladder)/2]
		upperHalf := ladder[len(ladder)/2:]
		if breakingPoint <= upperHalf[0] {
			return HighestBreakingPoint(lowerHalf, breakingPoint, budget, usedBudget+1)
		} else {
			return HighestBreakingPoint(upperHalf, breakingPoint, budget, usedBudget+1)
		}
	}
}
