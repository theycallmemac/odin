package commands

import (
	"fmt"
	"testing"
)

func TestDescribeJob(t *testing.T) {
	cases := []struct {
		Name     string
		A, B, C  string
		Expected error
	}{
		{"test describe flag -i with good ID", "describe", "-i", "763e8d23ad0f", nil},
		{"test describe flag -i with bad ID", "describe", "-i", "doesnt-exist", nil},
	}
	for i, testCase := range cases {
		t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
			RootCmd.SetArgs([]string{testCase.A, testCase.B, testCase.C})
			actual := RootCmd.Execute()
			if actual != testCase.Expected {
				t.Errorf("TestDescribeJob %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)
			}
		})
	}
}
