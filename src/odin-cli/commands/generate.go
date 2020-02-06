package commands

import (
    "fmt"
    "github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
    Use:   "generate ",
    Short: "creates config files for an Odin job",
    Long:  `This subcommand creates config files for an Odin job`,
    Run: func(cmd *cobra.Command, args []string) {
            generateJob(cmd, args)
    },
}

func init() {
    RootCmd.AddCommand(GenerateCmd)
}

func generateJob(cmd *cobra.Command, args []string) {
    fmt.Println("generate job")
}
