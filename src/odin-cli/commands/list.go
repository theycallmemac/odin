package commands

import (
    "fmt"
    "bytes"
    "os"
    "github.com/spf13/cobra"
)

// define the ListCmd's metadata and run operation
var ListCmd = &cobra.Command{
    Use:   "list",
    Short: "lists the user's current Odin jobs",
    Long:  `This subcommand lists the user's current Odin jobs`,
    Run: func(cmd *cobra.Command, args []string) {
            listJob()
    },
}

// add ListCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
    RootCmd.AddCommand(ListCmd)
}

// this function is called as the run operation for the ListCmd
// parameters: nil
// returns: nil
func listJob() {
    response := makePostRequest("http://localhost:3939/jobs/list", bytes.NewBuffer([]byte(fmt.Sprintf("%d", os.Getuid()))))
    fmt.Print(response)
}
