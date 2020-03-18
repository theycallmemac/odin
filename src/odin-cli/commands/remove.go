package commands

import (
    "bytes"
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

// define the RemoveCmd's metadata and run operation
var RemoveCmd = &cobra.Command{
    Use:   "remove",
    Short: "removes a user's job by ID",
    Long:  `This subcommand remove a user's job by ID`,
    Run: func(cmd *cobra.Command, args []string) {
        id, _:= cmd.Flags().GetString("id")
        removeJob(id)
    },
}

// add RemoveCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
    RootCmd.AddCommand(RemoveCmd)
    RemoveCmd.Flags().StringP("id", "i", "", "id")
}

// this function is called as the run operation for the RemoveCmd
// parameters: id (a string of the required id)
// returns: nil
func removeJob(id string) {
    response := makePutRequest("http://localhost:3939/jobs", bytes.NewBuffer([]byte(id + " " + fmt.Sprintf("%d", os.Getuid()))))
    fmt.Println(response)
}
