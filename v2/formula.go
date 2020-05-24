package formula

import (
	"golang.org/x/xerrors"
	"math"
)

// Formula is a parsed formula that is ready to be evaluated. It is safe to use concurrently from multiple
// goroutines.
type Formula struct {
	parser *astParser
	// evaluate is the function called when the formula is evaluated.
	evaluate func(vars vars) (float64, error)
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
		return nil, xerrors.Errorf("error parsing formula: %w", err)
	}
	f := &Formula{evaluate: eval, parser: p}
	f.registerDefaults()
	return f, nil
}

// RegisterFunc registers a custom function to be usable by the formula. This function allows an arbitrary number of input
// floats and one output float. The paramCount passed indicates the number of input floats expected. If less than the
// required paramCount arguments are passed to the function, Eval will return an ErrInsufficientArgs error. The function
// does not need to internally check the correct arg length. Functions must be registered with the formula before evaluating.
//
// Example:
//
//  // Add sinc function: https://en.wikipedia.org/wiki/Sinc_function
//  RegisterFunc("sinc", 1, func(args ...float64) float64 {
//     if args[0] == 0 {
//        return 1
//     }
//     return math.Sin(args[0]) / args[0]
//  })
//
func (formula *Formula) RegisterFunc(name string, paramCount int, f func(args ...float64) float64) {
	formula.parser.functions[name] = availableFunc{function: f, paramCount: paramCount}
}

// Eval evaluates a formula using the variables passed. If an unknown variable/constant or function is encountered,
// ErrUnknownVariable or ErrUnknownFunc is returned respectively. If a known function is passed with too few arguments,
// ErrInsufficientArgs is returned.
//
// Some special math constants are already included. They are automatically defined unless over-ridden
// by variables. These are: π, 𝜋, pi, Φ, phi, e, E.
func (formula *Formula) Eval(variables ...Variable) (float64, error) {
	// Add special constants
	variableMap := vars{
		"π":  math.Pi,
		"𝜋":  math.Pi,
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

// MustEval calls Eval but panics if Eval returns an error.
func (formula *Formula) MustEval(variables ...Variable) float64 {
	f, err := formula.Eval(variables...)
	if err != nil {
		panic(err)
	}
	return f
}

// registerDefaults registers all functions found in the functions.go file to the formula. This is done for
// each formula automatically, so these functions do not need to be added manually.
func (formula *Formula) registerDefaults() {
	formula.RegisterFunc("abs", 1, abs)
	formula.RegisterFunc("acos", 1, acos)
	formula.RegisterFunc("acosh", 1, acosh)
	formula.RegisterFunc("asin", 1, asin)
	formula.RegisterFunc("asinh", 1, asinh)
	formula.RegisterFunc("atan", 1, atan)
	formula.RegisterFunc("atan2", 2, atan2)
	formula.RegisterFunc("atanh", 1, atanh)
	formula.RegisterFunc("cbrt", 1, cbrt)
	formula.RegisterFunc("ceil", 1, ceil)
	formula.RegisterFunc("copysign", 2, copysign)
	formula.RegisterFunc("cos", 1, cos)
	formula.RegisterFunc("cosh", 1, cosh)
	formula.RegisterFunc("dim", 2, dim)
	formula.RegisterFunc("erf", 1, erf)
	formula.RegisterFunc("erfc", 1, erfc)
	formula.RegisterFunc("erfcinv", 1, erfcinv)
	formula.RegisterFunc("erfinv", 1, erfinv)
	formula.RegisterFunc("exp", 1, exp)
	formula.RegisterFunc("exp2", 1, exp2)
	formula.RegisterFunc("expm1", 1, expm1)
	formula.RegisterFunc("floor", 1, floor)
	formula.RegisterFunc("gamma", 1, gamma)
	formula.RegisterFunc("hypot", 2, hypot)
	formula.RegisterFunc("j0", 1, j0)
	formula.RegisterFunc("j1", 1, j1)
	formula.RegisterFunc("jn", 2, jn)
	formula.RegisterFunc("log", 1, log)
	formula.RegisterFunc("log10", 1, log10)
	formula.RegisterFunc("log1p", 1, log1p)
	formula.RegisterFunc("log2", 1, log2)
	formula.RegisterFunc("logb", 1, logb)
	formula.RegisterFunc("max", 1, max)
	formula.RegisterFunc("min", 1, min)
	formula.RegisterFunc("mod", 2, mod)
	formula.RegisterFunc("nextafter", 2, nextafter)
	formula.RegisterFunc("pow", 2, pow)
	formula.RegisterFunc("pow10", 1, pow10)
	formula.RegisterFunc("remainder", 2, remainder)
	formula.RegisterFunc("round", 1, round)
	formula.RegisterFunc("roundtoeven", 1, roundtoeven)
	formula.RegisterFunc("sin", 1, sin)
	formula.RegisterFunc("sinh", 1, sinh)
	formula.RegisterFunc("sqrt", 1, sqrt)
	formula.RegisterFunc("tan", 1, tan)
	formula.RegisterFunc("tanh", 1, tanh)
	formula.RegisterFunc("trunc", 1, trunc)
	formula.RegisterFunc("y0", 1, y0)
	formula.RegisterFunc("y1", 1, y1)
	formula.RegisterFunc("yn", 1, yn)

	formula.registerExtra()
}
