package commands

import (
    "bytes"
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

// define the ModifyCmd's metadata and run operation
var ModifyCmd = &cobra.Command{
    Use:   "modify",
    Short: "change details about an Odin job in-place",
    Long:  `This subcommand change details about an Odin job in-place`,
    Run: func(cmd *cobra.Command, args []string) {
            id, _:= cmd.Flags().GetString("id")
            name, _:= cmd.Flags().GetString("name")
            desc, _:= cmd.Flags().GetString("description")
            schedule, _:= cmd.Flags().GetString("schedule")
            port, _:= cmd.Flags().GetString("port")
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
    ModifyCmd.Flags().StringP("id", "i", "", "id")
    ModifyCmd.MarkFlagRequired("id")
    ModifyCmd.Flags().StringP("name", "n", "", "name")
    ModifyCmd.Flags().StringP("description", "d", "", "description")
    ModifyCmd.Flags().StringP("schedule", "s", "", "schedule")
    ModifyCmd.Flags().StringP("port", "p", "", "port")
}

// this function is called as the run operation for the ModifyCmd
// parameters: id (a string of the required id), name (a string to change the job name), desc (a string to change the job description), schedule (a string to change the job schedule), port (a string of the port to be used)
// returns: nil
func modifyJob(id string, name string, desc string, schedule string, port string) {
    if id != "" && name == "" && desc == "" && schedule == "" {
        fmt.Println("Please specify which field you want to modify in job " + id + "\n")
    } else {
        response := makePutRequest(fmt.Sprintf("http://localhost%s/jobs/info/", port), bytes.NewBuffer([]byte(id + " " + name + " " + desc + " " + schedule + " " + fmt.Sprintf("%d", os.Getuid()))))
        fmt.Println(response)
    }
}
