package calc

import (
	"regexp"
	"strconv"
	"strings"
)

//ParseFormula blah
func ParseFormula(formula string) (*FormulaParts, error) {
	f := new(FormulaParts)
	oIndex := -1
	operators := [5]string{"^", "*", "/", "+", "-"}
	reg, _ := regexp.Compile("\\d-|\\+|/|\\*|\\^")

	// First get the operator
	for i := 0; i < len(operators); i++ {
		oIndex = strings.Index(formula, operators[i])

		if oIndex > 0 {
			f.Operator = formula[oIndex : oIndex+1]
			break
		}
	}

	if oIndex <= 0 {
		f.Operator = "+"

		e1, _ := strconv.ParseFloat(strings.TrimSpace(formula), 64)

		f.Expression1 = e1
		f.Expression2 = float64(0)

		return f, nil
	}

	// Next, get the first expression
	e1, _ := strconv.ParseFloat(getNumberBeforeOperator(strings.TrimSpace(formula[:oIndex])), 64)
	f.Expression1 = e1

	// Finally, get the second expression
	e2, _ := strconv.ParseFloat(getNumberAfterOperator(strings.TrimSpace(formula[oIndex+1:len(formula)])), 64)
	f.Expression2 = e2

	// I need to know if there are multiple operators in the formula. If so, I need to parse and compute based on order of operations
	if matchIndexes := reg.FindAllString(formula, -1); len(matchIndexes) > 1 {
		result, _ := compute(f)
		formula = strings.Replace(formula, f.GetFormula(), strconv.FormatFloat(float64(result), 'f', -1, 64), -1)
		f, _ = ParseFormula(formula)
	}

	return f, nil
}

func getNumberBeforeOperator(formulaBeforeOperator string) string {
	// Chinn: this needs to change from taking the whole first part, to only taking the numbers in front of the operator (and the negative sign if present)
	var nIndex int
	var potentialNegative bool

	for i := len(formulaBeforeOperator) - 1; i > -1; i-- {
		c := string(formulaBeforeOperator[i])

		if match, _ := regexp.MatchString("\\+|\\*|\\/|\\^", c); match {
			// If it's an operator other than minus (or negative), we've reached the start of the number.
			nIndex = i
			break
		} else if match, _ := regexp.MatchString("\\d", c); !match {
			//If it's not a number or an operator it could be a decimal, or a negative sign.

			// If it's a decimal, keep going
			if c == "." {
				continue
			} else if c == "-" {
				nIndex = i
				potentialNegative = true
			}
		} else if potentialNegative {
			// It's a number, then it wasn't the negative sign and was instead the minus sign
			nIndex = i + 1
		}
	}

	if nIndex == 0 {
		return formulaBeforeOperator
	}

	return formulaBeforeOperator[nIndex+1:]
}

func getNumberAfterOperator(formulaAfterOperator string) string {
	// Chinn: this needs to change from taking the whole first part, to only taking the numbers in front of the operator (and the negative sign if present)
	var nIndex int

	for i := 0; i < len(formulaAfterOperator); i++ {
		c := string(formulaAfterOperator[i])

		if match, _ := regexp.MatchString("\\+|\\*|\\/|\\^", c); match {
			// If it's an operator other than minus (or negative), we've reached the start of the number.
			nIndex = i
			break
		} else if match, _ := regexp.MatchString("\\d", c); !match {
			//If it's not a number or an operator it could be a decimal, or a negative sign.

			// If it's a decimal, keep going
			if c == "." || (c == "-" && i == 0) {
				continue
			} else if c == "-" {
				nIndex = i
				break
			}
		}
	}

	if nIndex == 0 {
		return formulaAfterOperator
	}

	return formulaAfterOperator[:nIndex]
}
