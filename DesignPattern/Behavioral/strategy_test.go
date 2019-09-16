package behavioraldesignpattern

import (
	"testing"
)

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

func TestStrategy(t *testing.T) {
	operation := &Operation{
		Operator: &Addition{},
	}

	if operation.Operate(1, 2) != 3 {
		t.Fatal("Operate Error~")
	}

	operation.Operator = &Multiplication{}

	if operation.Operate(2, 2) != 4 {
		t.Fatal("Operate Error~")
	}

}
