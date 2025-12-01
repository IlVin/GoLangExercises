package test_example

import (
	"testing"
)

func TestF1(t *testing.T) {
	expectedResult := "F1"
	if gotResult := F1(); gotResult != expectedResult {
		t.Errorf("expected %s; got: %s", expectedResult, gotResult)
	}

}
