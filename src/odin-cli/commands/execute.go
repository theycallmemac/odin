package commands

import (
    "bytes"
    "fmt"
    "os"
    "log"
    "io/ioutil"
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
    dir, _ := os.Getwd()
    resp := makePostRequest("http://localhost:3939/execute/yaml", bytes.NewBuffer([]byte(dir+"/"+name)))
    fmt.Println(resp)
    fmt.Println("Executed successfully!")
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
