package commands

import (
    "fmt"
    "io/ioutil"
    "strings"
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
    GenerateCmd.Flags().StringP("file", "f", "", "file (required)")
    GenerateCmd.MarkFlagRequired("file")
}

func generateJob(cmd *cobra.Command, args []string) {
    name, _:= cmd.Flags().GetString("file")
    if strings.HasSuffix(name, ".yml") ||strings.HasSuffix(name, ".yaml") {
        data := []byte("provider:\n  name: 'odin'\n  version: '1.0.0'\njob:\n  name: ''\n  description: ''\n  language: ''\n  file: ''\n  schedule: ''\n\n")
        err := ioutil.WriteFile(name, data, 0644)
        if err != nil {
            panic(err)
        }
        fmt.Println(name + " config generated!")
    }
}
