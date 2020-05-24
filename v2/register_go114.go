// +build go1.14

package formula

func (formula *Formula) registerExtra() {
	formula.RegisterFunc("fma", 3, fma)
}
