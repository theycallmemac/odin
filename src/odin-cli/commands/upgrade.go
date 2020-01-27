package commands

import (
    "fmt"
    "github.com/spf13/cobra"
)

var UpgradeCmd = &cobra.Command{
    Use:   "upgrade",
    Short: "upgrade Odin to the latest release",
    Long:  `This subcommand will upgrade Odin to the latest release`,
    Run: func(cmd *cobra.Command, args []string) {
            upgradeOdin()
    },
}

func init() {
    RootCmd.AddCommand(UpgradeCmd)
}

func upgradeOdin() {
    fmt.Println("upgrade odin")
}
