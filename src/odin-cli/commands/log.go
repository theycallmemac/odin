package commands

import (
    "fmt"
    "github.com/spf13/cobra"
)

var LogCmd = &cobra.Command{
    Use:   "log",
    Short: "show metrics and logs associated with Odin jobs",
    Long:  `This subcommand will show metrics and logs associated with Odin jobs`,
    Run: func(cmd *cobra.Command, args []string) {
            logJob()
    },
}

func init() {
    RootCmd.AddCommand(LogCmd)
}

func logJob() {
    fmt.Println("log job")
}
