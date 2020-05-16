package commands

import (
    "fmt"
    "testing"

    "gopkg.in/yaml.v2"
)

func TestDeployConfig(t *testing.T) {
    cases := []struct {Name string; A, B, C string; Expected error} {
        {"test deploy flag -f with real file", "deploy", "-f", "testConfigs/scrape.yml", nil},
        {"test deploy flag -f with real file", "deploy", "-f", "testConfigs/job.yml", nil},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
	    RootCmd.SetArgs([]string{testCase.A, testCase.B, testCase.C})
            actual := RootCmd.Execute()
            if (actual != testCase.Expected) {t.Errorf("TestDeployConfig %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestDeployReadJobFile(t *testing.T) {
    cases := []struct {Name string; A string; Expected int} {
        {"read test job file", "testConfigs/job.yml", 246},
        {"read empty test job file", "testConfigs/empty.yml", 4},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v ", testCase.Name), func(t *testing.T) {
            actual := len(readJobFile(testCase.A))
            if (actual != testCase.Expected) {t.Errorf("TestDeployReadJobFile %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestDeployUnmarshalYaml(t *testing.T) {
    var cfg Config
    cases := []struct {Name string; A []byte; Expected Config} {
        {"parse an empty yaml file", readJobFile("testConfigs/empty.yml"), cfg},
        {"parse a standard odin yaml file", readJobFile("testConfigs/prune_containers.yml"), cfg},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.parse() ", testCase.A), func(t *testing.T) {
            yaml.Unmarshal([]byte(testCase.A), &testCase.Expected)
            actual := unmarsharlYaml(testCase.A)
            if (actual != testCase.Expected) {t.Errorf("TestDeployUnmarshalYaml %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestDeployEnsureDirectory(t *testing.T) {
    cases := []struct {Name string; A string; Expected bool} {
        {"check an existing directory", "testConfigs", true},
        {"check a non existent directory", "tastConfig", false},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.parse() ", testCase.A), func(t *testing.T) {
            actual := ensureDirectory(testCase.A)
            if (actual != testCase.Expected) {t.Errorf("TestDeployEnsureDirectory %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestDeployGetScheduleString(t *testing.T) {
    cases := []struct {Name string; A, B string; Expected string} {
        {"check an existing config file", "testConfigs/scrape.yml", ":3939", "* * * * *,"},
        {"check an existing config file", "testConfigs/prune_containers.yml", ":3939", "0 13 9 9 *,0 13 21 3 *,"},
        {"check an non existent config file", "testConfigs/notreal.yml", ":3939", ""},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.parse() ", testCase.A), func(t *testing.T) {
            actual := getScheduleString(testCase.A, testCase.B)
            if (actual != testCase.Expected) {t.Errorf("TestDeployGetScheduleString %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}
