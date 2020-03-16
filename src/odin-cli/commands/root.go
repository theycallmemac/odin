package commands

import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"

    "github.com/spf13/cobra"
)


// ----------------------- INIT COBRA ROOT CMD ---------------------- //
// ------------------------------------------------------------------ //

// define the RootCmd's metadata and run operation
var RootCmd = &cobra.Command{
    Use:   "odin",
    Short: "orchestrate your jobs",
    Long: `orchestrate your jobs for periodic execution`,
}


// this function is called as the run operation for the RootCmd
// parameters: nil
// returns: nil
func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}


// ------------------------- SHARED FUNCTIONS ----------------------- //
// ------------------------------------------------------------------ //

// this function is used to check that a file exists
// parameters: filename (a string containing the path to the file to be checked)
// returns: boolean (returns true if the file exists, false otherwise)
func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

// this function is used to perform a GET request
// parameters: link (a string containing the link to request)
// returns: io.ReadCloser (the body of the response from the request)
func makeGetRequest(link string) io.ReadCloser {
    client := &http.Client{}
    req, reqErr := http.NewRequest("GET", link, nil)
    if reqErr != nil {
        fmt.Println(reqErr)
    }
    res, respErr := client.Do(req)
    if respErr != nil {
        fmt.Println(respErr)
    }
    return res.Body
}

// this function is used to perform a POST request
// parameters: link (a string containing the link to request), data (a buffer of the data to be used in the request)
// returns: string (the body of the response from the request as a string)
func makePostRequest(link string, data *bytes.Buffer) string {
    client := &http.Client{}
    req, _ := http.NewRequest("POST", link, data)
    response, clientErr := client.Do(req)
    if clientErr != nil {
        fmt.Println(clientErr)
    }
    bodyBytes, _ := ioutil.ReadAll(response.Body)
    return string(bodyBytes)
}

// this function is used to perform a PUT request
// parameters: link (a string containing the link to request), data (a buffer of the data to be used in the request)
// returns: string (the body of the response from the request as a string)
func makePutRequest(link string, data *bytes.Buffer) string {
    client := &http.Client{}
    req, _ := http.NewRequest("PUT", link, data)
    response, clientErr := client.Do(req)
    if clientErr != nil {
        fmt.Println(clientErr)
    }
    bodyBytes, _ := ioutil.ReadAll(response.Body)
    return string(bodyBytes)
}


// -------------------------- SHARED STRUCTS ------------------------ //
// ------------------------------------------------------------------ //

// create Config type to tbe used for accessing config information
type Config struct {
    Provider ProviderType `yaml:"provider"`
    Job JobType `yaml:"job"`
}

// create ProviderType type to tbe used for accessing provider information in the config
type ProviderType struct {
    Name string `yaml:"name"`
    Version string `yaml:"version"`
}

// create JobType type to tbe used for accessing job information in the config
type JobType struct {
    ID string `yaml:"id"`
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    Language string `yaml:"language"`
    File string `yaml:"file"`
    Schedule string `yaml:"schedule"`
}

// create NewJob type to tbe used for accessing job information
type NewJob struct {
    ID string `yaml:"id"`
    UID string `yaml:"uid"`
    GID string `yaml:"gid"`
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    Language string `yaml:"language"`
    File string `yaml:"file"`
    Stats string `yaml:"stats"`
    Schedule string `yaml:"schedule"`
}

