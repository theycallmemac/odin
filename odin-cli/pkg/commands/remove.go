package commands

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// RemoveCmd is used to define the metadata and run operation for this command
var RemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "removes a user's job by ID",
	Long:    `This subcommand remove a user's job by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		port, _ := cmd.Flags().GetString("port")
		if port == "" {
			port = DefaultPort
		}
		removeJob(id, port)
	},
}

// add RemoveCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
	RootCmd.AddCommand(RemoveCmd)
	RemoveCmd.Flags().StringP("id", "i", "", "id used to specify a job to remove (required)")
	RemoveCmd.Flags().StringP("port", "p", "", "connect to a specific port (default: 3939)")
	RemoveCmd.MarkFlagRequired("id")
}

// this function is called as the run operation for the RemoveCmd
// parameters: id (a string of the required id), port (a string of the port to be used)
// returns: nil
func removeJob(id string, port string) {
	response := makePutRequest(fmt.Sprintf("http://localhost%s/jobs/delete", port), bytes.NewBuffer([]byte(id+" "+fmt.Sprintf("%d", os.Getuid()))))
	fmt.Println(response)
}
