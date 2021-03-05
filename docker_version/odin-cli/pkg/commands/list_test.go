package commands

import (
	"fmt"
	"testing"
)

func TestListJobs(t *testing.T) {
	cases := []struct {
		Name     string
		A        string
		Expected error
	}{
		{"test list", "list", nil},
	}
	for i, testCase := range cases {
		t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
			RootCmd.SetArgs([]string{testCase.A})
			actual := RootCmd.Execute()
			if actual != testCase.Expected {
				t.Errorf("TestListJobs %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)
			}
		})
	}
}
