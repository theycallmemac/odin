package resources

import (
    "fmt"
    "os"
    "testing"

    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/types"
)

var cfg types.JobConfig
var f *os.File

func TestExecutorYaml(t *testing.T) {
    cases := []struct {Name string; A string; Expected string} {
        {"parse an empty yaml file", "./testConfigs/empty.yml", ""},
        {"parse a standard odin yaml file", "./testConfigs/prune_containers.yml", "python3"},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
            actual, _ := ExecutorYaml(testCase.A)
            if (actual != testCase.Expected) {t.Errorf("TestGetYaml %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestSchedulerYaml(t *testing.T) {
    cases := []struct {Name string; A string; Expected string} {
        {"parse an empty yaml file", "testConfigs/empty.yml", ""},
        {"parse a standard odin yaml file", "testConfigs/prune_containers.yml", "every September 9th at 13:00 and every March 21st at 13:00"},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
            actual := SchedulerYaml(testCase.A)
            if (actual != testCase.Expected) {t.Errorf("TestGetYaml %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestReadFile(t *testing.T) {
    cases := []struct {Name string; A string; Expected *os.File} {
        {"read a yaml file", "testConfigs/readme.yml", f},
        {"read a txt file", "testConfigs/readme.txt", f},
        {"read a json", "testConfigs/readme.json", f},
        {"read a toml", "testConfigs/readme.toml", f},
        {"read a md file", "testConfigs/readme.md", f},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("read(%s) ", testCase.A), func(t *testing.T) {
            actual := ReadFile(testCase.A)
            if (actual == testCase.Expected) {t.Errorf("TestCompare %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestParseYaml(t *testing.T) {
    cases := []struct {Name string; A *types.JobConfig; B *os.File; Expected bool} {
        {"parse an empty yaml file", &cfg, ReadFile("testConfigs/empty.yml"), true},
        {"parse a standard odin yaml file", &cfg, ReadFile("testConfigs/prune_containers.yml"), true},
        {"parse a large yaml file", &cfg, ReadFile("testConfigs/large.yml"), true},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.parse(%s) ", testCase.B, testCase.A), func(t *testing.T) {
            actual := ParseYaml(testCase.A, testCase.B)
            if (actual != testCase.Expected) {t.Errorf("TestCompare %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}
