package commands

import (
    "github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
    Use:   "list",
    Short: "lists the user's current Odin jobs",
    Long:  `This subcommand lists the user's current Odin jobs`,
    Run: func(cmd *cobra.Command, args []string) {
            listJob()
    },
}

func init() {
    RootCmd.AddCommand(ListCmd)
}

func listJob() {
    c := getMongoClient()
    format("ID", "NAME", "DESCRIPTION", "LANGUAGE", "STATUS", "SCHEDULE")
    jobs := getAllJobs(c)
    for _, job := range jobs {
        format(job.ID, job.Name, job.Description, job.Language, job.Status, job.Schedule)
    }
}
