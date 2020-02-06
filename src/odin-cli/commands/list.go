package commands

import (
    "fmt"
    "io/ioutil"
    "github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
    Use:   "list",
    Short: "lists the user's current Odin jobs",
    Long:  `This subcommand lists the user's current Odin jobs`,
    Run: func(cmd *cobra.Command, args []string) {
            listJob()
    },
}

func init() {
    RootCmd.AddCommand(ListCmd)
}

func listJob() {
    response := makeGetRequest("http://localhost:3939/jobs")
    body, _ := ioutil.ReadAll(response)
    fmt.Println(string(body))
}
