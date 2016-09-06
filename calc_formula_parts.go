package calc

import "strconv"

// FormulaParts blah
type FormulaParts struct {
	Operator    string
	Expression1 float64
	Expression2 float64
}

// GetFormula recombines the parse
func (f FormulaParts) GetFormula() string {
	return strconv.FormatFloat(float64(f.Expression1), 'f', -1, 64) + f.Operator + strconv.FormatFloat(float64(f.Expression2), 'f', -1, 64)
}
