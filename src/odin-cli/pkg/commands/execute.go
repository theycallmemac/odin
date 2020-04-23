package commands

import (
    "bytes"
    "fmt"
    "os"
    "log"
    "io/ioutil"

    "github.com/spf13/cobra"
)

// define the ExecuteCmd's metadata and run operation
var ExecuteCmd = &cobra.Command{
    Use:   "execute",
    Short: "execute a job created by user",
    Long:  `This subcommand executes a job created by the user`,
    Run: func(cmd *cobra.Command, args []string) {
        port, _:= cmd.Flags().GetString("port")
        if port == "" {
            port = DefaultPort
        }
        executeJob(cmd, args, port)
    },
}

// add ExecuteCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
    RootCmd.AddCommand(ExecuteCmd)
    ExecuteCmd.Flags().StringP("file", "f", "", "file used to specify the job to execute (required)")
    ExecuteCmd.Flags().StringP("port", "p", "", "connect to a specific port (default: 3939)")
    ExecuteCmd.MarkFlagRequired("file")
}

// this function is called as the run operation for the ExecuteCmd
// parameters: cmd (the definition of *cmd.Command), args (an array of strings passed to the command), port (a string of the port to be used)
// returns: nil
func executeJob(cmd *cobra.Command, args []string, port string) {
    name, _:= cmd.Flags().GetString("file")
    contents := readJobFileExecute(name)
    fmt.Println(string(contents))
    dir, _ := os.Getwd()
    resp := makePostRequest(fmt.Sprintf("http://localhost%s/execute/yaml", port), bytes.NewBuffer([]byte(dir+"/"+name)))
    fmt.Println(resp)
    fmt.Println("Executed successfully!")
}

// this function is used to read a file
// parameters: name (a string containing the path to a file)
// returns: []byte (an array of bytes containing the contents of the file)
func readJobFileExecute(name string) []byte {
    file, err := os.Open(name)
    if err != nil {
        log.Fatal(err)
    }
    bytes, err := ioutil.ReadAll(file)
    defer file.Close()
    return bytes
}
