package calc

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//ComputeFormula blah
func ComputeFormula(formula string) (float64, []string, error) {
	defer ClearSteps()

	AddStep(formula)

	var decomposedFormula string
	var err error
	var result float64

	//Clean the formula
	if tmpFormula := removeSpaces(formula); tmpFormula != formula {
		formula = tmpFormula
		AddStep(formula)
	}

	if tmpFormula := convertShortHandMultiplication(formula); tmpFormula != formula {
		formula = tmpFormula
		AddStep(formula)
	}

	if tmpFormula := convertDoubleNegativeToAddition(formula); tmpFormula != formula {
		formula = tmpFormula
		AddStep(formula)
	}

	// Process the formula
	if decomposedFormula, err = decompose(formula); err != nil {
		return 0, GetSteps(), err
	}

	// Output the results
	result, _ = strconv.ParseFloat(decomposedFormula, 64)

	return result, GetSteps(), err
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

		if strings.Index(contents, "(") != -1 {
			if contents, err = decompose(contents); err != nil {
				return "", err
			}
		}

		if f, err = ParseFormula(contents); err != nil {
			return "", err
		} else if result, err = compute(f); err != nil {
			return "", err
		}

		formula = strings.Replace(formula, formula[iOpening:iClosing+1], strconv.FormatFloat(float64(result), 'f', -1, 64), -1)

		AddStep(formula)

		if formula, err = decompose(formula); err != nil {
			return "", err
		}
	} else {
		if f, err = ParseFormula(formula); err != nil {
			return "", err
		} else if result, err = compute(f); err != nil {
			return "", err
		}

		formula = strconv.FormatFloat(float64(result), 'f', -1, 64)
		AddStep(formula)
	}

	return formula, err
}

func compute(fp *FormulaParts) (float64, error) {
	switch fp.Operator {
	case "^":
		return Power(fp.Expression1, fp.Expression2), nil
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

func removeSpaces(formula string) string {
	formula = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, formula)

	return formula
}
