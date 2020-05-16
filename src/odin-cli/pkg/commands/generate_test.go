package commands

import (
    "fmt"
    "testing"
)

func TestGenerateFiles(t *testing.T) {
    cases := []struct {Name string; A, B, C, D, E string; Expected error} {
        {"test generate flags -f and -l", "generate", "-f", "test-job.yml", "-l", "python3", nil},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
	    RootCmd.SetArgs([]string{testCase.A, testCase.B, testCase.C, testCase.D, testCase.E})
            actual := RootCmd.Execute()
            if (actual != testCase.Expected) {t.Errorf("TestGenerateFiles %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestGenerateID(t *testing.T) {
    cases := []struct {Name string; Expected int} {
        {"test identifier generation", 12},
        {"test identifier generation", 12},
        {"test identifier generation", 12},
        {"test identifier generation", 12},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
            actual := len(generateId())
            if (actual != testCase.Expected) {t.Errorf("TestGenerateFiles %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestGenerateCreate(t *testing.T) {
    cases := []struct {Name string; A, B string; Expected string} {
        {"test python3 language", "test-job", "python3", "test-job.py"},
        {"test python language", "test-job", "python3", "test-job.py"},
        {"test python language", "test-job", "go", "test-job.go"},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
            actual := createLanguageFile(testCase.A, testCase.B)
            if (actual != testCase.Expected) {t.Errorf("TestGenerateCreate %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

