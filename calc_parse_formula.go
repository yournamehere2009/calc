package calc

import (
    "strings"
    "strconv"
    //"errors"
)

// FormulaParts blah
type FormulaParts struct {
    Operator string
    Expression1 float64
    Expression2 float64
}

//ParseFormula blah
func ParseFormula(formula string) (*FormulaParts, error) {
    f := new(FormulaParts);
    oIndex := -1;
    operators := [4]string{"+", "-", "*", "/"};

    // First get the operator
    for i := 0; i < len(operators); i++ {
        oIndex = strings.Index(formula, operators[i]);

        if oIndex != -1 {
            f.Operator = formula[oIndex : oIndex + 1];
            break
        }
    }

    if oIndex == -1 {
       f.Operator = "+";
       
       e1, _ := strconv.ParseFloat(strings.TrimSpace(formula), 64);

       f.Expression1 = e1;
       f.Expression2 = float64(0);

       return f, nil;
    }

    // Next, get the first expression
    e1, _ := strconv.ParseFloat(strings.TrimSpace(formula[:oIndex]), 64);
    f.Expression1 = e1;

    // Finally, get the second expression
    e2, _ := strconv.ParseFloat(strings.TrimSpace(formula[oIndex + 1 : len(formula)]), 64);
    f.Expression2 = e2;

    return f, nil;
}