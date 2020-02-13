package commands

import (
    "bytes"
    "fmt"
    "github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
    Use:   "remove",
    Short: "removes a user's job by ID",
    Long:  `This subcommand remove a user's job by ID`,
    Run: func(cmd *cobra.Command, args []string) {
        id, _:= cmd.Flags().GetString("id")
        removeJob(id)
    },
}

func init() {
    RootCmd.AddCommand(RemoveCmd)
    RemoveCmd.Flags().StringP("id", "i", "", "id")
}

func removeJob(id string) {
    response := makePutRequest("http://localhost:3939/jobs", bytes.NewBuffer([]byte(id)))
    fmt.Println(response)
}
