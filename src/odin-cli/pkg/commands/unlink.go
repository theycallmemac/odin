package commands

import (
    "fmt"
    "bytes"
    "os"
    "strconv"

    "github.com/spf13/cobra"
)

// define the UnunlinkCmd's metadata and run operation
var UnunlinkCmd = &cobra.Command{
    Use:   "unlink",
    Short: "unlinks the user's current Odin jobs",
    Long:  `This subcommand unlinks the user's current Odin jobs`,
    Run: func(cmd *cobra.Command, args []string) {
        from, _:= cmd.Flags().GetString("from")
        to, _:= cmd.Flags().GetString("to")
        port, _:= cmd.Flags().GetString("port")
        if port == "" {
            port = DefaultPort
        }
        unlinkJob(from, to, port)
    },
}

// add UnunlinkCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
    RootCmd.AddCommand(UnunlinkCmd)
    UnunlinkCmd.Flags().StringP("from", "f", "", "from")
    UnunlinkCmd.Flags().StringP("to", "t", "", "t")
    UnunlinkCmd.Flags().StringP("port", "p", "", "port")
    UnunlinkCmd.MarkFlagRequired("from")
    UnunlinkCmd.MarkFlagRequired("to")
}

// this function is called as the run operation for the UnunlinkCmd
// parameters: port (a string of the port to be used)
// returns: nil
func unlinkJob(from, to, port string) {
    uid := strconv.Itoa(os.Getuid())
    response := makePostRequest(fmt.Sprintf("http://localhost%s/links/delete", port), bytes.NewBuffer([]byte(fmt.Sprintf("%s", from + "_" + to + "_" + uid))))
    fmt.Print(response)
}
