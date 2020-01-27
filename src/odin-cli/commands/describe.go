package commands

import (
    "fmt"
    "github.com/spf13/cobra"
)

var DescribeCmd = &cobra.Command{
    Use:   "describe",
    Short: "describe a running Odin job",
    Long:  `This subcommand will describe a running Odin job created by the user`,
    Run: func(cmd *cobra.Command, args []string) {
            describeJob()
    },
}

func init() {
    RootCmd.AddCommand(DescribeCmd)
}

func describeJob() {
    fmt.Println("describe job")
}
