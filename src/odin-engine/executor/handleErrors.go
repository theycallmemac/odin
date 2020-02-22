package executor

import (
    "fmt"
    "os"
)

// this function is used to check if a directory exists
// parameters: name (a string containing the path to the directory)
// returns: boolean (returns the value from ProcessError)
func exists(name string) bool {
    _, err := os.Stat(name)
    return ProcessError(err, "dir")
}

// this function is used to switch between the different error types
// parameters: err (an error to use in a given case), errType (the string used to switch between cases)
// returns: boolean (returns the true if a case executes successfulyy, false if otherwise)
func ProcessError(err error, errType string) bool {
    switch errType {
        case "bool":
            return err != nil
        case "dir":
            return !os.IsNotExist(err)
    }
    return false
}


// this function is used to print out an error if it exists
// parameters: err (an error to print), errType (the string given to ProcessError)
// returns: nil
func ReviewError(err error, errType string) {
    if ProcessError(err, errType) {
        fmt.Println(err)
    }
}

