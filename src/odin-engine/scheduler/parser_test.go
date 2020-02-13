package scheduler

import (
    "fmt"
	"testing"
)

func TestIsTimeValid(t *testing.T) {
    cases := []struct {Name, A, B string; Expected int} {
	{"validate correct start of time", "18:04", "18", 1},
	{"validate correct end of time", "09:34", "34", 1},
	{"validate incorrect start of time", "23:58", "10", 0},
	{"validate incorrect end  of time ", "21:13", "05", 0},
    }
    var results []string
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("isTimeValid(%s) ", testCase.A), func(t *testing.T) {
            actual, _ := isTimeValid(testCase.A, testCase.B, results)
	    if (len(actual) != testCase.Expected) {t.Errorf("TestIsTimeValid %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestIsScheduleValid(t *testing.T) {
    cases := []struct {Name, A string; Expected bool} {
            {"read schedule with incorrect schedule string", getYaml("testConfigs/scrape_dcufm.yml"), false},
            {"read schedule with correct schedule string", getYaml("testConfigs/prune_containers.yml"), true},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("isScheduleValid(%s) ", testCase.A), func(t *testing.T) {
            actual := isScheduleValid(testCase.A)
            if (actual != testCase.Expected) {t.Errorf("TestIsTimeValid %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestGetCron(t *testing.T) {
    cases := []struct {Name string; A map[string]string; B string; Expected string} {
            {"read Dow values with correct key phrase", getDowMap(), "every Monday", "1"},
            {"read Dow values with correct key phrase", getDowMap(), "every Sunday", "7"},
            {"read Dow values with incorrect key phrase", getDowMap(), "every Mon", "*"},
            {"read Dow values with everday key phrase", getDowMap(), "everyday", "000"},
            {"read Dom values with correct key phrase", getDomMap(), "every 4th", "4"},
            {"read Dom values with correct key phrase", getDomMap(), "every 21st", "21"},
            {"read Dom values with incorrect key phrase", getDomMap(), "every 32nd", "*"},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("getCron(%s) ", testCase.A), func(t *testing.T) {
            actual := getCron(testCase.A, testCase.B)
            if (actual != testCase.Expected) {t.Errorf("TestIsTimeValid %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestGetCronMonth(t *testing.T) {
    cases := []struct {Name string; A map[string]string; B, C  string; Expected string} {
            {"read Mon values with correct key phrase", getMonMap(), "", "every May", "5"},
            {"read Mon values with correct key phrase", getMonMap(), "*", "every February", "2"},
            {"read Mon values with incorrect key phrase", getMonMap(), "*", "every Nov", "*"},
            {"read Mon values with incorrect key phrase", getMonMap(), "", "every Mar", "*"},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("getCronMonth(%s) ", testCase.A), func(t *testing.T) {
            actual, _ := getCronMonth(testCase.A, testCase.B, testCase.C)
            if (actual != testCase.Expected) {t.Errorf("TestIsTimeValid %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}
