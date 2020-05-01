package formula

import (
	"fmt"
	"math"
)

// Formula is a parsed formula that is ready to be evaluated. A formula may be re-used an unlimited amount of
// times and is safe to be used from different goroutines concurrently.
type Formula struct {
	parser *astParser
	// evaluate is the function called when the formula is evaluated.
	evaluate func(vars vars) float64
}

// New returns a new formula for a given string. The formula is parsed and may be evaluated if parsed
// successfully. If not successful, an error is returned and the formula is nil.
//
// Most functions from the math package that return a single float64 are supported. The equivalent function
// name is all lower-cased. Therefore RoundToEven becomes roundtoeven. See https://golang.org/pkg/math/.
func New(formula string) (*Formula, error) {
	p := &astParser{formula: formula, functions: make(map[string]availableFunc)}
	eval, err := p.parse()
	if err != nil {
		return nil, fmt.Errorf("error parsing formula: %v", err)
	}
	f := &Formula{evaluate: eval, parser: p}
	f.registerDefaults()
	return f, nil
}

// Func adds a function to be usable by the formula. This function allows any amount of input floats and one
// output float. The paramCount passed indicates the amount of input floats are expected. If less than the
// required paramCount arguments are passed to the function, 0 is returned, so that the function passed does
// not need to check for the correct arg length.
// Functions must be added to the formula before evaluating, and need only to be added once.
func (formula *Formula) Func(name string, paramCount int, f func(args ...float64) float64) {
	formula.parser.functions[name] = availableFunc{function: f, paramCount: paramCount}
}

// Eval evaluates a formula using the variables passed. Any variable in the formula parsed that is not passed
// to this method is considered zero.
//
// Some special math constants are automatically included. They are automatically defined unless over-ridden
// by variables. These are: π, pi, Φ, phi, e, E.
func (formula *Formula) Eval(variables ...Variable) float64 {
	// Add special constants
	variableMap := vars{
		"π":  math.Pi,
		"pi": math.Pi,

		"Φ":   math.Phi,
		"phi": math.Phi,

		"e": math.E,
		"E": math.E,

		"nan": math.NaN(),
	}

	for _, variable := range variables {
		variableMap[variable.name] = variable.value
	}
	return formula.evaluate(variableMap)
}

// registerDefaults registers all functions found in the functions.go file to the formula. This is done for
// each formula automatically, so these functions do not need to be added manually.
func (formula *Formula) registerDefaults() {
	formula.Func("abs", 1, abs)
	formula.Func("acos", 1, acos)
	formula.Func("acosh", 1, acosh)
	formula.Func("asin", 1, asin)
	formula.Func("asinh", 1, asinh)
	formula.Func("atan", 1, atan)
	formula.Func("atan2", 2, atan2)
	formula.Func("atanh", 1, atanh)
	formula.Func("cbrt", 1, cbrt)
	formula.Func("ceil", 1, ceil)
	formula.Func("copysign", 2, copysign)
	formula.Func("cos", 1, cos)
	formula.Func("cosh", 1, cosh)
	formula.Func("dim", 2, dim)
	formula.Func("erf", 1, erf)
	formula.Func("erfc", 1, erfc)
	formula.Func("erfcinv", 1, erfcinv)
	formula.Func("erfinv", 1, erfinv)
	formula.Func("exp", 1, exp)
	formula.Func("exp2", 1, exp2)
	formula.Func("expm1", 1, expm1)
	formula.Func("floor", 1, floor)
	formula.Func("gamma", 1, gamma)
	formula.Func("hypot", 2, hypot)
	formula.Func("j0", 1, j0)
	formula.Func("j1", 1, j1)
	formula.Func("jn", 2, jn)
	formula.Func("log", 1, log)
	formula.Func("log10", 1, log10)
	formula.Func("log1p", 1, log1p)
	formula.Func("log2", 1, log2)
	formula.Func("logb", 1, logb)
	formula.Func("max", 1, max)
	formula.Func("min", 1, min)
	formula.Func("mod", 2, mod)
	formula.Func("nextafter", 2, nextafter)
	formula.Func("pow", 2, pow)
	formula.Func("pow10", 1, pow10)
	formula.Func("remainder", 2, remainder)
	formula.Func("round", 1, round)
	formula.Func("roundtoeven", 1, roundtoeven)
	formula.Func("sin", 1, sin)
	formula.Func("sinh", 1, sinh)
	formula.Func("sqrt", 1, sqrt)
	formula.Func("tan", 1, tan)
	formula.Func("tanh", 1, tanh)
	formula.Func("trunc", 1, trunc)
	formula.Func("y0", 1, y0)
	formula.Func("y1", 1, y1)
	formula.Func("yn", 1, yn)

	formula.registerExtra()
}
