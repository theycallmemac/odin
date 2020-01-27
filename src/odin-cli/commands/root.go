package commands

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"/*
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/readpref"*/
)



// ----------------------- INIT COBRA ROOT CMD ---------------------- //
// ------------------------------------------------------------------ //
var RootCmd = &cobra.Command{
    Use:   "odin",
    Short: "orchestrate your jobs",
    Long: `orchestrate your jobs for periodic execution`,
}

func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}



// ------------------------- SHARED FUNCTIONS ----------------------- //
// ------------------------------------------------------------------ //
func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}



// -------------------------- SHARED STRUCTS ------------------------ //
// ------------------------------------------------------------------ //
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


type NewJob struct {
    ID string `yaml:"id"`
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    Language string `yaml:"language"`
    File string `yaml:"file"`
    ScheduleString string `yaml:"schedulestring"`
}


/*
// -------------------------- MONGO FUNCTIONS ----------------------- //
// ------------------------------------------------------------------ //
func getMongoClient() *mongo.Client {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
	log.Fatal(err)
    }
    err = client.Connect(context.Background())
    if err != nil {
	log.Fatal(err)
    }
    return client
}

func insertIntoMongo(client *mongo.Client, job NewJob) interface{} {
    collection := client.Database("myDatabase").Collection("myCollection")
    insertResult, err := collection.InsertOne(context.TODO(), job)
    if err != nil {
        log.Fatalln("Error on inserting new job", err)
    }
    return insertResult.InsertedID
}

func getJobByValue(client *mongo.Client, filter bson.M) Job {
    var job NewJob
    collection := client.Database("myDatabase").Collection("myCollection")
    documentReturned := collection.FindOne(context.TODO(), filter)
    documentReturned.Decode(&job)
    return job
}
*/
