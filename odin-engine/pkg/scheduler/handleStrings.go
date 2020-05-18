package scheduler

import (
    "strings"
)

// this function is used to strip characters from string
// parameters:  str (a string to strip from), chars (characters to strip off)
// returns: string (the stripped string)
func stripChars(str, chars string) string {
    return strings.Map(func(r rune) rune {
        if strings.IndexRune(chars, r) < 0 {
            return r
        }
        return -1
    }, str)
}

// this function is used to split a string on a keyword
// parameters: line (a string to be split), delimiter (a string to split on)
// returns: []string (an array of the tokens resulting from splitting `line` on `delimiter`)
func splitOnKeyword(line string, delimiter string) []string {
    return strings.Split(trimEdges(line, "\t \n"), delimiter)
}

// this function is used to trim special characters from either end of a string
// parameters: line (a string to be trimmed), chars (a string of character to trim offn)
// returns: string (the trimmed string) 
func trimEdges(line string, chars string) string {
    return strings.TrimRight(strings.TrimLeft(line, chars), chars)
}
