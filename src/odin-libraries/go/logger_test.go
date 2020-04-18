package odinlib

import (
    "fmt"
    "reflect"
    "testing"

    "go.mongodb.org/mongo-driver/mongo"
)

func TestLog(t *testing.T) {
    cases := []struct {Name string; A, B, C, D, E string; Expected bool;} {
        {"watch operation", "watch", "watch this variable", "1000", "fake_id", "1587219680", true},
        {"condition operation", "condition", "evaluate and watch this condition", "1001", "fake_id", "1587219680", true},
        {"result operation", "result", "exit on this result", "1002", "fake_id", "1587219680", true},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("Log(%s, %s, %s, %s, %s) ", testCase.A, testCase.B, testCase.C, testCase.D, testCase.E), func(t *testing.T) {
            actual := Log(testCase.A, testCase.B, testCase.C, testCase.D, testCase.E)
            if (actual != testCase.Expected) {t.Errorf("TestLog %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestFindAndInsert(t *testing.T) {
    cases := []struct {Name string; A, B, C, D, E string; Expected bool;} {
        {"watch operation", "watch", "watch this variable", "1000", "fake_id", "1587219680", true},
        {"condition operation", "condition", "evaluate and watch this condition", "1001", "fake_id", "1587219680", true},
        {"result operation", "result", "exit on this result", "1002", "fake_id", "1587219680", true},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("FindAndInsert(%s, %s, %s, %s, %s) ", testCase.A, testCase.B, testCase.C, testCase.D, testCase.E), func(t *testing.T) {
            actual := FindAndInsert(testCase.A, testCase.B, testCase.C, testCase.D, testCase.E)
            if (actual != testCase.Expected) {t.Errorf("TestFindAndInsert %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestSetupClient(t *testing.T) {
    var client *mongo.Client
    cases := []struct {Name string; Expected *mongo.Client;} {
        {"working client", client},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("SetupClient() "), func(t *testing.T) {
            actual := SetupClient()
            if (reflect.TypeOf(actual) != reflect.TypeOf(testCase.Expected)) {t.Errorf("TestSetupClient %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestGetMongoClient(t *testing.T) {
    var client *mongo.Client
    cases := []struct {Name string; A string; Expected *mongo.Client;} {
        {"working client", "mongodb://localhost:27017", client},
        {"broken client", "nolinkhere", nil},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("GetMongoClient(%s) ", testCase.A), func(t *testing.T) {
            actual := getMongoClient(testCase.A)
            if (reflect.TypeOf(actual) != reflect.TypeOf(testCase.Expected)) {t.Errorf("TestGetMongoClient %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}
