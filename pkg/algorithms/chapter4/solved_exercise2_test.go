package chapter4

import (
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

var eq = func(l, r DayStockPrice) bool {
	if l.day == r.day {
		return true
	} else {
		return false
	}
}

func TestMostProfitAscending3(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 10,
	}
	b := DayStockPrice{
		day:   1,
		price: 11,
	}

	c := DayStockPrice{
		day:   2,
		price: 12,
	}

	actual := MostProfit([]DayStockPrice{a, b, c})
	expected := []DayStockPrice{a, c}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestMostProfitAscending6(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 10,
	}
	b := DayStockPrice{
		day:   1,
		price: 11,
	}

	c := DayStockPrice{
		day:   2,
		price: 12,
	}
	d := DayStockPrice{
		day:   3,
		price: 13,
	}
	e := DayStockPrice{
		day:   4,
		price: 14,
	}

	f := DayStockPrice{
		day:   5,
		price: 15,
	}

	actual := MostProfit([]DayStockPrice{a, b, c, d, e, f})
	expected := []DayStockPrice{a, f}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestMostProfitAscending8(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 10,
	}
	b := DayStockPrice{
		day:   1,
		price: 11,
	}

	c := DayStockPrice{
		day:   2,
		price: 9,
	}
	d := DayStockPrice{
		day:   3,
		price: 5,
	}
	e := DayStockPrice{
		day:   4,
		price: 6,
	}

	f := DayStockPrice{
		day:   5,
		price: 14,
	}

	g := DayStockPrice{
		day:   6,
		price: 15,
	}
	h := DayStockPrice{
		day:   7,
		price: 17,
	}
	actual := MostProfit([]DayStockPrice{a, b, c, d, e, f, g, h})
	expected := []DayStockPrice{e, h}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestMostProfitMiddle6(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 10,
	}
	b := DayStockPrice{
		day:   1,
		price: 11,
	}

	c := DayStockPrice{
		day:   2,
		price: 12,
	}
	d := DayStockPrice{
		day:   3,
		price: 130,
	}
	e := DayStockPrice{
		day:   4,
		price: 14,
	}

	f := DayStockPrice{
		day:   5,
		price: 15,
	}

	actual := MostProfit([]DayStockPrice{a, b, c, d, e, f})
	expected := []DayStockPrice{a, d}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestMostProfitMiddle5(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 10,
	}
	b := DayStockPrice{
		day:   1,
		price: 11,
	}

	c := DayStockPrice{
		day:   2,
		price: 1200,
	}
	d := DayStockPrice{
		day:   3,
		price: 130,
	}
	e := DayStockPrice{
		day:   4,
		price: 14,
	}

	actual := MostProfit([]DayStockPrice{a, b, c, d, e})
	expected := []DayStockPrice{a, c}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestNoProfitDescending5(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 10,
	}
	b := DayStockPrice{
		day:   1,
		price: 9,
	}

	c := DayStockPrice{
		day:   2,
		price: 8,
	}
	d := DayStockPrice{
		day:   3,
		price: 7,
	}
	e := DayStockPrice{
		day:   4,
		price: 6,
	}

	actual := MostProfit([]DayStockPrice{a, b, c, d, e})
	expected := []DayStockPrice{}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}
