package commands

import (
    "fmt"
    "testing"
)

func TestExecuteYaml(t *testing.T) {
    cases := []struct {Name string; A, B, C string; Expected error} {
        {"test execute flags -f", "execute", "-f", "testConfigs/scrape.yml", nil},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
	    RootCmd.SetArgs([]string{testCase.A, testCase.B, testCase.C})
            actual := RootCmd.Execute()
            if (actual != testCase.Expected) {t.Errorf("TestExecuteYaml %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestExecuteReadFile(t *testing.T) {
    cases := []struct {Name string; A string; Expected int} {
        {"test existing yaml file", "testConfigs/scrape.yml", 207},
        {"test existing yaml file", "testConfigs/prune_containers.yml", 242},
        {"test empty yaml file", "testConfigs/empty.yml", 4},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
            actual := len(readJobFileExecute(testCase.A))
            if (actual != testCase.Expected) {t.Errorf("TestExecuteReadFile %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

