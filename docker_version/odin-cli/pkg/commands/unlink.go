package commands

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// UnlinkCmd is used to define the metadata and run operation for this command
var UnlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "unlinks the user's current Odin jobs",
	Long:  `This subcommand unlinks the user's current Odin jobs`,
	Run: func(cmd *cobra.Command, args []string) {
		from, _ := cmd.Flags().GetString("from")
		to, _ := cmd.Flags().GetString("to")
		port, _ := cmd.Flags().GetString("port")
		if port == "" {
			port = DefaultPort
		}
		unlinkJob(from, to, port)
	},
}

// add UnlinkCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
	RootCmd.AddCommand(UnlinkCmd)
	UnlinkCmd.Flags().StringP("from", "f", "", "from")
	UnlinkCmd.Flags().StringP("to", "t", "", "t")
	UnlinkCmd.Flags().StringP("port", "p", "", "port")
	UnlinkCmd.MarkFlagRequired("from")
	UnlinkCmd.MarkFlagRequired("to")
}

// this function is called as the run operation for the UnlinkCmd
// parameters: port (a string of the port to be used)
// returns: nil
func unlinkJob(from, to, port string) {
	uid := strconv.Itoa(os.Getuid())
	response := makePostRequest(fmt.Sprintf("http://localhost%s/links/delete", port), bytes.NewBuffer([]byte(fmt.Sprintf("%s", from+"_"+to+"_"+uid))))
	fmt.Print(response)
}
