package scheduler

import (
    "fmt"
    "strings"
)

type StringFormat struct {
    Minute string
    Hour string
    Dom string
    Mon string
    Dow string
}

func isTimeValid(time string, matchMe string, results []string) ([]string, float64) {
    var addOn float64 = 0
    timeSplit := splitOnKeyword(time, ":")
    for _, ts := range timeSplit {
        if compare(ts, matchMe) {
            results = append(results, "0")
            addOn = 0.5
        }
    }
    return results, addOn
}

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
                if compare(valid, word) {
                    results = append(results, "0")
                    break
                }
            }
        }
    }
    return len(results) == len(words)+int(timeSplitAddOn)
}

func getCron(values map[string]string, key string) string {
    var result string
    if values[key] == "" {
        result = "*"
    } else {
        result = values[key]
    }
    return result
}

func getCronMonth(values map[string]string, currentDom string, key string) (string, string) {
    var result string
    newKey := strings.Join(strings.Split(key, " ")[0:2]," ")
    if values[newKey] == "" {
        result = "*"
    } else {
        result = values[newKey]
        if currentDom == "*" {
            currentDom = stripChars(strings.Split(key, " ")[len(strings.Split(key, " "))-1], "strfdh")
        }
    }
    return result, currentDom
}

func Execute(filePath string) []StringFormat {
    var asf []StringFormat
    dowValues := getDowMap()
    domValues := getDomMap()
    monValues := getMonMap()
    var formattedRules [][]string
    yaml := getYaml(filePath)

    if isScheduleValid(yaml) {
        for _, rule := range splitOnKeyword(yaml, "and") {
            formattedRules = append(formattedRules, splitOnKeyword(rule, "at"))
        }
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
            asf = append(asf, sf)
        }
    } else {
        fmt.Println("Odin cannot recognise the schedule found in your Yaml config file.")
    }
    return asf
}
