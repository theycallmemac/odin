package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/shirou/gopsutil/process"
	"github.com/spf13/cobra"
)

// NodesCmd is used to define the metadata and run operations for this command
var NodesCmd = &cobra.Command{
	Use:   "nodes [add|get]",
	Short: "interacts with odin engine nodes",
	Long:  `This subcommand interacts with odin engine nodes`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		switch args[0] {
		case "get":
			getNodesJob()
		case "add":
			name, _ := cmd.Flags().GetString("name")
			http, _ := cmd.Flags().GetString("addr")
			raft, _ := cmd.Flags().GetString("raft")
			if name == "" || http == "" || raft == "" {
				fmt.Println("You must supply values for:\n  - The name of the new node (e.g -n worker1)\n  - The address of the new node (e.g -a :39391)\n  - The raft address of the new node (e.g -r :12001)")
			} else {
				addNodesJob(name, http, raft)
			}
		}
	},
}

// add NodesCmd and it's respective flags
// parameters: nil
// returns: nil
func init() {
	RootCmd.AddCommand(NodesCmd)
	NodesCmd.Flags().StringP("name", "n", "", "specify the name of the new worker node (e.g worker1)")
	NodesCmd.Flags().StringP("addr", "a", "", "specify the http address of the new worker node (e.g :39391)")
	NodesCmd.Flags().StringP("raft", "r", "", "specify the raft address of the new worker node (e.g :12001)")
}

// Data is used to export an error type
type Data struct {
	error error
}

// this function is called as the get operation for the NodesCmd
// parameters: nil
// returns: nil
func getNodesJob() {
	pids, _ := process.Pids()
	fmt.Println(format("PID", "NAME", "HTTP PORT", "RAFT PORT"))
	var http, raft string
	for _, pid := range pids {
		proc, _ := process.NewProcess(pid)
		name, _ := proc.Name()
		if name == "odin-engine" {
			cmdline, _ := proc.Cmdline()
			line := strings.Split(string(cmdline), " ")
			if len(line) > 6 {
				name, http, raft = line[2], line[4], line[6]
			} else {
				name = line[2]
				http = ":3939"
				raft = ":12000"
			}
			fmt.Println(format(fmt.Sprint(pid), name, http, raft))
		}
	}
}

// this function is called as the add operation for the NodesCmd
// parameters: name (a string of the specified worker name), http (a string of the specified worker http address), raft (a string og the specified worker raft address)
// returns: nil
func addNodesJob(name, http, raft string) {
	c := make(chan Data)
	if len(http[1:]) > 5 {
		fmt.Println("The port " + http + " does not exist.")
		os.Exit(2)
	}
	if len(raft[1:]) > 5 {
		fmt.Println("The port " + raft + " does not exist.")
		os.Exit(2)
	}
	go createNode(c, name, http, raft)
	res := <-c
	if res.error != nil {
		fmt.Println("Failed to execute command, only root may add more nodes")
	} else {
		fmt.Println("Node deployed!")
	}
}

// this function is used to create a new node for the odin engine
// parameters: ch (a channel to pass information back to the calling function), name (a string of the specified worker name), http (a string of the specified worker http address), raft (a string og the specified worker raft address)
// returns: nil
func createNode(ch chan<- Data, name, http, raft string) {
	cmd := exec.Command("/bin/odin-engine", "-id", name, "-http", http, "-raft", raft, "-join", ":3939", "/etc/odin/"+name, "&")
	err := cmd.Start()
	ch <- Data{
		error: err,
	}
}

// this function is used to format the output of of the node list
// parameters: pid, name, http, raft (four strings corresponding to individual node data)
// returns: string (a space formatted string used for display)
func format(pid string, name string, http string, raft string) string {
	return fmt.Sprintf("%-20s%-20s%-20s%-20s\n", pid, name, http, raft)
}
