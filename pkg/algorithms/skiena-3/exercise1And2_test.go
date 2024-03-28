package skiena_3

import (
	"fmt"
	"strings"
	"testing"
)

// Assumes string is only right and left parentheses
func balancedParentheses(s string) bool {
	var a int
	for {
		if a >= len(s) {
			break
		} else if strings.HasPrefix(s, ")") {
			break
		} else if s[a:a+1] == "(" {
			a = a + 1
			continue
		} else { //s[a] == ")"
			s = fmt.Sprintf("%v%v", s[0:a-1], s[a+1:len(s)])
			a = 0
			continue
		}
	}
	return len(s) == 0
}

func TestBalancedParentheses(t *testing.T) {

	actual := balancedParentheses("(())")
	if !actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}

	actual = balancedParentheses("((()))")
	if !actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}

	actual = balancedParentheses("((()()))")
	if !actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}

	actual = balancedParentheses("((())))")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}

	actual = balancedParentheses("(((()))")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}

	actual = balancedParentheses(")())))(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}
	actual = balancedParentheses("))((()))(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}
	actual = balancedParentheses(")((()))(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}

	actual = balancedParentheses(")")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}

	actual = balancedParentheses(")(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}

	actual = balancedParentheses("(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}
}
