package commands

import (
    "fmt"
    "os"
    "log"
    "io/ioutil"
    "net/http"
    "strings"
    "github.com/spf13/cobra"
)

var ExecuteCmd = &cobra.Command{
    Use:   "execute",
    Short: "execute a job created by user",
    Long:  `This subcommand executes a job created by the user`,
    Run: func(cmd *cobra.Command, args []string) {
            executeJob(cmd, args)
    },
}

func init() {
    RootCmd.AddCommand(ExecuteCmd)
    ExecuteCmd.Flags().StringP("file", "f", "", "file (required)")
    ExecuteCmd.MarkFlagRequired("file")
}

func executeJob(cmd *cobra.Command, args []string) {
    name, _:= cmd.Flags().GetString("file")
    contents := readJobFileExecute(name)
    fmt.Println(string(contents))
    resp, _ := http.NewRequest("POST", "localhost:3939/execute", strings.NewReader("/home/odin/go/src/odin/src/odin-cli/job.yml"))
    fmt.Println(resp)
}
func readJobFileExecute(name string) []byte {
    file, err := os.Open(name)
    if err != nil {
        log.Fatal(err)
    }
    bytes, err := ioutil.ReadAll(file)
    defer file.Close()
    return bytes
}
