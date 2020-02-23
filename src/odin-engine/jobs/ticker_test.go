package jobs

import (
    "fmt"
    "testing"
)

var unsorted Queue
var sorted Queue

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
                    if (actual[inc].Schedule[0] != testCase.Expected[inc].Schedule[0]) {t.Errorf("TestGetYaml %d failed - expected: '%v' got: '%v'", i+1, actual[inc].Schedule, testCase.Expected[inc].Schedule)}
                }
        })
    }
}

