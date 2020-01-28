package commands

import (
    "fmt"

    "github.com/spf13/cobra"
    "go.mongodb.org/mongo-driver/bson"
)

var DescribeCmd = &cobra.Command{
    Use:   "describe",
    Short: "describe a running Odin job",
    Long:  `This subcommand will describe a running Odin job created by the user`,
    Run: func(cmd *cobra.Command, args []string) {
            id, _:= cmd.Flags().GetString("id")
            describeJob(id)
    },
}

func init() {
    RootCmd.AddCommand(DescribeCmd)
    DescribeCmd.Flags().StringP("id", "i", "", "id (required)")
    DescribeCmd.MarkFlagRequired("id")
}

func describeJob(id string) {
    c := getMongoClient()
    job := getJobByValue(c, bson.M{"id": id})
    fmt.Println(job.Name + " - " + job.Description)
}
