package executor

import (
        "fmt"
	"testing"
)

func TestExists(t *testing.T) {
    cases := []struct {Name string; A string; Expected bool} {
        {"check if directory exists", "/bin", true},
        {"check if file  exists", "/bin/bash", true},
        {"check if directory does not exist", "/bin123", false},                                                                               {"check if file does not exist", "/bin/file.out", false},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
            actual := exists(testCase.A)
            if (actual != testCase.Expected) {t.Errorf("TestGetYaml %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}
