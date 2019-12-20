package parser

import "strings"

func compare(a, b string) bool {
    if len(a) != len(b) {
        return false
    }
    for i := 0; i < len(a); i++ {
        if a[i] == b[i] {
            continue
        }
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func stripChars(str, chars string) string {
    return strings.Map(func(r rune) rune {
        if strings.IndexRune(chars, r) < 0 {
            return r
        }
        return -1
    }, str)
}

func splitOnKeyword(line string, delimiter string) []string {
    return strings.Split(trimEdges(line, "\t \n"), delimiter)
}

func trimEdges(line string, chars string) string {
    return strings.TrimRight(strings.TrimLeft(line, chars), chars)
}
