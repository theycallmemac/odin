package commands

import (
    "strings"
    "context"
    "crypto/rand"
    "fmt"
    "os"
    "log"
    "net/http"
    "io"
    "io/ioutil"

    "github.com/spf13/cobra"
    "gopkg.in/yaml.v2"
    "go.mongodb.org/mongo-driver/mongo/readpref"
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
    ss := getScheduleString(name)
    jobPath := setupJobEnvironment(yaml, name, id)
    if jobPath == "" {
        os.Exit(2)
    }
    c := getMongoClient()
    err := c.Ping(context.Background(), readpref.Primary())
    if err != nil {
	    log.Fatal("Couldn't connect to the database", err)
    } else {
	    log.Println("Connected!")
    }
    job := NewJob{ID: id, Name: yaml.Job.Name, Description: yaml.Job.Description, Language: yaml.Job.Language, File: jobPath + "/" + yaml.Job.File, ScheduleString: ss}
    inserted := insertIntoMongo(c, job)
    fmt.Println(inserted)
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

func getScheduleString(name string) string {
    dir, _ := os.Getwd()
    absPath := dir + "/" + name
    ss := makePostRequest("http://localhost:3939/schedule", strings.NewReader(absPath))
    return ss
}

func makePostRequest(link string, data io.Reader) string {
    client := &http.Client{}
    req, _ := http.NewRequest("POST", link, data)
    response, clientErr := client.Do(req)
    if clientErr != nil {
        fmt.Println(clientErr)
    }
    bodyBytes, _ := ioutil.ReadAll(response.Body)
    return string(bodyBytes)
}

func setupJobEnvironment(cfg Config, name string, id string) string {
    path := "/etc/odin/jobs/"
    jobPath := path + id
    if ensureDirectory(jobPath) {
        setupJobEnvironment(cfg, name, generateId())
    } else {
        os.MkdirAll(jobPath, 0644)
        input, err := ioutil.ReadFile(cfg.Job.File)
        if err != nil {
            fmt.Println(err)
            return ""
        }
        ioutil.WriteFile(jobPath + "/" + cfg.Job.File, input, 0644)
        MarshalledCfg, _ := yaml.Marshal(cfg)
        ioutil.WriteFile(jobPath + "/" + name, MarshalledCfg, 0644)
    }
    return jobPath
}
