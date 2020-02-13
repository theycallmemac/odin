package executor

import (
    "fmt"
    "os"
    "testing"
)


func TestDoesNotExist(t *testing.T) {
    cases := []struct {Name, A string; B int; Expected bool} {
        {"check if an existing file exists", "testConfigs/prune_containers.yml", 1, true},
        {"check if a non-existing file exists", "testConfigs/test4.yml", 1, false},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("execute(%s) ", testCase.A), func(t *testing.T) {
            actual := Execute(testCase.A, testCase.B)
            if (actual != testCase.Expected) {t.Errorf("TestExecute %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)} else {
                os.Remove("output.txt")
            }
        })
    }
}

func TestExecute(t *testing.T) {
    cases := []struct {Name, A string; B int;Expected bool} {
        {"run an existing python file", "testConfigs/prune_containers.yml", 1, true},
        {"run an existing node file", "testConfigs/refresh_token.yml", 1, true},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("execute(%s) ", testCase.A), func(t *testing.T) {
            actual := Execute(testCase.A, testCase.B)
            if (actual != testCase.Expected) {t.Errorf("TestExecute %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)} else {
                _, err := os.Stat("output.txt")
                if err != nil {
                    os.Remove("output.txt")
                }
            }
        })
    }
}
