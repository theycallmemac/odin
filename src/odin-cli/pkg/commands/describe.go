package commands

import (
    "bytes"
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

// define the DescribeCmd's metadata and run operation
var DescribeCmd = &cobra.Command{
    Use:   "describe",
    Short: "describe a running Odin job",
    Long:  `This subcommand will describe a running Odin job created by the user`,
    Run: func(cmd *cobra.Command, args []string) {
            id, _:= cmd.Flags().GetString("id")
            describeJob(id)
    },
}

// add DescribeCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
    RootCmd.AddCommand(DescribeCmd)
    DescribeCmd.Flags().StringP("id", "i", "", "id (required)")
    DescribeCmd.MarkFlagRequired("id")
}

// this function is called as the run operation for the DescribeCmd
// parameters: id (a string of the required id)
// returns: nil
func describeJob(id string) {
    response := makePostRequest("http://localhost:3939/jobs/info/description", bytes.NewBuffer([]byte(id + " " + fmt.Sprintf("%d", os.Getuid()))))
    fmt.Println(response)
}
