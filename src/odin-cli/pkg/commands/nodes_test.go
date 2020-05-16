package commands

import (
    "fmt"
    "testing"
)

func TestNodesGet(t *testing.T) {
    cases := []struct {Name string; A, B string; Expected error} {
        {"test nodes get", "nodes", "get", nil},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
	    RootCmd.SetArgs([]string{testCase.A, testCase.B})
            actual := RootCmd.Execute()
            if (actual != testCase.Expected) {t.Errorf("TestNodes %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

