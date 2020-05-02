// +build go1.14

package formula

import "math"

// fma ...
func fma(params ...float64) float64 {
	return math.FMA(params[0], params[1], params[2])
}
