package formula

import (
	"math"
	"testing"
)

func TestFormula_Eval(t *testing.T) {
	formula, err := New("(1 * 2 / 3 + 4 - 5 * (21*z+pow(x*3, 3)))")
	if err != nil {
		t.Error(err)
		return
	}
	actual := formula.MustEval(Var("x", 4.5), Var("z", 5))
	expected := 1.0*2.0/3.0 + 4.0 - 5.0*(21.0*5.0+math.Pow(4.5*3.0, 3.0))
	if expected != actual {
		t.Errorf("expected formula result and Go result to be equal: expected: %v, actual: %v", expected, actual)
		return
	}
}

func TestSpecialConstants(t *testing.T) {
	formula, err := New("π + z")
	if err != nil {
		t.Error(err)
		return
	}
	actual := formula.MustEval(Var("z", 5))
	expected := math.Pi + 5
	if expected != actual {
		t.Errorf("expected formula result and Go result to be equal: expected: %v, actual: %v", expected, actual)
		return
	}
}

func TestMinMax(t *testing.T) {
	formula, err := New("min(1,2,3) + max(3,2,1)")
	if err != nil {
		t.Error(err)
		return
	}
	actual := formula.MustEval()
	expected := 4.0
	if expected != actual {
		t.Errorf("expected formula result and Go result to be equal: expected: %v, actual: %v", expected, actual)
		return
	}
}
