package formula

import (
	"fmt"
)

// ErrPanic is returned when a registered function panics.
type ErrPanic struct {
	// Func is the name of the function.
	Func string
	// Pos is the character position of Func.
	Pos int
	// Reason for panic.
	Reason string
	// Filename of where panic occurred.
	File string
	// Line number of where panic occurred.
	Line int
}

// Error implements error.
func (e *ErrPanic) Error() string {
	return fmt.Sprintf("panic func: %s (pos:%d): %s [%s@L%d]", e.Func, e.Pos, e.Reason, e.File, e.Line)
}

// ErrInsufficientArgs is returned when a function in a formula requires more arguments than that provided.
type ErrInsufficientArgs struct {
	// Func is the name of the function.
	Func string
	// Pos is the character position of Func.
	Pos int
	// Actual is the number of arguments provided to Func.
	Actual int
	// Expected is the minimum number of arguments expected by Func.
	Expected int
}

// Error implements error.
func (e *ErrInsufficientArgs) Error() string {
	return fmt.Sprintf("insufficient args: %s (pos:%d)", e.Func, e.Pos)
}

// ErrUnknownFunc is returned when a formula contains an unrecognized function.
type ErrUnknownFunc struct {
	// Func is the name of the unknown function encountered.
	Func string
	// Pos is the character position of the unknown Func.
	Pos int
}

// Error implements error.
func (e *ErrUnknownFunc) Error() string {
	return fmt.Sprintf("unknown func: %s (pos:%d)", e.Func, e.Pos)
}

// ErrUnknownVariable is returned when a formula contains an unrecognized variable or constant.
type ErrUnknownVariable struct {
	// Var is the name of the unknown variable or constant.
	Var string
	// Pos is the character position of the unknown variable or constant.
	Pos int
}

// Error implements error.
func (e *ErrUnknownVariable) Error() string {
	return fmt.Sprintf("unknown var: %s (pos:%d)", e.Var, e.Pos)
}
