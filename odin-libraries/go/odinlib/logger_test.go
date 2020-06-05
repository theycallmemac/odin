package odinlib

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	cases := []struct {
		Name          string
		A, B, C, D, E string
		Expected      bool
	}{
		{"watch operation", "watch", "watch this variable", "1000", "fake_id", "1587219680", true},
		{"condition operation", "condition", "evaluate and watch this condition", "1001", "fake_id", "1587219680", true},
		{"result operation", "result", "exit on this result", "1002", "fake_id", "1587219680", true},
	}
	for i, testCase := range cases {
		t.Run(fmt.Sprintf("Log(%s, %s, %s, %s, %s) ", testCase.A, testCase.B, testCase.C, testCase.D, testCase.E), func(t *testing.T) {
			actual := Log(testCase.A, testCase.B, testCase.C, testCase.D, testCase.E)
			if actual != testCase.Expected {
				t.Errorf("TestLog %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)
			}
		})
	}
}
