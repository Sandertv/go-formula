// +build go1.14

package formula

func (formula *Formula) registerExtra() {
	formula.Func("fma", 3, fma)
}
