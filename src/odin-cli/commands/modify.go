package commands

import (
    "fmt"
    "github.com/spf13/cobra"
)

var ModifyCmd = &cobra.Command{
    Use:   "modify",
    Short: "change details about an Odin job in-place",
    Long:  `This subcommand change details about an Odin job in-place`,
    Run: func(cmd *cobra.Command, args []string) {
            modifyJob()
    },
}

func init() {
    RootCmd.AddCommand(ModifyCmd)
}

func modifyJob() {
    fmt.Println("modify job")
}
