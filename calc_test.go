package calc_test

import (
	"testing"

	"github.com/yournamehere2009/calc"
)

var addTests = []struct {
	a        float64 // input
	b        float64 // input
	expected float64 // expected result
}{
	{2, 2, 4},
	{10, 11, 21},
	{-1, 1, 0},
	{-1, 2, 1},
}

var subtractTests = []struct {
	a        float64 // input
	b        float64 // input
	expected float64 // expected result
}{
	{2, 2, 0},
	{10, 11, -1},
	{-1, 1, -2},
	{11, -2, 13},
}

var multiplyTests = []struct {
	a        float64 // input
	b        float64 // input
	expected float64 // expected result
}{
	{2, 2, 4},
	{10, 11, 110},
	{-1, 1, -1},
	{11, -2, -22},
}

var divideTests = []struct {
	a        float64 // input
	b        float64 // input
	expected float64 // expected result
}{
	{2, 2, 1},
	{10, 5, 2},
	{5, 2, 2.5},
	{-3, -1, 3},
	{4, 0, 0}, // Expect an error
}

var parseFormulaTests = []struct {
	f string  // formula
	a float64 // expected expression 1
	b float64 // expected expression 2
	o string  // expected operator
}{
	{"10+2", 10, 2, "+"},
	{"7-2", 7, 2, "-"},
	{"15/5", 15, 5, "/"},
	{"20*3", 20, 3, "*"},
	{"3", 3, 0, "+"},
	{"-3", -3, 0, "+"},
}

var computeFormulaTests = []struct {
	f        string  // formula
	expected float64 // expected result
}{
	{"10+2", 12},
	{"(2+2)+2", 6},
	{"(2.5+2)+2", 6.5},
	{"(2.5+2)+(2+8)", 14.5},
	{"(2.5*(2+5))+(2+(8-4))", 23.5},
	{"((10*2)/5)", 4},
	{"3(3)", 9},
}

var addWorkStepTests = []struct {
	step       string // formula
	timesToAdd int    // formula
	totalSteps int    // expected result
}{
	{"10+2", 1, 1},
	{"10+2", 3, 3},
}

var showWorkTests = []struct {
	f          string // formula
	totalSteps int
	expected   float64 // expected result
}{
	{"10+2", 2, 12},
	{"10--2", 3, 12},
	{"(10+2)-3", 3, 9},
}

func TestAdd(t *testing.T) {
	for _, tt := range addTests {
		result := calc.Add(tt.a, tt.b)

		if result != tt.expected {
			t.Errorf("Add(%f,%f): expected %f, actual %f", tt.a, tt.b, tt.expected, result)
		}
	}
}

func TestSubtract(t *testing.T) {
	for _, tt := range subtractTests {
		result := calc.Subtract(tt.a, tt.b)

		if result != tt.expected {
			t.Errorf("Subtract(%f,%f): expected %f, actual %f", tt.a, tt.b, tt.expected, result)
		}
	}
}

func TestMultiply(t *testing.T) {
	for _, tt := range multiplyTests {
		result := calc.Multiply(tt.a, tt.b)

		if result != tt.expected {
			t.Errorf("Multiply(%f,%f): expected %f, actual %f", tt.a, tt.b, tt.expected, result)
		}
	}
}

func TestDivide(t *testing.T) {
	for _, tt := range divideTests {
		result, err := calc.Divide(tt.a, tt.b)

		if tt.b != 0 && result != tt.expected {
			t.Errorf("Divide(%f,%f): expected %f, actual %f", tt.a, tt.b, tt.expected, result)
		} else if tt.b == 0 && err == nil {
			t.Errorf("Divide(%f,%f): expected a divide by zero error", tt.a, tt.b)
		}
	}
}

func TestParseFormula(t *testing.T) {
	for _, tt := range parseFormulaTests {
		formula, _ := calc.ParseFormula(tt.f)

		if formula.Operator != tt.o {
			t.Errorf("ParseFormula(%v): expected operator %v, actual %v", tt.f, tt.o, formula.Operator)
		} else if formula.Expression1 != tt.a {
			t.Errorf("ParseFormula(%v): expected first expression %f, actual %f", tt.f, tt.a, formula.Expression1)
		} else if formula.Expression2 != tt.b {
			t.Errorf("ParseFormula(%v): expected second expression %f, actual %f", tt.f, tt.b, formula.Expression2)
		}
	}
}

func TestComputeFormula(t *testing.T) {
	for _, tt := range computeFormulaTests {
		result, _, _ := calc.ComputeFormula(tt.f)

		if result != tt.expected {
			t.Errorf("ComputeFormula(%v): expected %f, actual %f", tt.f, tt.expected, result)
		}
	}
}

func TestWorkStep(t *testing.T) {
	for _, tt := range addWorkStepTests {
		for i := 0; i < tt.timesToAdd; i++ {
			calc.AddStep(tt.step)
		}

		workSteps := calc.GetSteps()

		if len(workSteps) != tt.totalSteps {
			t.Errorf("AddStep(%v) Show Work: expected steps %d, actual %d, %v", tt.step, tt.totalSteps, len(workSteps), workSteps)
		}
		calc.ClearSteps()
	}
}

func TestComputeFormulaShowWork(t *testing.T) {
	for _, tt := range showWorkTests {
		_, workSteps, _ := calc.ComputeFormula(tt.f)

		if len(workSteps) != tt.totalSteps {
			t.Errorf("ComputeFormula(%v) Show Work: expected steps %d, actual %d, %v", tt.f, tt.totalSteps, len(workSteps), workSteps)
		}
		calc.ClearSteps()
	}
}
