package resources

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/theycallmemac/odin/odin-engine/pkg/types"
)

var cfg types.JobConfig
var f *os.File
var b []byte

func TestExecutorYaml(t *testing.T) {
	cases := []struct {
		Name     string
		A        string
		Expected string
	}{
		{"parse an empty yaml file", "./testConfigs/empty.yml", ""},
		{"parse a standard odin yaml file", "./testConfigs/prune_containers.yml", "python3"},
		{"parse an non-existent yaml file", "./testConfigs/false.yml", "Failed to read file."},
	}
	for i, testCase := range cases {
		t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
			actual, _ := ExecutorYaml(testCase.A)
			if actual != testCase.Expected {
				t.Errorf("TestExecutorYaml %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)
			}
		})
	}
}

func TestSchedulerYaml(t *testing.T) {
	cases := []struct {
		Name     string
		A        string
		Expected string
	}{
		{"parse an empty yaml file", "testConfigs/empty.yml", ""},
		{"parse a standard odin yaml file", "testConfigs/prune_containers.yml", "every September 9th at 13:00 and every March 21st at 13:00"},
		{"parse an non-existent yaml file", "./testConfigs/false.yml", "Failed to read file."},
	}
	for i, testCase := range cases {
		t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
			actual := SchedulerYaml(testCase.A)
			if actual != testCase.Expected {
				t.Errorf("TestSchedulerYaml %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)
			}
		})
	}
}

func TestNotEmpty(t *testing.T) {
	cases := []struct {
		Name     string
		A        string
		Expected bool
	}{
		{"check a non-empty string", "hello", true},
		{"check an empty string", "", false},
	}
	for i, testCase := range cases {
		t.Run(fmt.Sprintf("read(%s) ", testCase.A), func(t *testing.T) {
			actual := NotEmpty(testCase.A)
			if actual != testCase.Expected {
				t.Errorf("TestNotEmpty %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	cases := []struct {
		Name     string
		A        string
		Expected *os.File
	}{
		{"read a yaml file", "testConfigs/readme.yml", f},
		{"read a txt file", "testConfigs/readme.txt", f},
		{"read a json", "testConfigs/readme.json", f},
		{"read a toml", "testConfigs/readme.toml", f},
		{"read a md file", "testConfigs/readme.md", f},
	}
	for i, testCase := range cases {
		t.Run(fmt.Sprintf("read(%s) ", testCase.A), func(t *testing.T) {
			actual := ReadFile(testCase.A)
			if actual == testCase.Expected {
				t.Errorf("TestReadFile %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)
			}
		})
	}
}

func TestReadFileBytes(t *testing.T) {
	cases := []struct {
		Name     string
		A        string
		Expected []byte
	}{
		{"read a yaml file", "testConfigs/readme.yml", b},
		{"read a txt file", "testConfigs/readme.txt", b},
		{"read a json", "testConfigs/readme.json", b},
		{"read a toml", "testConfigs/readme.toml", b},
		{"read a md file", "testConfigs/readme.md", b},
		{"read a non-existent yaml file", "testConfigs/false.yml", b},
	}
	for i, testCase := range cases {
		t.Run(fmt.Sprintf("read(%s) ", testCase.A), func(t *testing.T) {
			actual := ReadFileBytes(testCase.A)
			if reflect.TypeOf(actual) != reflect.TypeOf(testCase.Expected) {
				t.Errorf("TestReadFileBytes %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)
			}
		})
	}
}

func TestParseYaml(t *testing.T) {
	cases := []struct {
		Name     string
		A        *types.JobConfig
		B        *os.File
		Expected bool
	}{
		{"parse an empty yaml file", &cfg, ReadFile("testConfigs/empty.yml"), true},
		{"parse a standard odin yaml file", &cfg, ReadFile("testConfigs/prune_containers.yml"), true},
		{"parse a large yaml file", &cfg, ReadFile("testConfigs/large.yml"), true},
		{"parse a non-existent yaml file", &cfg, ReadFile("testConfigs/false.yml"), false},
	}
	for i, testCase := range cases {
		t.Run(fmt.Sprintf("%v.parse(%s) ", testCase.B, testCase.A), func(t *testing.T) {
			actual := ParseYaml(testCase.A, testCase.B)
			if actual != testCase.Expected {
				t.Errorf("TestParseYaml %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)
			}
		})
	}
}

func TestUnmarshalYaml(t *testing.T) {
	var cfg types.EngineConfig
	cases := []struct {
		Name     string
		A        []byte
		Expected types.EngineConfig
	}{
		{"parse an empty yaml file", ReadFileBytes("testConfigs/empty.yml"), cfg},
		{"parse a standard odin yaml file", ReadFileBytes("testConfigs/prune_containers.yml"), cfg},
		{"parse a large yaml file", ReadFileBytes("testConfigs/large.yml"), cfg},
	}
	for i, testCase := range cases {
		t.Run(fmt.Sprintf("%v.parse() ", testCase.A), func(t *testing.T) {
			actual := UnmarsharlYaml(testCase.A)
			if actual != testCase.Expected {
				t.Errorf("TestUnmarshalYaml %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)
			}
		})
	}
}
