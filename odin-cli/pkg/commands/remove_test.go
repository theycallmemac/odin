package commands

import (
    "fmt"
    "testing"
)

func TestRemove(t *testing.T) {
    cases := []struct {Name string; A, B, C string; Expected error} {
        {"test remove flag -i with good ID", "remove", "-i", "763e8d23ad0f", nil},
        {"test remove flag -i with good ID", "remove", "-i", "a13ad8b9c77c", nil},
        {"test remove flag -i with bad ID", "remove", "-i", "doesnt-exist", nil},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
	    RootCmd.SetArgs([]string{testCase.A, testCase.B, testCase.C})
            actual := RootCmd.Execute()
            if (actual != testCase.Expected) {t.Errorf("TestRemove %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

