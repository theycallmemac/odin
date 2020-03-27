package jobs

import (
    "fmt"
    "reflect"
    "testing"
    "time"
)

var unsorted Queue
var sorted Queue
var jobs []NewJob

func TestSortQueue(t *testing.T) {
    node1 := Node{Schedule:[]int{80}}
    node2 := Node{Schedule:[]int{1204}}
    node3 := Node{Schedule:[]int{365}}
    node4 := Node{Schedule:[]int{201}}
    sorted.Items = append(sorted.Items, node1, node4, node3, node2)
    unsorted.Items = append(unsorted.Items, node1, node2, node3, node4)
    cases := []struct {Name string; A []Node; Expected []Node} {
        {"sort an unsorted array of nodes", unsorted.Items, sorted.Items},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
                actual := testCase.A
                channel := make(chan int)
                go sortQueue(actual, channel)
                <-channel
                for inc, _ := range unsorted.Items {
                    if (actual[inc].Schedule[0] != testCase.Expected[inc].Schedule[0]) {t.Errorf("TestSortQueue %d failed - expected: '%v' got: '%v'", i+1, actual[inc].Schedule, testCase.Expected[inc].Schedule)}
                }
        })
    }
}

func TestCheckHead(t *testing.T) {
    node1 := Node{Schedule:[]int{80}}
    node2 := Node{Schedule:[]int{80}}
    unsorted.Items = append(sorted.Items, node1, node2)
    var mapItems80 = map[int][]Node {
        80: unsorted.Items,
    }
    var mapItems0 = map[int][]Node {
        0: unsorted.Items,
    }
    cases := []struct {Name string; A map[int][]Node; Expected bool} {
        {"check head of map with no zero value", mapItems80, false},
        {"check head of map with a zero value", mapItems0, true},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
            actual := checkHead(testCase.A)
            if (actual != testCase.Expected) {t.Errorf("TestCheckHead %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestGroupItems(t *testing.T) {
    node1 := Node{Schedule:[]int{80}}
    node2 := Node{Schedule:[]int{1204}}
    node3 := Node{Schedule:[]int{365}}
    node4 := Node{Schedule:[]int{201}}
    var returnMap = map[int][]Node {80: {node1}, 201: {node4}, 365: {node3},1204: {node2},
    }
    cases := []struct {Name string; A []Node; Expected map[int][]Node} {
        {"group array of sorted nodes", sorted.Items, returnMap},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
            actual := groupItems(testCase.A)
            if (!reflect.DeepEqual(actual, testCase.Expected)) {t.Errorf("TestGroupItems %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestCronToSeconds(t *testing.T) {
    var result int
    var expected int
    cases := []struct {Name string; A string; Expected []int} {
        {"every minute", "* * * * *,", []int{0}},
        {"every hour", "0 * * * *,", []int{0}},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
            if i == 0 {
                testCase.Expected[0] = time.Now().Second()
                actual := cronToSeconds(testCase.A)
                result = 59 - actual[0]
                expected = testCase.Expected[0]
            } else {
                testCase.Expected[0] = (time.Now().Minute() * 60) + time.Now().Second()
                actual := cronToSeconds(testCase.A)
                result = actual[0]
                expected = 3599 - testCase.Expected[0]
            }
            if (result != expected) {t.Errorf("TestCronToSeconds %d failed - expected: '%v' got: '%v'", i+1, result, expected)}
        })
    }
}

func TestFillQueue(t *testing.T) {
    var filled []Node
    node1 := NewJob{Schedule:"* * * * *,"}
    node2 := NewJob{Schedule:"0 * * * *,"}
    jobs = append(jobs, node1, node2)
    cases := []struct {Name string; A []NewJob; Expected []Node} {
        {"group array of sorted nodes", jobs, filled},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.get() ", testCase.A), func(t *testing.T) {
            actual := fillQueue(testCase.A)
            if ((reflect.TypeOf(actual)) != reflect.TypeOf(testCase.Expected)) {t.Errorf("TestFillQueue %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

