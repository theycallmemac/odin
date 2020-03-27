package jobs

import (
    "fmt"
    "testing"
)

func TestNotDirectory(t *testing.T) {
    cases := []struct {Name string; A string; Expected bool} {
        {"check a directory which exists", "/etc", false},
        {"check a directory which does not exist", "/bin/banter", true},
        {"check a directory which does not exist", "/dev/ram", true},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
            actual := notDirectory(testCase.A)
            if (actual != testCase.Expected) {t.Errorf("TestNotDirectory %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestMakeDirectory(t *testing.T) {
    cases := []struct {Name string; A string; Expected bool} {
        {"make a directory under user home directory", "./here/this", true},
        {"make a directory under user home directory", "./now/this", true},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
            actual := makeDirectory(testCase.A)
            if (actual != testCase.Expected) {t.Errorf("TestMakeDirectory %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestChownR(t *testing.T) {
    cases := []struct {Name string; A string; B int; Expected bool} {
        {"chown a directory which exists", "./here", 1000, true},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
            actual := ChownR(testCase.A, testCase.B, testCase.B)
            if (actual != testCase.Expected) {t.Errorf("TestChownR %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

