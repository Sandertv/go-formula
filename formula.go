package formula

import "fmt"

// Formula is a parsed formula that is ready to be evaluated. A formula may be re-used an unlimited amount of
// times and is safe to be used from different goroutines concurrently.
type Formula struct {
	parser *astParser
	// evaluate is the function called when the formula is evaluated.
	evaluate func(vars vars) float64
}

// New returns a new formula for a given string. The formula is parsed and may be evaluated if parsed
// successfully. If not successful, an error is returned and the formula is nil.
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
func (formula *Formula) Eval(variables ...Variable) float64 {
	variableMap := vars{}
	for _, variable := range variables {
		variableMap[variable.name] = variable.value
	}
	return formula.evaluate(variableMap)
}

// registerDefaults registers all functions found in the functions.go file to the formula. This is done for
// each formula automatically, so these functions do not need to be added manually.
func (formula *Formula) registerDefaults() {
	formula.Func("tan", 1, tan)
	formula.Func("cos", 1, cos)
	formula.Func("sin", 1, sin)
	formula.Func("pow", 2, pow)
	formula.Func("log", 1, log)
	formula.Func("sqrt", 1, sqrt)
	formula.Func("abs", 1, abs)
	formula.Func("max", 2, max)
	formula.Func("min", 2, min)
	formula.Func("ceil", 1, ceil)
	formula.Func("floor", 1, floor)
	formula.Func("round", 1, round)
}
