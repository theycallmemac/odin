package commands

import (
    "fmt"
    "bytes"
    "io/ioutil"

    "github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
    Use:   "status",
    Short: "return the status of an Odin job",
    Long:  `This subcommand will return the status of an Odin job`,
    Run: func(cmd *cobra.Command, args []string) {
            all, _ := cmd.Flags().GetBool("all")
            id, _:= cmd.Flags().GetString("id")
            if all {
                statusAll()
            } else if id != "" {
                statusJob(id)
            } else {
                cmd.Help()
            }
    },
}

func init() {
    var All bool
    RootCmd.AddCommand(StatusCmd)
    StatusCmd.Flags().BoolVarP(&All, "all", "a", false, "all")
    StatusCmd.Flags().StringP("id", "i", "", "id")
}

func statusJob(id string) {
    response := makePostRequest("http://localhost:3939/jobs/info/status", bytes.NewBuffer([]byte(id)))
    fmt.Println(response)
}

func statusAll() {
    response := makeGetRequest("http://localhost:3939/jobs/info/status/all")
    data, _ := ioutil.ReadAll(response)
    fmt.Println(string(data))
}
