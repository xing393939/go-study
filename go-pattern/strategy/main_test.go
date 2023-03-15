package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	operator := Operator{}

	operator.setStrategy(&add{})
	result := operator.calculate(1, 2)
	if result != 3 {
		t.Errorf("err")
	}

	operator.setStrategy(&reduce{})
	result = operator.calculate(2, 1)
	if result != 1 {
		t.Errorf("err")
	}
}
