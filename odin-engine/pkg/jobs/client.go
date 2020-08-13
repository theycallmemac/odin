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

	"github.com/lnquy/cron"
	"github.com/theycallmemac/odin/odin-engine/pkg/fsm"
	"github.com/theycallmemac/odin/odin-engine/pkg/repository"
	"github.com/theycallmemac/odin/odin-engine/pkg/resources"
	"github.com/theycallmemac/odin/odin-engine/pkg/types"

	"gopkg.in/yaml.v2"
)

// URI is used to store the address to the MongoDB instance used by the Odin Engine
var URI = resources.UnmarsharlYaml(resources.ReadFileBytes(getHome() + "/odin-config.yml")).Mongo.Address

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

// RunLinks is used to run jobs linked to a job which has just been executed
// parameters: links (a string array of Job ID's to execute), uid (a uint32 of that user's id), httpAddr (a string port of the master node), store (a fsm.Store containing information about other nodes)
// returns: nil
func RunLinks(repo repository.Repository, links []string, uid uint32, httpAddr string, store fsm.Store) {
	ctx := context.Background()
	var jobs []Node
	var node Node
	for _, link := range links {
		id := string(link)
		job, err := repo.GetJobById(ctx, id, fmt.Sprint(uid))
		if err != nil {
			fmt.Printf("[FAILED] Failed to get job (%s) (%d) while running links", id, uid)
			continue
		}
		node.ID, node.Lang, node.File, node.Links = job.ID, job.Language, job.File, job.Links
		uid, _ := strconv.ParseUint(job.UID, 10, 32)
		gid, _ := strconv.ParseUint(job.GID, 10, 32)
		node.UID = uint32(uid)
		node.GID = uint32(gid)
		jobs = append(jobs, node)
	}
	var en ExecNode
	jobsArray, _ := json.Marshal(jobs)
	en.Items = jobsArray
	en.Store = store
	buffer, _ := json.Marshal(en)
	go MakePostRequest("http://localhost"+httpAddr+"/execute", bytes.NewBuffer(buffer))
}
