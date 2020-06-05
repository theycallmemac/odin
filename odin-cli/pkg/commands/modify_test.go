package commands

import (
	"fmt"
	"testing"
)

func TestModifyDetails(t *testing.T) {
	cases := []struct {
		Name          string
		A, B, C, D, E string
		Expected      error
	}{
		{"test log with flag -i and -s", "modify", "-i", "a13ad8b9c77c", "-s", "every minute", nil},
		{"test log with flag -i and -d ", "modify", "-i", "763e8d23ad0f", "-d", "new description", nil},
	}
	for i, testCase := range cases {
		t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
			RootCmd.SetArgs([]string{testCase.A, testCase.B, testCase.C, testCase.D, testCase.E})
			actual := RootCmd.Execute()
			if actual != testCase.Expected {
				t.Errorf("TestModifyDetails %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)
			}
		})
	}
}
