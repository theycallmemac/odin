package executor

import (
    "fmt"
    "os"
    "testing"
)


func TestDoesNotExist(t *testing.T) {
    cases := []struct {Name, A string; Expected bool} {
        {"check if an existing file exists", "testConfigs/prune_containers.yml", true},
        {"check if a non-existing file exists", "testConfigs/test4.yml", false},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("execute(%s) ", testCase.A), func(t *testing.T) {
            actual := Execute(testCase.A)
            if (actual != testCase.Expected) {t.Errorf("TestExecute %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)} else {
                os.Remove("output.txt")
            }
        })
    }
}

func TestExecute(t *testing.T) {
    cases := []struct {Name, A string; Expected bool} {
        {"run an existing python file", "testConfigs/prune_containers.yml", true},
        {"run an existing node file", "testConfigs/refresh_token.yml", true},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("execute(%s) ", testCase.A), func(t *testing.T) {
            actual := Execute(testCase.A)
            if (actual != testCase.Expected) {t.Errorf("TestExecute %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)} else {
                _, err := os.Stat("output.txt")
                if err != nil {
                    os.Remove("output.txt")
                }
            }
        })
    }
}
