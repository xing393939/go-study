package option3

import "testing"

func TestNewConnect(t *testing.T) {
	obj1, _ := NewConnect("localhost")
	obj2, _ := NewConnect("localhost", WithTimeout(100))
	if obj1.timeout == obj2.timeout {
		t.Errorf("err")
	}
}
