package jobs

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

// create NewJob type to tbe used for accessing and storing job information
type NewJob struct {
    ID string `yaml:"id"`
    UID string `yaml:"uid"`
    GID string `yaml:"gid"`
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    Language string `yaml:"language"`
    File string `yaml:"file"`
    Status string `yaml:"status"`
    Schedule string `yaml:"schedule"`
}

// this function is used to set up a MongoDB client and test it with a ping command
// parameters: nil
// returns: *mogno.Client (a client)
func SetupClient() *mongo.Client {
    c := getMongoClient()
    err := c.Ping(context.Background(), readpref.Primary())
    if err != nil {
        os.Exit(2)
    }
    return c
}

// this function is used to get a MongoDB Client and set it's options
// parameters: none
// returns: *mogno.Client (a client)
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

// this function is used to add information to the MongoDB instance
// parameters: client (a *mongo.Client), d (a byte array containing marshaled JSON), and path (a string to set as the new job.File)
// returns: interface{} (an interface on the insertion results)
func InsertIntoMongo(client *mongo.Client, d []byte, path string) string {
    var job NewJob
    json.Unmarshal(d, &job)
    job.File = path
    if string(GetJobByValue(client, bson.M{"id": string(job.ID)}).ID) == string(job.ID) {
        return "Job already exists!"
    } else {
        collection := client.Database("myDatabase").Collection("myCollection")
        _, err := collection.InsertOne(context.TODO(), job)
        if err != nil {
            log.Fatalln("Error on inserting new job", err)
        }
        return "Job deployed successfully!"
    }
}

// this function is used to return a job in MongoDB by filtering on a certain value pertaining to that job
// parameters: client (a *mongo.Client), filter (a bson encoding of a job id)
// returns: NewJob (the fetched job)
func GetJobByValue(client *mongo.Client, filter bson.M) NewJob {
    var job NewJob
    collection := client.Database("myDatabase").Collection("myCollection")
    documentReturned := collection.FindOne(context.TODO(), filter)
    documentReturned.Decode(&job)
    return job
}

// this function is used to return all jobs in MongoDB
// parameters: client (a *mongo.Client)
// returns: []NewJob (all jobs in the Mongo instance)
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

// this function is used to format the output of MongoDB contents
// parameters: id, name, description, status, schedule (five strings corresponding to individual job data)
// returns: string (a space formatted string used for display)
func Format(id string, name, string, description string, status string, schedule string) string {
    return fmt.Sprintf("%-38s%-20s%-20s%-20s%-20s\n", id, name, description, status, schedule)
}

// this function is used to modify a job in MongoDB
// parameters: client (a *mongo.Client), job (a NewJob structure)
// returns: int64 (value of the number of entries modified)
func UpdateJobByValue(client *mongo.Client, job NewJob) int64 {
    update := bson.M{"$set": bson.M{"name": job.Name, "description": job.Description, "schedule": job.Schedule,},}
    collection := client.Database("myDatabase").Collection("myCollection")
    updateResult, _ := collection.UpdateOne(context.TODO(), bson.M{"id": job.ID}, update)
    return updateResult.ModifiedCount
}

// this function is used to delete a job in MongoDB
// parameters: parameters: client (a *mongo.Client), filter (a bson encoding of a job id)
// returns: int64 (value of the number of entries deleted)
func DeleteJobByValue(client *mongo.Client, filter bson.M) int64 {
    collection := client.Database("myDatabase").Collection("myCollection")
    deleteResult, _ := collection.DeleteOne(context.TODO(), filter)
    return deleteResult.DeletedCount
}

