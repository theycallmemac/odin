package scheduler

import (
    "fmt"
    "testing"

    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/pkg/resources"
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
            {"read schedule with incorrect schedule string", resources.SchedulerYaml("testConfigs/scrape_dcufm.yml"), false},
            {"read schedule with correct schedule string", resources.SchedulerYaml("testConfigs/prune_containers.yml"), true},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("isScheduleValid(%s) ", testCase.A), func(t *testing.T) {
            actual := isScheduleValid(testCase.A)
            if (actual != testCase.Expected) {t.Errorf("TestIsScheduleValid %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestGetCron(t *testing.T) {
    cases := []struct {Name string; A map[string]string; B string; Expected string} {
            {"read Dow values with correct key phrase", getDowMap(), "every Monday", "1"},
            {"read Dow values with correct key phrase", getDowMap(), "every Sunday", "7"},
            {"read Dow values with incorrect key phrase", getDowMap(), "every Mon", "*"},
            {"read Dow values with everday key phrase", getDowMap(), "everyday", "*"},
            {"read Dom values with correct key phrase", getDomMap(), "every 4th", "4"},
            {"read Dom values with correct key phrase", getDomMap(), "every 21st", "21"},
            {"read Dom values with incorrect key phrase", getDomMap(), "every 32nd", "*"},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("getCron(%s) ", testCase.A), func(t *testing.T) {
            actual := getCron(testCase.A, testCase.B)
            if (actual != testCase.Expected) {t.Errorf("TestGetCron %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestGetCronMonth(t *testing.T) {
    cases := []struct {Name string; A map[string]string; B, C  string; Expected string} {
            {"read Mon values with correct key phrase", getMonMap(), "", "every", "*"},
            {"read Mon values with correct key phrase", getMonMap(), "", "every May", "5"},
            {"read Mon values with correct key phrase", getMonMap(), "*", "every February", "2"},
            {"read Mon values with incorrect key phrase", getMonMap(), "*", "every Nov", "*"},
            {"read Mon values with incorrect key phrase", getMonMap(), "", "every Mar", "*"},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("getCronMonth(%s) ", testCase.A), func(t *testing.T) {
            actual, _ := getCronMonth(testCase.A, testCase.B, testCase.C)
            if (actual != testCase.Expected) {t.Errorf("TestGetCronMonth %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestExecute(t *testing.T) {
    cases := []struct {Name string; A string; Expected int} {
            {"read config file with incorrect schedule", "testConfigs/scrape_dcufm.yml", 0},
            {"read config file with correct two time schedule", "testConfigs/prune_containers.yml", 2},
            {"read config file with 'every hour' schedule", "testConfigs/hour.yml", 1},
            {"read config file with 'every minute' schedule", "testConfigs/minute.yml", 1},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("Execute(%s) ", testCase.A), func(t *testing.T) {
            actual := len(Execute(testCase.A))
            if (actual != testCase.Expected) {t.Errorf("TestExecute %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}
