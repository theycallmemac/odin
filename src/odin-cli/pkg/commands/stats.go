package commands

import (
    "fmt"
    "bytes"
    "io/ioutil"

    "github.com/spf13/cobra"
)

// define the RemoveCmd's metadata and run operation
var StatsCmd = &cobra.Command{
    Use:   "stats",
    Short: "return the stats of an Odin job",
    Long:  `This subcommand will return the stats of an Odin job`,
    Run: func(cmd *cobra.Command, args []string) {
        all, _ := cmd.Flags().GetBool("all")
        id, _:= cmd.Flags().GetString("id")
        port, _:= cmd.Flags().GetString("port")
        if port == "" {
            port = DefaultPort
        }
        if all {
            statsAll(port)
        } else if id != "" {
            statsJob(id, port)
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
    RootCmd.AddCommand(StatsCmd)
    StatsCmd.Flags().BoolVarP(&All, "all", "a", false, "all")
    StatsCmd.Flags().StringP("id", "i", "", "id (required)")
    StatsCmd.Flags().StringP("port", "p", "", "port")
    StatsCmd.MarkFlagRequired("id")
}

// this function is called as a run operation for the StatsCmd to get the stats of a single job
// parameters: id (a string of the required id), port (a string of the port to be used)
// returns: nil
func statsJob(id string, port string) {
    response := makePostRequest(fmt.Sprintf("http://localhost%s/jobs/info/stats", port), bytes.NewBuffer([]byte(id)))
    fmt.Println(response)
}

// this function is called as a run operation for the StatsCmd to get the stats of all jobs
// parameters: port (a string of the port to be used)
// returns: nil
func statsAll(port string) {
    response := makeGetRequest(fmt.Sprintf("http://localhost%s/jobs/info/stats/all", port))
    data, _ := ioutil.ReadAll(response)
    fmt.Println(string(data))
}
