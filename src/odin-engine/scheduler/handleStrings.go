package scheduler

import (
    "strings"
)

// this function is used to compare two strings
func compare(a, b string) bool {
    return a == b
}

// this function is used to strip characters from string
func stripChars(str, chars string) string {
    return strings.Map(func(r rune) rune {
        if strings.IndexRune(chars, r) < 0 {
            return r
        }
        return -1
    }, str)
}

// this function is used to
func splitOnKeyword(line string, delimiter string) []string {
    return strings.Split(trimEdges(line, "\t \n"), delimiter)
}

func trimEdges(line string, chars string) string {
    return strings.TrimRight(strings.TrimLeft(line, chars), chars)
}
