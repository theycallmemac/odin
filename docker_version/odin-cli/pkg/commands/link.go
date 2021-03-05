package commands

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// LinkCmd is used to define metadata and run operation for this command
var LinkCmd = &cobra.Command{
	Use:   "link",
	Short: "links the user's current Odin jobs",
	Long:  `This subcommand links the user's current Odin jobs`,
	Run: func(cmd *cobra.Command, args []string) {
		from, _ := cmd.Flags().GetString("from")
		to, _ := cmd.Flags().GetString("to")
		port, _ := cmd.Flags().GetString("port")
		if port == "" {
			port = DefaultPort
		}
		linkJob(from, to, port)
	},
}

// add LinkCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
	RootCmd.AddCommand(LinkCmd)
	LinkCmd.Flags().StringP("from", "f", "", "from")
	LinkCmd.Flags().StringP("to", "t", "", "to")
	LinkCmd.Flags().StringP("port", "p", "", "port")
	LinkCmd.MarkFlagRequired("from")
	LinkCmd.MarkFlagRequired("to")
}

// this function is called as the run operation for the LinkCmd
// parameters: port (a string of the port to be used)
// returns: nil
func linkJob(from, to, port string) {
	uid := strconv.Itoa(os.Getuid())
	response := makePostRequest(fmt.Sprintf("http://localhost%s/links/add", port), bytes.NewBuffer([]byte(fmt.Sprintf("%s", from+"_"+to+"_"+uid))))
	fmt.Print(response)
}
