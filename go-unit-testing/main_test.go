package main

import "testing"

var testScenarios = []struct {
	scenario string
	dividend float64
	divisor  float64
	expected float64
	isError  bool
}{
	{"Divide 100 by 10", 100.0, 10.0, 10.0, false},
	{"Divide 10 by 2", 10.0, 2.0, 5.0, false},
	{"Attempt to divide by 0", 10.0, 0.0, 0, true},
	{"Divide 0 by 100", 0.0, 100.0, 0.0, false},
}

func TestDivide(t *testing.T) {
	for _, test := range testScenarios {
		result, err := divide(test.dividend, test.divisor)
		if test.isError {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one", err.Error())
			}
		}
		if result != test.expected {
			t.Errorf("Expected %f but got %f", test.expected, result)
		}
	}
}

// tests below are not needed because they are covered by the table driven test above
func TestDivideByZero(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Expected an error but did not get one")
	}
}

func TestDivideByNonZero(t *testing.T) {
	_, err := divide(10.0, 2.0)
	if err != nil {
		t.Error("Did not expect an error but got one", err.Error())
	}
}
