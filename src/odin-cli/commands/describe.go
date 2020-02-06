package commands

import (
    "fmt"
    "bytes"
    //"io/ioutil"

    "github.com/spf13/cobra"
)

var DescribeCmd = &cobra.Command{
    Use:   "describe",
    Short: "describe a running Odin job",
    Long:  `This subcommand will describe a running Odin job created by the user`,
    Run: func(cmd *cobra.Command, args []string) {
            id, _:= cmd.Flags().GetString("id")
            describeJob(id)
    },
}

func init() {
    RootCmd.AddCommand(DescribeCmd)
    DescribeCmd.Flags().StringP("id", "i", "", "id (required)")
    DescribeCmd.MarkFlagRequired("id")
}

func describeJob(id string) {
    response := makePostRequest("http://localhost:3939/jobs/info/description", bytes.NewBuffer([]byte(id)))
    fmt.Println(response)
}
