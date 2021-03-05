package jobs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/user"
	"strconv"
	"strings"
	"time"

	"github.com/lnquy/cron"
	"github.com/theycallmemac/odin/odin-engine/pkg/fsm"
	"github.com/theycallmemac/odin/odin-engine/pkg/resources"
	"github.com/theycallmemac/odin/odin-engine/pkg/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"gopkg.in/yaml.v2"
)

// URI is used to store the address to the MongoDB instance used by the Odin Engine
var URI = resources.UnmarsharlYaml(resources.ReadFileBytes(getHome() + "/odin-config.yml")).Mongo.Address

// NewJob is a type to be used for accessing and storing job information
type NewJob struct {
	ID          string `yaml:"id"`
	UID         string `yaml:"uid"`
	GID         string `yaml:"gid"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Language    string `yaml:"language"`
	File        string `yaml:"file"`
	Stats       string `yaml:"stats"`
	Schedule    string `yaml:"schedule"`
	Runs        int
	Links       string
}

// JobStats is a type to be used for accessing and storing job stats information
type JobStats struct {
	ID          string
	Description string
	Type        string
	Value       string
}


// getHome is used to get the path to the user's home directory
// parameters: nil
// return string (the path to the user's home)
func getHome() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

// unmarsharlYaml is used to unmarshal YAML
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

// SetupClient is used to set up a MongoDB client and test it with a ping command
// parameters: nil
// returns: *mogno.Client (a client)
func SetupClient() (*mongo.Client, error) {
	c := getMongoClient()
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := c.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Cannot connect to MongoDB: ", err)
		c.Disconnect(context.TODO())
	}
	return c, err
}

// getMongoClient is used to get a MongoDB Client and set it's options
// parameters: none
// returns: *mogno.Client (a client)
func getMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(URI)
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

// InsertIntoMongo is used to add information to the MongoDB instance
// parameters: client (a *mongo.Client), d (a byte array containing marshaled JSON), and path (a string to set as the new job.File)
// returns: interface{} (an interface on the insertion results)
func InsertIntoMongo(client *mongo.Client, d []byte, path string, uid string) string {
	var job NewJob
	json.Unmarshal(d, &job)
	job.File = path
	job.Runs = 0
	if string(GetJobByValue(client, bson.M{"id": string(job.ID)}, uid).ID) == string(job.ID) {
		return "Job with ID: " + job.ID + " already exists\n"
	}
        collection := client.Database("odin").Collection("jobs")
        _, err := collection.InsertOne(context.TODO(), job)
        client.Disconnect(context.TODO())
        if err != nil {
                log.Fatalln("Error on inserting new job", err)
        }
        return "Job: " + job.ID + " deployed successfully\n"
}

// GetJobStats is used to retrieve the stats associated with each job from the MongoDB instance
// parameters: client (a *mongo.Client), id (a string representation of a job's id)
// returns: []JobStats (the collection of fetched job stats)
func GetJobStats(client *mongo.Client, id string) []JobStats {
	var statMap map[string]string
	var jobStats JobStats
	var statsList []JobStats
	collection := client.Database("odin").Collection("observability")
	documents, _ := collection.Find(context.TODO(), bson.M{"id": id})
	client.Disconnect(context.TODO())
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

// GetJobByValue is used to return a job in MongoDB by filtering on a certain value pertaining to that job
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

// GetUserJobs is used to return a specific user's jobs from MongoDB
// parameters: client (a *mongo.Client), uid (a string of that user's id)
// returns: []NewJob (all jobs in the Mongo instance)
func GetUserJobs(client *mongo.Client, uid string) []NewJob {
	var jobs []NewJob
	collection := client.Database("odin").Collection("jobs")
	documents, _ := collection.Find(context.TODO(), bson.D{})
	client.Disconnect(context.TODO())
	for documents.Next(context.TODO()) {
		var job NewJob
		documents.Decode(&job)
		if job.UID == uid || uid == "0" {
			jobs = append(jobs, job)
		}
	}
	return jobs
}

// GetAll is used to return all jobs in MongoDB
// parameters: client (a *mongo.Client)
// returns: []NewJob (all jobs in the Mongo instance)
func GetAll(client *mongo.Client) []NewJob {
	var jobs []NewJob
	collection := client.Database("odin").Collection("jobs")
	documents, _ := collection.Find(context.TODO(), bson.D{})
	client.Disconnect(context.TODO())
	for documents.Next(context.TODO()) {
		var job NewJob
		documents.Decode(&job)
		jobs = append(jobs, job)
	}
	return jobs
}

// Format is used to format the output of MongoDB stat contents
// parameters: id, description, valType, value (four strings corresponding to individual job stats)
// returns: string (a space formatted string used for display)
func Format(id string, description string, valType string, value string) string {
	return fmt.Sprintf("%-20s%-20s%-20s%-20s\n", id, description, valType, value)
}

// SchFormat is used to parse and format the output of the MongoDB schedule contents
// parameters: id, name, description, schedule (four strings corresponding to individual job data)
// returns: string (a space formatted string used for display)
func SchFormat(id string, name, string, description string, links string, schedule string) string {
	var finalSchedule = ""
	var tmpSchedule = ""
	if schedule == "0 5 31 2 *" {
		finalSchedule = "never"
	} else if schedule != "SCHEDULE" {
		scheduleArray := strings.Split(schedule, ",")
		for i, item := range scheduleArray {
			descriptor, _ := cron.NewDescriptor()
			tmpSchedule, _ = descriptor.ToDescription(item, cron.Locale_en)
			if i+1 == len(scheduleArray) {
				finalSchedule += tmpSchedule
			} else {
				finalSchedule += tmpSchedule + " & "
			}
		}
	} else {
		finalSchedule = schedule
	}
	return fmt.Sprintf("%-20s%-20s%-20s%-20s%-20s\n", id, name, description, links, finalSchedule)
}

// UpdateJobByValue is used to modify a job in MongoDB
// parameters: client (a *mongo.Client), job (a NewJob structure)
// returns: int64 (value of the number of entries modified)
func UpdateJobByValue(client *mongo.Client, job NewJob) int64 {
	update := bson.M{"$set": bson.M{"name": job.Name, "description": job.Description, "schedule": job.Schedule, "runs": job.Runs}}
	collection := client.Database("odin").Collection("jobs")
	updateResult, err := collection.UpdateOne(context.TODO(), bson.M{"id": job.ID}, update)
	client.Disconnect(context.TODO())
	if err != nil {
		return int64(0)
	}
	return updateResult.ModifiedCount
}

// DeleteJobByValue is used to delete a job in MongoDB
// parameters: parameters: client (a *mongo.Client), filter (a bson encoding of a job id), uid (a string of the user's ID)
// returns: bool (whether a job was deleted or not)
func DeleteJobByValue(client *mongo.Client, filter bson.M, uid string) bool {
	job := GetJobByValue(client, filter, uid)
	if job.ID == "" || job.UID != uid {
		return false
	}
	collection := client.Database("odin").Collection("jobs")
	_, err := collection.DeleteOne(context.TODO(), filter)
	client.Disconnect(context.TODO())
	if err != nil {
		return false
	}
	return true
}

// AddJobLink is used to add links the job is associated with
// parameters: client (a *mongo.Client), from (a string of a job ID to give a new link), to (a string of a job ID to create a link to), uid (a string of the user's ID)
// returns: int64 (value of the number of entries modified)
func AddJobLink(client *mongo.Client, from string, to string, uid string) int64 {
	job := GetJobByValue(client, bson.M{"id": string(from)}, uid)
	if strings.Contains(job.Links, to) {
		return 0
	}
	job.Links = job.Links + to + ","
	update := bson.M{"$set": bson.M{"links": job.Links}}
	collection := client.Database("odin").Collection("jobs")
	updateResult, _ := collection.UpdateOne(context.TODO(), bson.M{"id": from}, update)
	client.Disconnect(context.TODO())
	return updateResult.ModifiedCount
}

// DeleteJobLink is used to delete links the job is associated with
// parameters: client (a *mongo.Client), from (a string of a job ID to remove a link from), to (a string of a job ID to remove), uid (a string of the user's ID)
// returns: int64 (value of the number of entries modified)
func DeleteJobLink(client *mongo.Client, from string, to string, uid string) int64 {
	var newLinks string
	job := GetJobByValue(client, bson.M{"id": string(from)}, uid)
	links := strings.Split(job.Links, ",")
	for _, link := range links {
		if link != to && link != "" {
			newLinks = newLinks + link + ","
		}
	}
	update := bson.M{"$set": bson.M{"links": newLinks}}
	collection := client.Database("odin").Collection("jobs")
	updateResult, _ := collection.UpdateOne(context.TODO(), bson.M{"id": job.ID}, update)
	client.Disconnect(context.TODO())
	return updateResult.ModifiedCount
}


// RunLinks is used to run jobs linked to a job which has just been executed
// parameters: links (a string array of Job ID's to execute), uid (a uint32 of that user's id), httpAddr (a string port of the master node), store (a fsm.Store containing information about other nodes)
// returns: nil
func RunLinks(links []string, uid uint32, httpAddr string, store fsm.Store) {
	client, _ := SetupClient()
	var jobs []Node
	var node Node
	for _, link := range links {
		job := GetJobByValue(client, bson.M{"id": string(link)}, fmt.Sprint(uid))
		node.ID, node.Lang, node.File, node.Links = job.ID, job.Language, job.File, job.Links
		uid, _ := strconv.ParseUint(job.UID, 10, 32)
		gid, _ := strconv.ParseUint(job.GID, 10, 32)
		node.UID = uint32(uid)
		node.GID = uint32(gid)
		jobs = append(jobs, node)
	}
	client.Disconnect(context.TODO())
	var en ExecNode
	jobsArray, _ := json.Marshal(jobs)
	en.Items = jobsArray
	en.Store = store
	buffer, _ := json.Marshal(en)
	go MakePostRequest("http://localhost"+httpAddr+"/execute", bytes.NewBuffer(buffer))
}
