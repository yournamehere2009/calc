package calc

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

//ComputeFormula blah
func ComputeFormula(formula string) (float64, error) {
	var decomposedFormula string
	var err error
	var f *FormulaParts

	formula = convertShortHandMultiplication(formula)
	formula = convertDoubleNegativeToAddition(formula)

	if decomposedFormula, err = decompose(formula); err != nil {
		return 0, err
	}

	if f, err = ParseFormula(decomposedFormula); err != nil {
		return 0, err
	}

	return compute(f)
}

func decompose(formula string) (string, error) {
	var result float64
	var err error
	var f *FormulaParts

	// If we still have parenthesis in the formula
	if strings.Index(formula, "(") != -1 {
		// Find an opening followed by a closing with no other parenthesis in-between

		//Starting parenthesis
		iOpening := 0
		iClosing := 0

		// Loop over every character
		for i, r := range formula {
			if string(r) == "(" {
				iOpening = i
			} else if string(r) == ")" {
				iClosing = i
				break
			}
		}

		contents := formula[iOpening+1 : iClosing]

		if f, err = ParseFormula(contents); err != nil {
			return "", err
		} else if result, err = compute(f); err != nil {
			return "", err
		}

		formula = strings.Replace(formula, formula[iOpening:iClosing+1], strconv.FormatFloat(float64(result), 'f', -1, 64), -1)

		if formula, err = decompose(formula); err != nil {
			return "", err
		}
	}
	return formula, err
}

func compute(fp *FormulaParts) (float64, error) {
	switch fp.Operator {
	case "+":
		return Add(fp.Expression1, fp.Expression2), nil
	case "-":
		return Subtract(fp.Expression1, fp.Expression2), nil
	case "*":
		return Multiply(fp.Expression1, fp.Expression2), nil
	case "/":
		return Divide(fp.Expression1, fp.Expression2)
	default:
		return float64(0), errors.New("Unrecognized operator")
	}
}

func convertShortHandMultiplication(formula string) string {
	if match, _ := regexp.MatchString("\\d\\(", formula); match {
		r, _ := regexp.Compile("\\d\\(")

		i := r.FindStringIndex(formula)

		formula = formula[0:i[0]+1] + "*" + formula[i[1]-1:len(formula)]
	}

	return strings.TrimSpace(formula)
}

func convertDoubleNegativeToAddition(formula string) string {
	negativeIndex := 0

	for negativeIndex > -1 {
		negativeIndex = strings.Index(formula, "--")

		if negativeIndex != -1 {
			formula = strings.Replace(formula, formula[negativeIndex:negativeIndex+2], "+", -1)
		}
	}

	return formula
}
