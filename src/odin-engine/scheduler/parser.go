package scheduler

import (
    "fmt"
    "strings"
)

// create StringFormat type to tbe used for accessing time information
type StringFormat struct {
    Minute string
    Hour string
    Dom string
    Mon string
    Dow string
}

// this function is used to judge the validity of the time in a schedule by returning an array of 0's for each matching segment
// parameters: time (the time to be checked), matchMe (a time to match against), results (an array of 0's)
// returns: results (resultant array of strings), addOn (an offset to be returned and used)
func isTimeValid(time string, matchMe string, results []string) ([]string, float64) {
    var addOn float64 = 0
    timeSplit := splitOnKeyword(time, ":")
    for _, ts := range timeSplit {
        if ts == matchMe {
            results = append(results, "0")
            addOn = 0.5
        }
    }
    return results, addOn
}

// this function is used to judge the validity of a schedule
// parameters: schedule (the schedule to be checked)
// returns: boolean (true indicating the schedule is valid, false if otherwise)
func isScheduleValid(schedule string) bool {
    var results []string
    var timeSplitAddOn float64 = 0
    var addOn float64
    words := strings.Fields(schedule)
    for _, word := range words {
        for _, valid := range strings.Fields(getValidKeywords()) {
            if strings.Contains(word, ":") {
                results, addOn = isTimeValid(word, valid, results)
                timeSplitAddOn += addOn
            } else {
                if valid == word {
                    results = append(results, "0")
                    break
                }
            }
        }
    }
    return len(results) == len(words)+int(timeSplitAddOn)
}

// this function is used to get the cron values for the day of the week and the day of the month
// parameters: values (a map of odin formatted time strings to number valued string), key (a string - value to use as an index for the map)
// returns: string (the value of the map at the given key)
func getCron(values map[string]string, key string) string {
    var result string
    if values[key] == "" {
        // if the value doesn't exist, we assume *
        result = "*"
    } else {
        // if the value does exist we set the result to the value
        result = values[key]
    }
    return result
}

// this function is used to get the cron value for a month
// parameters: values (a map of odin formatted time strings to number valued string), currentDom (the current value for dom), key (a string - value to use as an index for the map)
// returns: two strings (the result of the month cron value and the value of currentDom)
func getCronMonth(values map[string]string, currentDom string, key string) (string, string) {
    var result string
    var newKey string
    splitKey := strings.Split(key, " ")
    if len(splitKey) == 1 {
        newKey = splitKey[0]
    } else {
        newKey = strings.Join(splitKey[0:2]," ")
    }
    if values[newKey] == "" {
        // if the value doesn't exist, we assume *
        result = "*"
    } else {
        // if the value does exist we set the result to the value
        result = values[newKey]
        // if the current Day of the month is already *, we remove strfdh from the last element of the key string
        if currentDom == "*" {
            currentDom = stripChars(strings.Split(key, " ")[len(strings.Split(key, " "))-1], "strfdh")
        }
    }
    return result, currentDom
}

// this function is used to start the scheduler
// parameters: filePath (a string containing the path to a file)
// returns: []StringFormat (an array of cron string formats)
func Execute(filePath string) []StringFormat {
    var stringFormat []StringFormat
    // day of week, day of month and month of the year values are returned from functions in keywords.go
    dowValues := getDowMap()
    domValues := getDomMap()
    monValues := getMonMap()

    // initalize a 2-D string aray and get the yaml passed to the scheduler
    var formattedRules [][]string
    yaml := getYaml(filePath)

    if isScheduleValid(yaml) {
        if yaml == "every minute" {
            var sf StringFormat
            sf.Minute, sf.Hour, sf.Dom, sf.Mon, sf.Dow = "*","*","*","*","*"
            stringFormat = append(stringFormat, sf)
            return stringFormat
        }
        // if the schedule is deemed valid, the string is split on the and & at keywords into rules
        for _, rule := range splitOnKeyword(yaml, "and") {
            formattedRules = append(formattedRules, splitOnKeyword(rule, "at"))
        }
        // the rules are iterated over and converted into the representative cron values
        for i, _ := range formattedRules {
            var sf StringFormat
            key := trimEdges(formattedRules[i][0], "\t \n")
            sf.Dow = getCron(dowValues, key)
            sf.Dom = getCron(domValues, key)
            sf.Mon, sf.Dom = getCronMonth(monValues, sf.Dom, key)
            sf.Hour = splitOnKeyword(formattedRules[i][1], ":")[0]
            if splitOnKeyword(formattedRules[i][1], ":")[1] == "00" {
                sf.Minute = "0"
            } else {
                sf.Minute = splitOnKeyword(formattedRules[i][1], ":")[1]
            }
            stringFormat = append(stringFormat, sf)
        }
    } else {
        // if the schedule isn't valid the the schedule does nothing and alerts this fact
        fmt.Println("Odin cannot recognise the schedule found in your Yaml config file.")
    }
    return stringFormat
}
