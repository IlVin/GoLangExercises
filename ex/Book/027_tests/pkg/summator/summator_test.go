package summator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	s := Summator{}

	s.Add(3.3)

	if "3.300000" != s.String() {
		t.Errorf("3.300000 != %s", s)
	}

	t.Log(s)
}
