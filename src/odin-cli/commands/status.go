package commands

import (
    "fmt"
    "github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
    Use:   "modify",
    Short: "return the status of an Odin job",
    Long:  `This subcommand will return the status of an Odin job`,
    Run: func(cmd *cobra.Command, args []string) {
            statusJob()
    },
}

func init() {
    RootCmd.AddCommand(StatusCmd)
}

func statusJob() {
    fmt.Println("status job")
}
