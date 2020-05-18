package commands

import (
    "fmt"
    "testing"
)

func TestLog(t *testing.T) {
    cases := []struct {Name string; A, B, C string; Expected error} {
        {"test log with flag -i", "log", "-i", "a13ad8b9c77c", nil},
        {"test log with flag -i ", "log", "-i", "763e8d23ad0f", nil},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
	    RootCmd.SetArgs([]string{testCase.A, testCase.B, testCase.C})
            actual := RootCmd.Execute()
            if (actual != testCase.Expected) {t.Errorf("TestLog %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

