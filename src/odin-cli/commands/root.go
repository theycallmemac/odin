package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:   "odin",
    Short: "orchestrate your jobs",
    Long: `orchestrate your jobs for periodic execution`,
}

type Config struct {
    Provider ProviderType `yaml:"provider"`
    Job JobType `yaml:"job"`
}

type ProviderType struct {
    Name string `yaml:"name"`
    Version string `yaml:"version"`
}

type JobType struct {
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    Language string `yaml:"language"`
    File string `yaml:"file"`
    Schedule string `yaml:"schedule"`
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

