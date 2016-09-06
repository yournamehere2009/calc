package calc

import "testing"

var getNumberBeforeTests = []struct {
	f        string // Formula Before Operator
	expected string // Expected number
}{
	{"2+2", "2"},
	{"2+3+-22", "-22"},
	{"7", "7"},
	{"15", "15"},
}

var getNumberAfterTests = []struct {
	f        string // Formula Before Operator
	expected string // Expected number
}{
	{"2+2", "2"},
	{"-22+3+-22", "-22"},
	{"775-3", "775"},
	{"15", "15"},
}

func TestGetNumberBeforeOperator(t *testing.T) {
	for _, tt := range getNumberBeforeTests {
		result := getNumberBeforeOperator(tt.f)

		if result != tt.expected {
			t.Errorf("getNumberBeforeOperator(%v), expected %v, actual %v", tt.f, tt.expected, result)
		}
	}
}

func TestGetNumberAfterOperator(t *testing.T) {
	for _, tt := range getNumberAfterTests {
		result := getNumberAfterOperator(tt.f)

		if result != tt.expected {
			t.Errorf("getNumberAfterOperator(%v), expected %v, actual %v", tt.f, tt.expected, result)
		}
	}
}
