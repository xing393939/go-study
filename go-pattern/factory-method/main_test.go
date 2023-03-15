package factory_method

import (
	"testing"
)

func TestNewPersonFactory(t *testing.T) {
	newBaby := NewPersonFactory(1)
	baby := newBaby("john")
	if baby.name != "john" {
		t.Errorf("err")
	}

	newTeenager := NewPersonFactory(16)
	teen := newTeenager("jill")
	if teen.name != "jill" {
		t.Errorf("err")
	}
}
