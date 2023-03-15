package abstract_factory

import (
	"testing"
)

func TestNewPerson(t *testing.T) {
	got := NewPerson("tt.args.name", 18)
	got.Greet()
}
