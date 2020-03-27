package commands

import (
    "fmt"
    "bytes"
    "io/ioutil"

    "github.com/spf13/cobra"
)

// define the RemoveCmd's metadata and run operation
var statsCmd = &cobra.Command{
    Use:   "stats",
    Short: "return the stats of an Odin job",
    Long:  `This subcommand will return the stats of an Odin job`,
    Run: func(cmd *cobra.Command, args []string) {
        all, _ := cmd.Flags().GetBool("all")
        id, _:= cmd.Flags().GetString("id")
        if all {
            statsAll()
        } else if id != "" {
            statsJob(id)
        } else {
            cmd.Help()
        }
    },
}

// add RemoveCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
    var All bool
    RootCmd.AddCommand(statsCmd)
    statsCmd.Flags().BoolVarP(&All, "all", "a", false, "all")
    statsCmd.Flags().StringP("id", "i", "", "id")
}

// this function is called as a run operation for the statsCmd to get the stats of a single job
// parameters: id (a string of the required id)
// returns: nil
func statsJob(id string) {
    response := makePostRequest("http://localhost:3939/jobs/info/stats", bytes.NewBuffer([]byte(id)))
    fmt.Println(response)
}

// this function is called as a run operation for the statsCmd to get the stats of all jobs
// parameters: nil
// returns: nil
func statsAll() {
    response := makeGetRequest("http://localhost:3939/jobs/info/stats/all")
    data, _ := ioutil.ReadAll(response)
    fmt.Println(string(data))
}
