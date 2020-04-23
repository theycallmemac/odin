package commands

import (
    "crypto/rand"
    "fmt"
    "io/ioutil"
    "log"
    "strings"
    "github.com/spf13/cobra"
)

// define the GenerateCmd's metadata and run operation
var GenerateCmd = &cobra.Command{
    Use:   "generate ",
    Short: "creates config files for an Odin job",
    Long:  `This subcommand creates config files for an Odin job`,
    Run: func(cmd *cobra.Command, args []string) {
            generateJob(cmd, args)
    },
}

// add GenereateCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
    RootCmd.AddCommand(GenerateCmd)
    GenerateCmd.Flags().StringP("file", "f", "", "file (required)")
    GenerateCmd.Flags().StringP("lang", "l", "", "lang (required)")
    GenerateCmd.MarkFlagRequired("file")
    GenerateCmd.MarkFlagRequired("lang")
}

// this function is called as the run operation for the GenerateCmd
// parameters: cmd (the definition of *cmd.Command), args (an array of strings passed to the command)
// returns: nil
func generateJob(cmd *cobra.Command, args []string) {
    name, _:= cmd.Flags().GetString("file")
    lang, _:= cmd.Flags().GetString("lang")
    if strings.HasSuffix(name, ".yml") || strings.HasSuffix(name, ".yaml") {
        id := generateId()
        data := []byte("provider:\n  name: 'odin'\n  version: '1.0.0'\njob:\n  id: '" + id + "'\n  name: ''\n  description: ''\n  language: ''\n  file: ''\n  schedule: ''\n\n")
        err := ioutil.WriteFile(name, data, 0644)
        if err != nil {
            panic(err)
        }
        var extension string
        switch lang {
            case "go":
                extension = ".go"
            case "golang":
                extension = ".go"
            case "python":
                extension = ".py"
            case "python3":
                extension = ".py"
        }
        ioutil.WriteFile(strings.Split(name, ".")[0] + extension, []byte(""), 0644)
        fmt.Println("Config and language files generated!")
    }
}

// this function is used to generate a unqiue id
// parameters: nil
// returns: string (the generated id)
func generateId() string {
    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        log.Fatal(err)
    }
    id := fmt.Sprintf("%x%x", b[0:4], b[4:6])
    return id
}
