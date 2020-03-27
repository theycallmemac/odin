package commands

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "os/user"
    "strconv"
    "syscall"

    "github.com/spf13/cobra"
    "gopkg.in/yaml.v2"
)

// define the DeployCmd's metadata and run operation
var DeployCmd = &cobra.Command{
    Use:   "deploy",
    Short: "deploy a job created by user",
    Long:  `This subcommand deploys a job created by the user`,
    Run: func(cmd *cobra.Command, args []string) {
            deployJob(cmd, args)
    },
}

// add DeployCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
    RootCmd.AddCommand(DeployCmd)
    DeployCmd.Flags().StringP("file", "f", "", "file (required)")
    DeployCmd.MarkFlagRequired("file")
}

// this function is called as the run operation for the DeployCmd
// parameters: cmd (the definition of *cmd.Command), args (an array of strings passed to the command)
// returns: nil
func deployJob(cmd *cobra.Command, args []string) {
    name, _:= cmd.Flags().GetString("file")
    yaml := unmarsharlYaml(readJobFile(name))
    currentDir, _ := os.Getwd()
    var job NewJob
    job.ID = yaml.Job.ID
    job.UID = fmt.Sprint(syscall.Getuid())
    group, _ := user.LookupGroup("odin")
    gid, _ := strconv.Atoi(group.Gid)
    job.GID = strconv.Itoa(gid)
    job.Name = yaml.Job.Name
    job.Description =  yaml.Job.Description
    job.Language = yaml.Job.Language
    job.File = currentDir + "/" + yaml.Job.File
    job.Schedule = getScheduleString(name)
    jobJSON, _ := json.Marshal(job)
    body := makePostRequest("http://localhost:3939/jobs", bytes.NewBuffer(jobJSON))
    fmt.Println(body)
}

// this function is used to read a file
// parameters: name (a string containing the path to a file)
// returns: []byte (an array of bytes containing the contents of the file)
func readJobFile(name string) []byte {
    file, err := os.Open(name)
    if err != nil {
        log.Fatal(err)
    }
    bytes, err := ioutil.ReadAll(file)
    defer file.Close()
    return bytes
}

// this function is used to unmarshal YAML
// parameters: byteArray (an array of bytes representing the contents of a file)
// returns: Config (a struct form of the YAML)
func unmarsharlYaml(byteArray []byte) Config {
   var cfg Config
    err := yaml.Unmarshal([]byte(byteArray), &cfg)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    return cfg
}

// this function is used to check if a directory exists
// parameters: dir (a string containing the path to the checked directory)
// returns: boolean (true is no error in checking the existence of the directory, false if otherwise)
func ensureDirectory(dir string) bool {
    if  _, err := os.Stat(dir); os.IsNotExist(err) {
        return false
    }
    return true
}

// this function is used to get the schedule string using the path to the file
// parameters: name (a string containing the path to a file)
// returns: ss (the generated schedule string)
func getScheduleString(name string) string {
    dir, _ := os.Getwd()
    absPath := []byte(dir + "/" + name)
    ss := makePostRequest("http://localhost:3939/schedule", bytes.NewBuffer(absPath))
    return ss
}

