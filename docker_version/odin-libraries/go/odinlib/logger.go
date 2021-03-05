package odinlib

import (
	"bytes"
	"net/http"
)

// Log is used to make a Post Request containing collected job data. This is made to the Odin Engine.
func Log(varType string, desc string, value string, id string, timestamp string) bool {
	str := []byte(varType + "," + desc + "," + value + "," + id + "," + timestamp)
	request, _ := http.NewRequest("POST", "http://localhost:3939/stats/add", bytes.NewBuffer(str))
	client := &http.Client{}
	response, err := client.Do(request)
        if err != nil {
            return false
        }
	defer response.Body.Close()
	return true
}
