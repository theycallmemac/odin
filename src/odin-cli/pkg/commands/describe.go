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
        port, _:= cmd.Flags().GetString("port")
        if port == "" {
            port = DefaultPort
        }
        describeJob(id, port)
    },
}

// add DescribeCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
    RootCmd.AddCommand(DescribeCmd)
    DescribeCmd.Flags().StringP("id", "i", "", "id (required)")
    DescribeCmd.Flags().StringP("port", "p", "", "port")
    DescribeCmd.MarkFlagRequired("id")
}

// this function is called as the run operation for the DescribeCmd
// parameters: id (a string of the required id), port (a string of the port to be used)
// returns: nil
func describeJob(id string, port string) {
    response := makePostRequest(fmt.Sprintf("http://localhost%s/jobs/info/description", port), bytes.NewBuffer([]byte(id + "_" + fmt.Sprintf("%d", os.Getuid()))))
    fmt.Println(response)
}
