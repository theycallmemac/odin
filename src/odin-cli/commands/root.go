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
var RootCmd = &cobra.Command{
    Use:   "odin",
    Short: "orchestrate your jobs",
    Long: `orchestrate your jobs for periodic execution`,
}

func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}



// ------------------------- SHARED FUNCTIONS ----------------------- //
// ------------------------------------------------------------------ //
func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

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
type Config struct {
    Provider ProviderType `yaml:"provider"`
    Job JobType `yaml:"job"`
}

type ProviderType struct {
    Name string `yaml:"name"`
    Version string `yaml:"version"`
}

type JobType struct {
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    Language string `yaml:"language"`
    File string `yaml:"file"`
    Schedule string `yaml:"schedule"`
}


type NewJob struct {
    ID string `yaml:"id"`
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    Language string `yaml:"language"`
    File string `yaml:"file"`
    Status string `yaml:"status"`
    Schedule string `yaml:"schedule"`
}


