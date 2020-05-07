package formula

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"reflect"
	"strconv"
)

// astParser handles the parsing of the AST produced by parsing the formula passed into the astParser as an
// AST expression. It 'links' functions and variables when they are encountered.
type astParser struct {
	// formula is the formula that ought to be parsed.
	formula string
	// functions is a map of functions added to the formula which may be executed by the formula. The
	// functions are indexed by their names.
	functions map[string]availableFunc
}

// availableFunc represents a function that was made available to the function to use.
type availableFunc struct {
	// function is the function that is called when the formula calls the function.
	function func(args ...float64) float64
	// paramCount is the minimum parameter count that must be passed to this function. If the amount of
	// parameters passed is lower than paramCount, the function above is not called.
	paramCount int
}

// parse parses the formula in the astParser into a function that may be executed by passing a vars map into
// it. If the parsing was not successful, an error is returned.
func (p *astParser) parse() (eval func(vars vars) float64, err error) {
	expr, err := parser.ParseExpr(p.formula)
	if err != nil {
		return nil, fmt.Errorf("error parsing expression: %v", err)
	}
	return p.parseExpr(expr)
}

// parseExpr parses the expression passed by checking what type it is and applying the correct parser. An
// error is returned if the expression parsed returned one or if the expression was not one of the allowed
// types.
func (p *astParser) parseExpr(e ast.Expr) (eval func(vars vars) float64, err error) {
	switch expr := e.(type) {
	case *ast.BasicLit:
		eval, err = p.parseBasicLit(expr)
	case *ast.Ident:
		eval, err = p.parseIdent(expr)
	case *ast.BinaryExpr:
		eval, err = p.parseBinaryExpr(expr)
	case *ast.ParenExpr:
		eval, err = p.parseParenExpr(expr)
	case *ast.CallExpr:
		eval, err = p.parseCallExpr(expr)
	default:
		return nil, fmt.Errorf("cannot parse unknown expression %v", reflect.TypeOf(e).Elem().String())
	}
	return
}

// parseBinaryExpr parses a binary expression. This is an expression that has an operator in it to add,
// subtract, multiply etc. Each binary expression only has one operator and 2 expressions. The AST package
// splits the formula up correctly itself.
func (p *astParser) parseBinaryExpr(expr *ast.BinaryExpr) (eval func(vars vars) float64, err error) {
	x, err := p.parseExpr(expr.X)
	if err != nil {
		return nil, fmt.Errorf("cannot parse binary expression X: %v", err)
	}
	y, err := p.parseExpr(expr.Y)
	if err != nil {
		return nil, fmt.Errorf("cannot parse binary expression Y: %v", err)
	}

	switch expr.Op {
	case token.ADD:
		eval = func(vars vars) float64 {
			return x(vars) + y(vars)
		}
	case token.SUB:
		eval = func(vars vars) float64 {
			return x(vars) - y(vars)
		}
	case token.MUL:
		eval = func(vars vars) float64 {
			return x(vars) * y(vars)
		}
	case token.QUO:
		eval = func(vars vars) float64 {
			return x(vars) / y(vars)
		}
	case token.REM:
		eval = func(vars vars) float64 {
			return math.Mod(x(vars), y(vars))
		}
	}
	return
}

// parseBasicLit parses a basic literal, provided the literal is a numeric one, like a float or an integer.
// Both integers and floats are parsed as a float64.
func (p *astParser) parseBasicLit(lit *ast.BasicLit) (func(vars vars) float64, error) {
	switch lit.Kind {
	case token.INT:
		val, err := strconv.Atoi(lit.Value)
		if err != nil {
			return nil, fmt.Errorf("invalid value for token.INT %v: %v", lit.Value, err)
		}
		value := float64(val)
		return wrapFunc(value), nil
	case token.FLOAT:
		val, err := strconv.ParseFloat(lit.Value, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for token.FLOAT %v: %v", lit.Value, err)
		}
		return wrapFunc(val), nil
	default:
		return nil, fmt.Errorf("literal must be of type token.INT or token.FLOAT, got %v", lit.Kind)
	}
}

// parseIdent parses an identifier. (generally a variable that needs to be substituted with what is found in
// the vars map passed)
func (p *astParser) parseIdent(ident *ast.Ident) (func(vars vars) float64, error) {
	return func(vars vars) float64 {
		value, ok := vars[ident.Name]
		if !ok {
			return math.NaN()
		}
		return value
	}, nil
}

// parseParenExpr parses an expression within parentheses and returns the function returned by the expression
// within those parentheses.
func (p *astParser) parseParenExpr(expr *ast.ParenExpr) (func(vars vars) float64, error) {
	return p.parseExpr(expr.X)
}

// parseCallExpr parses a call expression. It parses all parameters inside of the function and evaluates them
// when the function is evaluated.
func (p *astParser) parseCallExpr(expr *ast.CallExpr) (func(vars vars) float64, error) {
	args := make([]func(vars vars) float64, len(expr.Args))
	for i, arg := range expr.Args {
		f, err := p.parseExpr(arg)
		if err != nil {
			return nil, fmt.Errorf("error parsing function parameter: %v", err)
		}
		args[i] = f
	}
	return func(vars vars) float64 {
		f, ok := p.functions[expr.Fun.(*ast.Ident).Name]
		if !ok {
			return math.NaN()
		}
		if len(expr.Args) < f.paramCount {
			// Too few arguments supplied to the function.
			return math.NaN()
		}
		argValues := make([]float64, len(expr.Args))
		for i, argValue := range args {
			argValues[i] = argValue(vars)
		}
		return f.function(argValues...)
	}, nil
}

// wrapFunc returns a function that wraps around the value passed and returns it.
func wrapFunc(value float64) func(vars vars) float64 {
	return func(vars vars) float64 {
		return value
	}
}
