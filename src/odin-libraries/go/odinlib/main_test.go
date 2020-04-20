package odinlib

import (
    "fmt"
    "testing"
)

func TestSetup(t *testing.T) {
    cases := []struct {Name string; A string; Expected1 bool; Expected2 string;} {
        {"read a yaml file", "testConfigs/job.yml", true, ""},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("Setup(%s) ", testCase.A), func(t *testing.T) {
            actual1, actual2 := Setup(testCase.A)
            fmt.Println(actual1, actual2)
            if (actual1 != testCase.Expected1) {t.Errorf("TestSetup %d failed - expected: '%v' got: '%v'", i+1, actual1, testCase.Expected1)}
            if (actual2 != testCase.Expected2) {t.Errorf("TestSetup %d failed - expected: '%v' got: '%v'", i+1, actual2, testCase.Expected2)}
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
            if i == 1 {
                ENV_CONFIG = false
            }
            actual := Condition(testCase.A, testCase.B)
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
            if i == 1 {
                ENV_CONFIG = false
            }
            actual := Watch(testCase.A, testCase.B)
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
            if i == 1 {
                ENV_CONFIG = false
            }
            actual := Result(testCase.A, testCase.B)
            if (actual != testCase.Expected) {t.Errorf("TestResult %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}
