package odinlib

import (
    "bytes"
    "net/http"
)

func Log(varType string, desc string, value string, id string, timestamp string) bool {
    str := []byte(varType + "," + desc + "," + value + "," + id + "," + timestamp)
    request, _ := http.NewRequest("POST", "http://localhost:3939/stats/add", bytes.NewBuffer(str))
    client := &http.Client{}
    response, err := client.Do(request)
    defer response.Body.Close()
    if err != nil {
        return false
    }
    return true
}

