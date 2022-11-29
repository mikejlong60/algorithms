package chapter4

import "math"

type DayStockPrice struct {
	day   int
	price int
}

var maxS = func(xs []DayStockPrice) DayStockPrice {
	var currMax = DayStockPrice{
		day:   math.MinInt,
		price: math.MinInt,
	}
	for _, y := range xs {
		if y.price >= currMax.price {
			currMax = y
		}
	}
	return currMax
}

var minS = func(xs []DayStockPrice) DayStockPrice {
	var currMin = DayStockPrice{
		day:   math.MaxInt,
		price: math.MaxInt,
	}
	for _, y := range xs {
		if y.price <= currMin.price {
			currMin = y
		}
	}
	return currMin
}

var maxMinS = func(max, min DayStockPrice) []DayStockPrice {
	if max.price > math.MinInt && min.price > math.MinInt && max.price > min.price { //&& max.day > min.day {
		return []DayStockPrice{min, max}
	} else {
		return []DayStockPrice{}
	}
}

func MostProfit(xs []DayStockPrice) []DayStockPrice {
	if len(xs) == 0 {
		return xs
	} else if len(xs) == 1 {
		return []DayStockPrice{}
	} else if len(xs) == 2 {
		return maxMinS(maxS(xs), minS(xs))
	} else if len(xs) == 3 { //return highest positive price difference over time or empty
		return maxMinS(maxS(xs), minS(xs))
	} else if len(xs) == 4 { //return highest positive price difference over time or empty
		return maxMinS(maxS(xs), minS(xs))
	} else {
		//slice the array in half and send it off recursively
		i := len(xs) / 2
		left := xs[0:i]
		right := xs[i:]

		a := MostProfit(left)
		b := MostProfit(right)
		return MostProfit(append(a, b...))
	}
}
