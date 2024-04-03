package skiena_3

import (
	"testing"
)

type parenPos struct {
	paren string
	pos   int
}

// Assumes string is only right and left parentheses
func balancedParentheses(ss string) (bool, int, int) {
	maxDistance := 0
	a := 0
	makeParenPos := func(s string) []parenPos {
		r := []parenPos{}
		for i := 0; i < len(s); i++ {
			r = append(r, parenPos{s[i : i+1], i})
		}
		return r
	}
	s := makeParenPos(ss)
	for {
		if a >= len(s) {
			break
		} else if s[0].paren == ")" {
			break
		} else if s[a].paren == "(" {
			a = a + 1
			continue
		} else { //s[a] == ")"
			m := s[a].pos - s[a-1].pos
			if m > maxDistance {
				maxDistance = m
			}
			sa := s[0 : a-1]
			sb := s[a+1 : len(s)]
			s = append(sa, sb...)
			a = 0
			continue
		}

	}
	if len(s) == 0 {
		return true, -1, maxDistance
	} else {
		return false, s[0].pos, maxDistance
	}
}

func TestBalancedParentheses(t *testing.T) {

	actual, startOfError, maxDistance := balancedParentheses("(())")
	if !actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}

	if startOfError != -1 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, -1)
	}

	if maxDistance != 3 {
		t.Errorf("Actual:%v, Expected:%v", maxDistance, 3)
	}

	actual, startOfError, maxDistance = balancedParentheses("((()))")
	if !actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}

	if startOfError != -1 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, -1)
	}
	if maxDistance != 5 {
		t.Errorf("Actual:%v, Expected:%v", maxDistance, 5)
	}

	actual, startOfError, maxDistance = balancedParentheses("((()()))")
	if !actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}
	if startOfError != -1 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, -1)
	}

	if maxDistance != 7 {
		t.Errorf("Actual:%v, Expected:%v", maxDistance, 7)
	}
	actual, startOfError, maxDistance = balancedParentheses("((())))")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 6 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 6)
	}

	actual, startOfError, maxDistance = balancedParentheses("(((()))")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 0 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 0)
	}

	actual, startOfError, maxDistance = balancedParentheses(")())))(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 0 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 0)
	}
	actual, startOfError, maxDistance = balancedParentheses("))((()))(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 0 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 0)
	}
	actual, startOfError, maxDistance = balancedParentheses(")((()))(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 0 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 0)
	}

	actual, startOfError, maxDistance = balancedParentheses(")")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 0 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 0)
	}

	actual, startOfError, maxDistance = balancedParentheses(")(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 0 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 0)
	}

	actual, startOfError, maxDistance = balancedParentheses("(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 0 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 0)
	}

	actual, startOfError, maxDistance = balancedParentheses("")
	if !actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}
	if startOfError != -1 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, -1)
	}

}
