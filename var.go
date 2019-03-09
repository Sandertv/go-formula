package formula

import (
	"fmt"
)

// Variable represents a variable with a specific name and value, that may be passed to a formula.
type Variable struct {
	name  string
	value float64
}

// Var returns a new variable that may be passed to a formula when evaluating it. All variables in the formula
// with that name will then adapt the value of the variable. The value passed must be a numeric value. If the
// value is not numeric, the function panics.
func Var(name string, value interface{}) Variable {
	return Variable{name: name, value: valueToFloat64(value)}
}

// valueToFloat converts a numeric value to a float64 value. If the value passed was not numeric, the function
// panics.
func valueToFloat64(value interface{}) float64 {
	switch val := value.(type) {
	case uint8:
		return float64(val)
	case int8:
		return float64(val)
	case uint16:
		return float64(val)
	case int16:
		return float64(val)
	case uint32:
		return float64(val)
	case int32:
		return float64(val)
	case uint64:
		return float64(val)
	case int64:
		return float64(val)
	case int:
		return float64(val)
	case uint:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return val
	default:
		panic(fmt.Sprintf("invalid variable type %T, must be numeric", value))
	}
}

// vars is a map of variables in a name => value map.
type vars map[string]float64
