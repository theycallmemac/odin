package commands

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path"
    "strings"

    "github.com/spf13/cobra"
)

// define the RecoverCmd's metadata and run operation
var RecoverCmd = &cobra.Command{
    Use:   "recover",
    Aliases: []string{"ls"},
    Short: "recovers files from a specific Odin job",
    Long:  `This subcommand recovers files from a specific Odin job`,
    Run: func(cmd *cobra.Command, args []string) {
        id, _ := cmd.Flags().GetString("id")
        src := "/etc/odin/jobs/" + id
        dest, _ := os.Getwd()
        recoverJob(src, dest)
    },
}

// add RecoverCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
    RootCmd.AddCommand(RecoverCmd)
    RecoverCmd.Flags().StringP("id", "i", "", "id to specific which job files to recover")
    RecoverCmd.MarkFlagRequired("id")
}

// this function is called as the run operation for the RecoverCmd
// parameters: port (a string of the port to be used)
// returns: nil
func recoverJob(src string, dest string) error {
    var err error
    var fds []os.FileInfo
    var srcfd *os.File
    var destfd *os.File
    if _, err = os.Stat(src); err != nil {
	return err
    }
    if fds, err = ioutil.ReadDir(src); err != nil {
	return err
    }
    var recoveries []string
    for _, fd := range fds {
	srcfp := path.Join(src, fd.Name())
	destfp := path.Join(dest, fd.Name())
	if srcfd, err = os.Open(srcfp); err != nil {
	    return err
	}
	defer srcfd.Close()
	if destfd, err = os.Create(destfp); err != nil {
	    return err
	}
	defer destfd.Close()
	if _, err = io.Copy(destfd, srcfd); err != nil {
	    return err
	} else {
            pathSplit := strings.Split(srcfp, "/")
            recoveries = append(recoveries, pathSplit[len(pathSplit)-1])
        }
    }
    for _, r := range recoveries {
        fmt.Println("File: " + r + " recovered!")
    }
    return nil
}
