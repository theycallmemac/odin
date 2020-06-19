package commands

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ModifyCmd is used to define the metadata and run operation for this command
var ModifyCmd = &cobra.Command{
	Use:     "modify",
	Aliases: []string{"mod"},
	Short:   "change details about an Odin job in-place",
	Long:    `This subcommand change details about an Odin job in-place`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")
		desc, _ := cmd.Flags().GetString("description")
		schedule, _ := cmd.Flags().GetString("schedule")
		port, _ := cmd.Flags().GetString("port")
		if port == "" {
			port = DefaultPort
		}
		modifyJob(id, name, desc, schedule, port)
	},
}

// add ModifyCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
	RootCmd.AddCommand(ModifyCmd)
	ModifyCmd.Flags().StringP("id", "i", "", "id used to specify a job to modify (required)")
	ModifyCmd.Flags().StringP("name", "n", "", "change the current name")
	ModifyCmd.Flags().StringP("description", "d", "", "change the current description")
	ModifyCmd.Flags().StringP("schedule", "s", "", "change the current schedule")
	ModifyCmd.Flags().StringP("port", "p", "", "connect to a specific port (default: 3939)")
	ModifyCmd.MarkFlagRequired("id")
}

// this function is called as the run operation for the ModifyCmd
// parameters: id (a string of the required id), name (a string to change the job name), desc (a string to change the job description), schedule (a string to change the job schedule), port (a string of the port to be used)
// returns: nil
func modifyJob(id string, name string, desc string, schedule string, port string) {
	if id != "" && name == "" && desc == "" && schedule == "" {
		fmt.Println("Please specify which field you want to modify in job " + id + "\n")
	} else {
		response := makePutRequest(fmt.Sprintf("http://localhost%s/jobs/info/update", port), bytes.NewBuffer([]byte(id+"_"+name+"_"+desc+"_"+schedule+"_"+fmt.Sprintf("%d", os.Getuid()))))
		fmt.Println(response)
	}
}
