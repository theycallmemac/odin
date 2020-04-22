package odinlib

import (
    "fmt"
    "testing"
)


func TestSetup(t *testing.T) {
    cases := []struct {Name string; A string; Expected string;} {
        {"read a yaml file", "testConfigs/job.yml", "13a772a848e1"},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("Setup(%s) ", testCase.A), func(t *testing.T) {
            odin, _ := Setup(testCase.A)
            if (odin.ID != testCase.Expected) {t.Errorf("TestSetup %d failed - expected: '%v' got: '%v'", i+1, testCase.Expected, odin.ID)}
        })
    }
}

func TestCondition(t *testing.T) {
    cases := []struct {Name string; A, B string; Expected bool;} {
        {"a working operation", "watch this variable", "1000", true},
        {"a broken operation", "watch this variable", "1000", false},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("Condition(%s, %s) ", testCase.A, testCase.B), func(t *testing.T) {
            odin, _ := Setup(testCase.A)
            if i == 1 {
                ENV_CONFIG = false
            }
            actual := odin.Condition(testCase.A, testCase.B)
            if (actual != testCase.Expected) {t.Errorf("TestCondition %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}
func TestWatch(t *testing.T) {
    ENV_CONFIG = true
    cases := []struct {Name string; A, B string; Expected bool;} {
        {"a working operation", "watch this variable", "1000", true},
        {"a broken operation", "watch this variable", "1000", false},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("Watch(%s, %s) ", testCase.A, testCase.B), func(t *testing.T) {
            odin, _ := Setup(testCase.A)
            if i == 1 {
                ENV_CONFIG = false
            }
            actual := odin.Watch(testCase.A, testCase.B)
            if (actual != testCase.Expected) {t.Errorf("TestWatch %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestResult(t *testing.T) {
    ENV_CONFIG = true
    cases := []struct {Name string; A, B string; Expected bool;} {
        {"a working operation", "watch this variable", "1000", true},
        {"a broken operation", "watch this variable", "1000", false},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("Result(%s, %s) ", testCase.A, testCase.B), func(t *testing.T) {
            odin, _ := Setup(testCase.A)
            if i == 1 {
                ENV_CONFIG = false
            }
            actual := odin.Result(testCase.A, testCase.B)
            if (actual != testCase.Expected) {t.Errorf("TestResult %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}
