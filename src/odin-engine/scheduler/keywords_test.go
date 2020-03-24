package scheduler

import (
    "fmt"
    "strings"
    "testing"
)

func TestGetValidKeywords(t *testing.T) {
    cases := []struct {Name string; Expected int} {
	{"validate correct start of time", 115},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("getValidKeywords()"), func(t *testing.T) {
            actual := len(strings.Split(getValidKeywords(), " "))
	    if (actual != testCase.Expected) {t.Errorf("TestGetValidKeywords %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestGetDowMap(t *testing.T) {
    cases := []struct {Name string; Expected int} {
	{"validate correct start of time", 8},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("getDowMap()"), func(t *testing.T) {
            actual := len(getDowMap())
	    if (actual != testCase.Expected) {t.Errorf("TestTestGetDowMap %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestGetDomMap(t *testing.T) {
    cases := []struct {Name string; Expected int} {
	{"validate correct start of time", 31},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("getDomMap()"), func(t *testing.T) {
            actual := len(getDomMap())
	    if (actual != testCase.Expected) {t.Errorf("TestTestGetDomMap %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestGetMonMap(t *testing.T) {
    cases := []struct {Name string; Expected int} {
	{"validate correct start of time", 12},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("getMonMap()"), func(t *testing.T) {
            actual := len(getMonMap())
	    if (actual != testCase.Expected) {t.Errorf("TestTestGetMonMap %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

