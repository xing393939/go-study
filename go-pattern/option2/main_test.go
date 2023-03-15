package option2

import "testing"

func TestNewConnect(t *testing.T) {
	obj1, _ := NewConnect("localhost", NewDefaultOptions())
	obj2, _ := NewConnect("localhost", &ConnectionOptions{Timeout: 100})
	if obj1.timeout == obj2.timeout {
		t.Errorf("err")
	}
}
