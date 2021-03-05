package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// DeployCmd is used to define the metadata and run operation for this command
var DeployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"add", "dep"},
	Short:   "deploy a job created by user",
	Long:    `This subcommand deploys a job created by the user`,
	Run: func(cmd *cobra.Command, args []string) {
		deployJob(cmd, args)
	},
}

// add DeployCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
	RootCmd.AddCommand(DeployCmd)
	DeployCmd.Flags().StringP("file", "f", "", "file used to specify the job to deploy (required)")
	DeployCmd.Flags().StringP("port", "p", "", "connect to a specific port (default: 3939)")
	DeployCmd.MarkFlagRequired("file")
}

// this function is called as the run operation for the DeployCmd
// parameters: cmd (the definition of *cmd.Command), args (an array of strings passed to the command)
// returns: nil
func deployJob(cmd *cobra.Command, args []string) {
	port, _ := cmd.Flags().GetString("port")
	if port == "" {
		port = DefaultPort
	}
	name, _ := cmd.Flags().GetString("file")
	yaml := unmarsharlYaml(readJobFile(name))
        if !semanticCheck(yaml) {
            os.Exit(2)
        }
	currentDir, _ := os.Getwd()
	var job NewJob
	job.ID = yaml.Job.ID
	job.UID = fmt.Sprint(syscall.Getuid())
	group, _ := user.LookupGroup("odin")
	gid, _ := strconv.Atoi(group.Gid)
	job.GID = strconv.Itoa(gid)
	job.Name = yaml.Job.Name
	job.Description = yaml.Job.Description
	job.File = currentDir + "/" + yaml.Job.File
	if yaml.Job.Language == "go" {
		job.Language = yaml.Job.Language
		cmd := exec.Command(job.Language, "build", job.File)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		_, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		job.File = job.File[:len(job.File)-3]
	} else {
		job.Language = yaml.Job.Language
	}
	job.Schedule = getScheduleString(name, port)
	jobJSON, _ := json.Marshal(job)
	body := makePostRequest(fmt.Sprintf("http://localhost%s/jobs/add", port), bytes.NewBuffer(jobJSON))
	fmt.Println(body)
}

// this function is used to check the Yaml Config of a job before a deployment
// parameters: yaml (a Config type containing the file fields)
// returns: bool (a boolean indicating the success the check process)
func semanticCheck(yaml Config) bool {
    var status = true
    if yaml.Provider.Version == "2.0.0" {
        if yaml.Job.Name == "" {
            printField("\n - Name")
            status = false
        }
        if yaml.Job.Description == "" {
            printField("\n - Description")
            status = false
        }
        if yaml.Job.Schedule == "" {
            printField("\n - Schedule")
            status = false
        }
    }
    return status
}

// this function is used to print a specific field
// parameters: field (a string containing the field to print)
// returns: nil 
func printField(field string) {
    if field != "" {
        fmt.Println(field + " field is missing from Yaml config.")
    }
}

// this function is used to read a file
// parameters: name (a string containing the path to a file)
// returns: []byte (an array of bytes containing the contents of the file)
func readJobFile(name string) []byte {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	bytes, _ := ioutil.ReadAll(file)
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
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

// this function is used to get the schedule string using the path to the file
// parameters: name (a string containing the path to a file), port (a string of the port to be used)
// returns: ss (the generated schedule string)
func getScheduleString(name string, port string) string {
	dir, _ := os.Getwd()
	absPath := []byte(dir + "/" + name)
	ss := makePostRequest(fmt.Sprintf("http://localhost%s/schedule", port), bytes.NewBuffer(absPath))
	return ss
}
