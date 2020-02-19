package commands

import (
    "fmt"
    "bytes"
    "github.com/spf13/cobra"
)

var LogCmd = &cobra.Command{
    Use:   "log",
    Short: "show metrics and logs associated with Odin jobs",
    Long:  `This subcommand will show metrics and logs associated with Odin jobs`,
    Run: func(cmd *cobra.Command, args []string) {
            id, _:= cmd.Flags().GetString("id")
            logJob(id)
    },
}

func init() {
    RootCmd.AddCommand(LogCmd)
    LogCmd.Flags().StringP("id", "i", "", "id")
    LogCmd.MarkFlagRequired("id")
}

func logJob(id string) {
    if id != "" {
        response := makePostRequest("http://localhost:3939/jobs/logs", bytes.NewBuffer([]byte(id)))
        fmt.Println(response)
    }
}
