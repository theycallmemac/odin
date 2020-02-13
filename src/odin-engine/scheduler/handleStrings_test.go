package scheduler

import (
    "fmt"
	"testing"
)

func TestCompare(t *testing.T) {
    cases := []struct {Name, A, B string; Expected bool} {
	{"compare equal strings", "abc", "abc", true},
	{"compare strings with a one missing character ","abcdefghijklmno", "abcdeghijklmno", false},
	{"compare empty strings", "", "", true},
	{"compare strings with multiple missing characters", "tuvwxyz", "txz", false},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf(" %s == %s ", testCase.A, testCase.B), func(t *testing.T) {
            actual := compare(testCase.A, testCase.B)
	    if (actual != testCase.Expected) {t.Errorf("TestCompare %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestStripChairs(t *testing.T) {
    cases := []struct {Name, A, B string; Expected string} {
	{"strip on a single letter", "abcd", "d", "abc"},
	{"strip on multiple letters","stuvstuvstuvstuvstuv", "stv", "uuuuu"},
	{"strip on special characters", "//($$*(^@^&$@%^@$@(*)|{}{|$$@#", "${}", "//(*(^@^&@%^@@(*)||@#"},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf(" %s.strip(%s) ", testCase.A, testCase.B), func(t *testing.T) {
            actual := stripChars(testCase.A, testCase.B)
	    if (actual != testCase.Expected) {t.Errorf("TestCompare %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestSplitOnKeyword(t *testing.T) {
    cases := []struct {Name, A, B string; Expected int} {
	{"split on a single character", "12:00", ":", 2},
	{"split on multiple characters", "the.fox.jumped.over.the.dog", ".", 6},
	{"split on a single word","every September 9th at 13:00 and every March 21st at 13:00", "and", 2},
	{"split on multiple words", "every September 9th at 13:00 and every March 21st at 13:00 and everyday at 00:00", "and", 3},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf(" %s.splitOn(%s) ", testCase.A, testCase.B), func(t *testing.T) {
            actual := len(splitOnKeyword(testCase.A, testCase.B))
	    if (actual != testCase.Expected) {t.Errorf("TestCompare %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestTrimEdges(t *testing.T) {
    cases := []struct {Name, A, B string; Expected string} {
	{"trim on left side of the string", " every September 9th at 13:00", "\t \n", "every September 9th at 13:00"},
	{"trim on right side of the string", "every September 9th at 13:00 ", ".\t \n", "every September 9th at 13:00"},
	{"trim on both sides of the string", " every September 9th at 13:00 ", "\t \n", "every September 9th at 13:00"},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf(" %s.trimEdgesFrom(%s) ", testCase.A, testCase.B), func(t *testing.T) {
            actual := trimEdges(testCase.A, testCase.B)
	    if (actual != testCase.Expected) {t.Errorf("TestCompare %d failed - expected: '%v' got: '%v'", i, actual, testCase.Expected)}
        })
    }
}
