package commands

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// ListCmd is used to define the metadata and run operation for this command
var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "lists the user's current Odin jobs",
	Long:    `This subcommand lists the user's current Odin jobs`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		if port == "" {
			port = DefaultPort
		}
		listJob(port)
	},
}

// add ListCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
	RootCmd.AddCommand(ListCmd)
	ListCmd.Flags().StringP("port", "p", "", "connect to a specific port (default: 3939)")
}

// this function is called as the run operation for the ListCmd
// parameters: port (a string of the port to be used)
// returns: nil
func listJob(port string) {
	response := makePostRequest(fmt.Sprintf("http://localhost%s/jobs/list", port), bytes.NewBuffer([]byte(fmt.Sprintf("%d", os.Getuid()))))
	fmt.Print(response)
}
