package singleton

import (
	"reflect"
	"testing"
)

func TestGetInsOr(t *testing.T) {
	if got := GetInsOr(); !reflect.DeepEqual(got, ins) {
		t.Errorf("err")
	}
}
