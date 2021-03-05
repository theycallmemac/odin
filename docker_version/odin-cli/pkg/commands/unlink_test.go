package commands

import (
	"fmt"
	"testing"
)

func TestUnlinkJobs(t *testing.T) {
	cases := []struct {
		Name          string
		A, B, C, D, E string
		Expected      error
	}{
		{"test unlink flags -f and -t", "unlink", "-f", "763e8d23ad0f", "-t", "a13ad8b9c77c", nil},
	}
	for i, testCase := range cases {
		t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
			RootCmd.SetArgs([]string{testCase.A, testCase.B, testCase.C, testCase.D, testCase.E})
			actual := RootCmd.Execute()
			if actual != testCase.Expected {
				t.Errorf("TestUnlinkJobs %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)
			}
		})
	}
}
