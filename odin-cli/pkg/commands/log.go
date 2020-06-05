package commands

import (
	"bytes"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
)

// LogCmd is used to define the metadata and run operation for this command
var LogCmd = &cobra.Command{
	Use:   "log",
	Short: "show metrics and logs associated with Odin jobs",
	Long:  `This subcommand will show metrics and logs associated with Odin jobs`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		port, _ := cmd.Flags().GetString("port")
		if port == "" {
			port = DefaultPort
		}
		logJob(id, port)
	},
}

// add LogCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
	RootCmd.AddCommand(LogCmd)
	LogCmd.Flags().StringP("id", "i", "", "id used to view a specific jobs logs (required)")
	LogCmd.Flags().StringP("port", "p", "", "connect to a specific port (default: 3939)")
	LogCmd.MarkFlagRequired("id")
}

// this function is called as the run operation for the LogCmd
// parameters: id (a string of the required id), port (a string of the port to be used)
// returns: nil
func logJob(id string, port string) {
	var gid int
	if id != "" {
		fileInfo, _ := os.Stat("/etc/odin/jobs/" + id)
		group, err := user.LookupGroup("odin")
		if err != nil {
			fmt.Println("User is not in the `odin` group")
			os.Exit(2)
		} else {
			gid, _ = strconv.Atoi(group.Gid)
		}
		if gid == int(fileInfo.Sys().(*syscall.Stat_t).Gid) {
			response := makePostRequest(fmt.Sprintf("http://localhost%s/jobs/logs", port), bytes.NewBuffer([]byte(id)))
			fmt.Println(response)
		} else {
			fmt.Println("Cannot access the logs for job " + id + "\n")
		}
	}
}
