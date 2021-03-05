package commands

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// DescribeCmd is used to define the metadata and run operation for this command
var DescribeCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"desc"},
	Short:   "describe a running Odin job",
	Long:    `This subcommand will describe a running Odin job created by the user`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		port, _ := cmd.Flags().GetString("port")
		if port == "" {
			port = DefaultPort
		}
		describeJob(id, port)
	},
}

// add DescribeCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
	RootCmd.AddCommand(DescribeCmd)
	DescribeCmd.Flags().StringP("id", "i", "", "id used to specify a job to describe (required)")
	DescribeCmd.Flags().StringP("port", "p", "", "connect to a specific port (default: 3939)")
	DescribeCmd.MarkFlagRequired("id")
}

// this function is called as the run operation for the DescribeCmd
// parameters: id (a string of the required id), port (a string of the port to be used)
// returns: nil
func describeJob(id string, port string) {
	response := makePostRequest(fmt.Sprintf("http://localhost%s/jobs/info/description", port), bytes.NewBuffer([]byte(id+"_"+fmt.Sprintf("%d", os.Getuid()))))
	fmt.Println(response)
}
