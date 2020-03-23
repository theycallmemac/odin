package jobs

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "os/user"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/readpref"

    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/resources"
    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/types"

    "gopkg.in/yaml.v2"
)

// create NewJob type to be used for accessing and storing job information
type NewJob struct {
    ID string `yaml:"id"`
    UID string `yaml:"uid"`
    GID string `yaml:"gid"`
    Name string `yaml:"name"`
    Description string `yaml:"description"`
    Language string `yaml:"language"`
    File string `yaml:"file"`
    Stats string `yaml:"stats"`
    Schedule string `yaml:"schedule"`
    Runs int
}

// create JobStats type to be used for accessing and storing job stats information
type JobStats struct {
    ID string
    Description string
    Type string
    Value string
}

// this function is used to unmarshal YAML
// parameters: byteArray (an array of bytes representing the contents of a file)
// returns: Config (a struct form of the YAML)
func unmarsharlYaml(byteArray []byte) types.EngineConfig {
    var cfg types.EngineConfig
    err := yaml.Unmarshal([]byte(byteArray), &cfg)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    return cfg
}

// this function is used to set up a MongoDB client and test it with a ping command
// parameters: nil
// returns: *mogno.Client (a client)
func SetupClient() *mongo.Client {
    c := getMongoClient()
    err := c.Ping(context.Background(), readpref.Primary())
    if err != nil {
        fmt.Println("Cannot connect to MongoDB")
        os.Exit(2)
    }
    return c
}

// this function is used to get a MongoDB Client and set it's options
// parameters: none
// returns: *mogno.Client (a client)
func getMongoClient() *mongo.Client {
    usr, _ := user.Current()
    clientOptions := options.Client().ApplyURI(resources.UnmarsharlYaml(resources.ReadFileBytes(usr.HomeDir + "/odin-config.yml")).Mongo.Address)
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
func InsertIntoMongo(client *mongo.Client, d []byte, path string, uid string) string {
    var job NewJob
    json.Unmarshal(d, &job)
    job.File = path
    job.Runs = 0
    if string(GetJobByValue(client, bson.M{"id": string(job.ID)}, uid).ID) == string(job.ID) {
        return "Job with ID: " + job.ID + " already exists\n"
    } else {
        collection := client.Database("odin").Collection("jobs")
        _, err := collection.InsertOne(context.TODO(), job)
        if err != nil {
            log.Fatalln("Error on inserting new job", err)
        }
        return "Job: " + job.ID + " deployed successfully\n"
    }
}

// this function is used to retrieve the stats associated with each job from the MongoDB instance
// paramters: parameters: client (a *mongo.Client), id (a string representation of a job's id)
// retuns []JobStats (the colleciton of fetched job stats)
func GetJobStats(client *mongo.Client, id string) []JobStats {
    var statMap map[string]string
    var jobStats JobStats
    var statsList []JobStats
    collection := client.Database("odin").Collection("observability")
    documents, _ := collection.Find(context.TODO(), bson.M{"id": id})
    for documents.Next(context.TODO()) {
        documents.Decode(&statMap)
        jobStats.ID = statMap["id"]
        jobStats.Description = statMap["desc"]
        jobStats.Type = statMap["type"]
        jobStats.Value = statMap["value"]
        statsList = append(statsList, jobStats)
    }
    return statsList
}

// this function is used to return a job in MongoDB by filtering on a certain value pertaining to that job
// parameters: client (a *mongo.Client), filter (a bson encoding of a job id), uid (a string of the user's ID)
// returns: NewJob (the fetched job)
func GetJobByValue(client *mongo.Client, filter bson.M, uid string) NewJob {
    var job NewJob
    collection := client.Database("odin").Collection("jobs")
    documentReturned := collection.FindOne(context.TODO(), filter)
    documentReturned.Decode(&job)
    if job.UID == uid {
        return job
    }
    var tmp NewJob
    return tmp
}

// this function is used to return a specific user's jobs from MongoDB
// parameters: client (a *mongo.Client), uid (a string of that user's id)
// returns: []NewJob (all jobs in the Mongo instance)
func GetUserJobs(client *mongo.Client, uid string) []NewJob {
    var jobs []NewJob
    collection := client.Database("odin").Collection("jobs")
    documents, _ := collection.Find(context.TODO(), bson.D{})
    for documents.Next(context.TODO()) {
        var job NewJob
        documents.Decode(&job)
        if job.UID == uid || uid == "0" {
            jobs = append(jobs, job)
        }
    }
    return jobs
}

// this function is used to return all jobs in MongoDB
// parameters: client (a *mongo.Client)
// returns: []NewJob (all jobs in the Mongo instance)
func GetAll(client *mongo.Client) []NewJob {
    var jobs []NewJob
    collection := client.Database("odin").Collection("jobs")
    documents, _ := collection.Find(context.TODO(), bson.D{})
    for documents.Next(context.TODO()) {
        var job NewJob
        documents.Decode(&job)
        jobs = append(jobs, job)
    }
    return jobs
}

// this function is used to format the output of MongoDB contents
// parameters: id, name, description, stats, schedule (five strings corresponding to individual job data)
// returns: string (a space formatted string used for display)
func Format(id string, name, string, description string, schedule string) string {
    return fmt.Sprintf("%-20s%-20s%-20s%-20s\n", id, name, description, schedule)
}

// this function is used to modify a job in MongoDB
// parameters: client (a *mongo.Client), job (a NewJob structure)
// returns: int64 (value of the number of entries modified)
func UpdateJobByValue(client *mongo.Client, job NewJob) int64 {
    update := bson.M{"$set": bson.M{"name": job.Name, "description": job.Description, "schedule": job.Schedule, "runs": job.Runs},}
    collection := client.Database("odin").Collection("jobs")
    updateResult, _ := collection.UpdateOne(context.TODO(), bson.M{"id": job.ID}, update)
    return updateResult.ModifiedCount
}

// this function is used to delete a job in MongoDB
// parameters: parameters: client (a *mongo.Client), filter (a bson encoding of a job id), uid (a string of the user's ID)
// returns: bool (whether a job was deleted or not)
func DeleteJobByValue(client *mongo.Client, filter bson.M, uid string) bool {
    job := GetJobByValue(client, filter, uid)
    if job.ID == "" || job.UID != uid {
        return false
    }
    collection := client.Database("odin").Collection("jobs")
    _, err := collection.DeleteOne(context.TODO(), filter)
    if err != nil {
        return false
    }
    return true
}
