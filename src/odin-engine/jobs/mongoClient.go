package jobs

import (
    "context"
    "fmt"
    "log"
    "os"
    "encoding/json"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/readpref"

)

type NewJob struct {
    ID string `yaml:"id"`
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    Language string `yaml:"language"`
    File string `yaml:"file"`
    Status string `yaml:"status"`
    Schedule string `yaml:"schedule"`
}

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

func InsertIntoMongo(client *mongo.Client, d []byte) interface{} {
    var job NewJob
    json.Unmarshal(d, &job)
    collection := client.Database("myDatabase").Collection("myCollection")
    insertResult, err := collection.InsertOne(context.TODO(), job)
    if err != nil {
        log.Fatalln("Error on inserting new job", err)
    }
    return insertResult.InsertedID
}

func GetJobByValue(client *mongo.Client, filter bson.M) NewJob {
    var job NewJob
    collection := client.Database("myDatabase").Collection("myCollection")
    documentReturned := collection.FindOne(context.TODO(), filter)
    documentReturned.Decode(&job)
    return job
}

func GetAll(client *mongo.Client) []NewJob {
    var jobs []NewJob
    collection := client.Database("myDatabase").Collection("myCollection")
    documents, _ := collection.Find(context.TODO(), bson.D{})
    for documents.Next(context.TODO()) {
        var job NewJob
        documents.Decode(&job)
        jobs = append(jobs, job)
    }
    return jobs
}

func Format(id string, name, string, description string, status string, schedule string) string {
    return fmt.Sprintf("%-38s%-20s%-20s%-20s%-20s\n", id, name, description, status, schedule)
}

func DeleteJobByValue(client *mongo.Client, filter bson.M) int64 {
    collection := client.Database("myDatabase").Collection("myCollection")
    deleteResult, _:= collection.DeleteOne(context.TODO(), filter)
    return deleteResult.DeletedCount
}

func SetupClient() *mongo.Client {
    c := getMongoClient()
    err := c.Ping(context.Background(), readpref.Primary())
    if err != nil {
        os.Exit(2)
    }
    return c
}
