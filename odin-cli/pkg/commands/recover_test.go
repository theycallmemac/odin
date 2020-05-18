package commands

import (
    "fmt"
    "testing"
)

func TestRecover(t *testing.T) {
    cases := []struct {Name string; A, B, C string; Expected error} {
        {"test recover flag -i with good ID", "recover", "-i", "763e8d23ad0f", nil},
        {"test recover flag -i with good ID", "recover", "-i", "a13ad8b9c77c", nil},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
	    RootCmd.SetArgs([]string{testCase.A, testCase.B, testCase.C})
            actual := RootCmd.Execute()
            if (actual != testCase.Expected) {t.Errorf("TestRecover %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

