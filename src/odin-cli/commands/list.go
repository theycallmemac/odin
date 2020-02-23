package commands

import (
    "fmt"
    "io/ioutil"
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
    response := makeGetRequest("http://localhost:3939/jobs")
    body, _ := ioutil.ReadAll(response)
    fmt.Println(string(body))
}
