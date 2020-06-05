package commands

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
)

// StatsCmd is used to define the metadata and run operation for this command
var StatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "return the stats of an Odin job",
	Long:  `This subcommand will return the stats of an Odin job`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		port, _ := cmd.Flags().GetString("port")
		if port == "" {
			port = DefaultPort
		}
		statsJob(id, port)
	},
}

// add RemoveCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
	RootCmd.AddCommand(StatsCmd)
	StatsCmd.Flags().StringP("id", "i", "", "id used to view a specific jobs stats (required)")
	StatsCmd.Flags().StringP("port", "p", "", "connect to a specific port (default: 3939)")
	StatsCmd.MarkFlagRequired("id")
}

// this function is called as a run operation for the StatsCmd to get the stats of a single job
// parameters: id (a string of the required id), port (a string of the port to be used)
// returns: nil
func statsJob(id string, port string) {
	response := makePostRequest(fmt.Sprintf("http://localhost%s/stats/get", port), bytes.NewBuffer([]byte(id)))
	fmt.Println(response)
}
