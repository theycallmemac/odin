package executor

import (
	"errors"
	"fmt"
	"testing"
)

func TestExists(t *testing.T) {
	cases := []struct {
		Name     string
		A        string
		Expected bool
	}{
		{"check if directory exists", "/bin", true},
		{"check if file  exists", "/bin/bash", true},
		{"check if directory does not exist", "/bin123", false},
		{"check if file does not exist", "/bin/file.out", false},
	}
	for i, testCase := range cases {
		t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
			actual := exists(testCase.A)
			if actual != testCase.Expected {
				t.Errorf("TestExists %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)
			}
		})
	}
}

func TestProcessError(t *testing.T) {
	cases := []struct {
		Name     string
		A        error
		B        string
		Expected bool
	}{
		{"check an empty error", errors.New(""), "bool", true},
		{"check an empty error", errors.New("/bin/banter"), "dir", true},
		{"check an empty error", errors.New(""), "nothing", false},
	}
	for i, testCase := range cases {
		t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
			actual := ProcessError(testCase.A, testCase.B)
			if actual != testCase.Expected {
				t.Errorf("TestProcessError %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)
			}
		})
	}
}

func TestReviewError(t *testing.T) {
	cases := []struct {
		Name     string
		A        error
		B        string
		Expected bool
	}{
		{"check an empty error", errors.New(""), "bool", true},
		{"check an empty error", errors.New("/bin/banter"), "dir", true},
	}
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
			ReviewError(testCase.A, testCase.B)
		})
	}
}
