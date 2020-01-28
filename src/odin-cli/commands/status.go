package commands

import (
    "fmt"

    "github.com/spf13/cobra"
    "go.mongodb.org/mongo-driver/bson"
)

var StatusCmd = &cobra.Command{
    Use:   "status",
    Short: "return the status of an Odin job",
    Long:  `This subcommand will return the status of an Odin job`,
    Run: func(cmd *cobra.Command, args []string) {
            all, _ := cmd.Flags().GetBool("all")
            id, _:= cmd.Flags().GetString("id")
            if all {
                statusAll()
            } else if id != "" {
                statusJob(id)
            } else {
                cmd.Help()
            }
    },
}

func init() {
    var All bool
    RootCmd.AddCommand(StatusCmd)
    StatusCmd.Flags().BoolVarP(&All, "all", "a", false, "all")
    StatusCmd.Flags().StringP("id", "i", "", "id")
}

func statusJob(id string) {
    c := getMongoClient()
    job := getJobByValue(c, bson.M{"id": id})
    fmt.Println(job.Name + " - " + job.Status)
}

func statusAll() {
    c := getMongoClient()
    jobs := getAllJobs(c)
    for _, job := range jobs {
        fmt.Println(job.Name + " - " + job.Status)
    }
}
