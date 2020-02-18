package commands

import (
    "fmt"
    "bytes"

    "github.com/spf13/cobra"
)

var ModifyCmd = &cobra.Command{
    Use:   "modify",
    Short: "change details about an Odin job in-place",
    Long:  `This subcommand change details about an Odin job in-place`,
    Run: func(cmd *cobra.Command, args []string) {
            id, _:= cmd.Flags().GetString("id")
            name, _:= cmd.Flags().GetString("name")
            desc, _:= cmd.Flags().GetString("description")
            schedule, _:= cmd.Flags().GetString("schedule")
            modifyJob(id, name, desc, schedule)
    },
}

func init() {
    RootCmd.AddCommand(ModifyCmd)
    ModifyCmd.Flags().StringP("id", "i", "", "id")
    ModifyCmd.MarkFlagRequired("id")
    ModifyCmd.Flags().StringP("name", "n", "", "name")
    ModifyCmd.Flags().StringP("description", "d", "", "description")
    ModifyCmd.Flags().StringP("schedule", "s", "", "schedule")
}

func modifyJob(id string, name string, desc string, schedule string) {
    response := makePutRequest("http://localhost:3939/jobs/info/", bytes.NewBuffer([]byte(id + "," + name + "," + desc + "," + schedule)))
    fmt.Println("modify job", response)
}
