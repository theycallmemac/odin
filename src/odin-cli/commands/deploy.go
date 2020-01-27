package commands

import (
    "strings"
    "crypto/rand"
    "fmt"
    "os"
    "log"
    "net/http"
    "io"
    "io/ioutil"

    "github.com/spf13/cobra"
    "gopkg.in/yaml.v2"
)

var DeployCmd = &cobra.Command{
    Use:   "deploy",
    Short: "deploy a job created by user",
    Long:  `This subcommand deploys a job created by the user`,
    Run: func(cmd *cobra.Command, args []string) {
            deployJob(cmd, args)
    },
}

func init() {
    RootCmd.AddCommand(DeployCmd)
    DeployCmd.Flags().StringP("file", "f", "", "file (required)")
    DeployCmd.MarkFlagRequired("file")
}

func deployJob(cmd *cobra.Command, args []string) {
    name, _:= cmd.Flags().GetString("file")
    byteArray := readJobFile(name)
    yaml := unmarsharlYaml(byteArray)
    id := generateId()
    fmt.Println(id)
    getScheduleString(name)
    setupJobEnvironment(yaml, name, id)
    // Create directory called /etc/odin/jobs/$id (check it doesnt already exist)
    // If it does exist, gently tell the user
    // If it doesn't exist, create the directory and place the relevant data into it
}

func readJobFile(name string) []byte {
    file, err := os.Open(name)
    if err != nil {
        log.Fatal(err)
    }
    bytes, err := ioutil.ReadAll(file)
    defer file.Close()
    return bytes
}

func unmarsharlYaml(byteArray []byte) Config {
   var cfg Config
    err := yaml.Unmarshal([]byte(byteArray), &cfg)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    return cfg
}

func generateId() string {
    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        log.Fatal(err)
    }
    id := fmt.Sprintf("%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
    return id
}

func ensureDirectory(dir string) bool {
    if  _, err := os.Stat(dir); os.IsNotExist(err) {
        return false
    }
    return true
}

func getScheduleString(name string) {
    code := makePostRequest("http://localhost:3939/schedule", strings.NewReader(name))
    fmt.Println(code)
}

func makePostRequest(link string, data io.Reader) int {
    client := &http.Client{}
    req, _ := http.NewRequest("POST", link, data)
    response, clientErr := client.Do(req)
    if clientErr != nil {
        fmt.Println(clientErr)
    }
    return response.StatusCode
}

func setupJobEnvironment(cfg Config, name string, id string) {
    path := "/etc/odin/jobs/"
    jobPath := path + id
    if ensureDirectory(jobPath) {
        setupJobEnvironment(cfg, name, generateId())
    } else {
        os.MkdirAll(jobPath, 0644)
        input, err := ioutil.ReadFile(cfg.Job.File)
        if err != nil {
            fmt.Println(err)
            return
        }
        ioutil.WriteFile(jobPath + "/" + cfg.Job.File, input, 0644)
        MarshalledCfg, _ := yaml.Marshal(cfg)
        ioutil.WriteFile(jobPath + "/" + name, MarshalledCfg, 0644)
    }
}
