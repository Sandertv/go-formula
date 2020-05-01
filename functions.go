package formula

import "math"

// The following functions are automatically exposed to any formula.

// abs ...
func abs(params ...float64) float64 {
	return math.Abs(params[0])
}

// acos ...
func acos(params ...float64) float64 {
	return math.Acos(params[0])
}

// acosh ...
func acosh(params ...float64) float64 {
	return math.Acosh(params[0])
}

// asin ...
func asin(params ...float64) float64 {
	return math.Asin(params[0])
}

// asinh ...
func asinh(params ...float64) float64 {
	return math.Asinh(params[0])
}

// atan ...
func atan(params ...float64) float64 {
	return math.Atan(params[0])
}

// atan2 ...
func atan2(params ...float64) float64 {
	return math.Atan2(params[0], params[1])
}

// atanh ...
func atanh(params ...float64) float64 {
	return math.Atanh(params[0])
}

// cbrt ...
func cbrt(params ...float64) float64 {
	return math.Cbrt(params[0])
}

// ceil ...
func ceil(params ...float64) float64 {
	return math.Ceil(params[0])
}

// copysign ...
func copysign(params ...float64) float64 {
	return math.Copysign(params[0], params[1])
}

// cos ...
func cos(params ...float64) float64 {
	return math.Cos(params[0])
}

// cosh ...
func cosh(params ...float64) float64 {
	return math.Cosh(params[0])
}

// dim ...
func dim(params ...float64) float64 {
	return math.Dim(params[0], params[1])
}

// erf ...
func erf(params ...float64) float64 {
	return math.Erf(params[0])
}

// erfc ...
func erfc(params ...float64) float64 {
	return math.Erfc(params[0])
}

// erfcinv ...
func erfcinv(params ...float64) float64 {
	return math.Erfcinv(params[0])
}

// erfinv ...
func erfinv(params ...float64) float64 {
	return math.Erfinv(params[0])
}

// exp ...
func exp(params ...float64) float64 {
	return math.Exp(params[0])
}

// exp2 ...
func exp2(params ...float64) float64 {
	return math.Exp2(params[0])
}

// expm1 ...
func expm1(params ...float64) float64 {
	return math.Expm1(params[0])
}

// floor ...
func floor(params ...float64) float64 {
	return math.Floor(params[0])
}

// gamma ...
func gamma(params ...float64) float64 {
	return math.Gamma(params[0])
}

// hypot ...
func hypot(params ...float64) float64 {
	return math.Hypot(params[0], params[1])
}

// j0 ...
func j0(params ...float64) float64 {
	return math.J0(params[0])
}

// j1 ...
func j1(params ...float64) float64 {
	return math.J1(params[0])
}

// jn ...
func jn(params ...float64) float64 {
	return math.Jn(int(params[0]), params[1])
}

// log ...
func log(params ...float64) float64 {
	return math.Log(params[0])
}

// log10 ...
func log10(params ...float64) float64 {
	return math.Log10(params[0])
}

// log1p ...
func log1p(params ...float64) float64 {
	return math.Log1p(params[0])
}

// log2 ...
func log2(params ...float64) float64 {
	return math.Log2(params[0])
}

// logb ...
func logb(params ...float64) float64 {
	return math.Logb(params[0])
}

// max ...
func max(params ...float64) float64 {
	// Assume at least 2 args
	var maxVal float64
	for i := range params {
		if i == 0 {
			maxVal = params[i]
		} else {
			maxVal = math.Max(params[i], maxVal)
		}
	}
	return maxVal
}

// min ...
func min(params ...float64) float64 {
	// Assume at least 2 args
	var minVal float64
	for i := range params {
		if i == 0 {
			minVal = params[i]
		} else {
			minVal = math.Min(params[i], minVal)
		}
	}
	return minVal
}

// mod ...
func mod(params ...float64) float64 {
	return math.Mod(params[0], params[1])
}

// nextafter ...
func nextafter(params ...float64) float64 {
	return math.Nextafter(params[0], params[1])
}

// pow ...
func pow(params ...float64) float64 {
	return math.Pow(params[0], params[1])
}

// pow10 ...
func pow10(params ...float64) float64 {
	return math.Pow10(int(params[0]))
}

// remainder ...
func remainder(params ...float64) float64 {
	return math.Remainder(params[0], params[1])
}

// round ...
func round(params ...float64) float64 {
	return math.Round(params[0])
}

// roundtoeven ...
func roundtoeven(params ...float64) float64 {
	return math.RoundToEven(params[0])
}

// sin ...
func sin(params ...float64) float64 {
	return math.Sin(params[0])
}

// sinh ...
func sinh(params ...float64) float64 {
	return math.Sinh(params[0])
}

// sqrt ...
func sqrt(params ...float64) float64 {
	return math.Sqrt(params[0])
}

// tan ...
func tan(params ...float64) float64 {
	return math.Tan(params[0])
}

// tanh ...
func tanh(params ...float64) float64 {
	return math.Tanh(params[0])
}

// trunc ...
func trunc(params ...float64) float64 {
	return math.Trunc(params[0])
}

// y0 ...
func y0(params ...float64) float64 {
	return math.Y0(params[0])
}

// y1 ...
func y1(params ...float64) float64 {
	return math.Y1(params[0])
}

// yn ...
func yn(params ...float64) float64 {
	return math.Yn(int(params[0]), params[1])
}
