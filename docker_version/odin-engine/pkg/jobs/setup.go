package jobs

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"strings"
)

var job NewJob

// GID is used to store the group id associated with a specific group
var GID int

// notDirectory is used to check whether or not a directory exists
// parameters: dir (a string of the directory path to check)
// returns: boolean (true if it exists, false otherwise)
func notDirectory(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return true
	}
	return false
}

// makeDirectory is used to create a new directory
// parameters: name (a string of the directory path to create)
// returns: bool
func makeDirectory(name string) bool {
	err := os.MkdirAll(name, 0654)
	if err != nil {
		if notDirectory(name) {
			return false
		}
	}
	return true
}

// ChownR is used to recursively change the owner of each subdir under /etc
// parameters: path (a string of the directory path to chown), uid (an int used to set the owner uid), gid (an int used to set the owner gid)
// returns: bool
func ChownR(path string, uid, gid int) bool {
	s := strings.Split(path, "/")
	for i := len(s) - 1; i > 2; i-- {
		err := os.Chown(strings.Join(s[:i], "/"), uid, gid)
		if err != nil {
			return false
		}
	}
	return true
}

// createNewPath is used to create a new path for the runtime file
// parameters: jobsPath (a string containing the jobs directory), file (a string containing the source runtime file)
// returns: string (the full path to the files new location)
func createNewPath(jobsPath string, file string) string {
	fileSlice := strings.Split(file, "/")
	job.File = string(fileSlice[len(fileSlice)-1])
	return jobsPath + job.ID + "/" + job.File

}

// createNewConfigPath( is used to create a new path for the config
// parameters: jobsPath (a string containing the jobs directory), file (a string containing the source config file)
// returns: string (the full path to the configs new location)
func createNewConfigPath(jobsPath string, file string) string {
	fileSlice := strings.Split(file, "/")
	file = string(fileSlice[len(fileSlice)-1])
	return jobsPath + job.ID + "/" + file
}

// copyFile is used copy the contents of the user generated files up to /etc/odin/jobs/<id>/filename.ext
// parameters: name (the name of the file to read contents from)
// returns: []byte (a byte array of the content read from the file)
func copyFile(name string) []byte {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		return nil
	}
	return content
}

// SetupEnvironment is used to kickstart the process for setting up the correct directories and files used by odin
// parameters: d (a byte array containing marshaled JSON)
// returns: string (the path to the newly created file)
func SetupEnvironment(d []byte) string {
	var originalConfig string
	err := json.Unmarshal(d, &job)

	if err != nil {
		return ""
	}

	jobsPath := "/etc/odin/jobs/"
	logsPath := "/etc/odin/logs/"

	if notDirectory(jobsPath) && notDirectory(logsPath) {
		makeDirectory(jobsPath)
		makeDirectory(logsPath)
	}

	originalFile := job.File
	if strings.Contains(job.File, ".") {
		originalConfig = job.File[:len(originalFile)-3] + ".yml"
	} else {
		originalConfig = job.File + ".yml"
	}

	newFilePath := createNewPath(jobsPath, job.File)
	newConfigPath := createNewPath(jobsPath, originalConfig)

	if notDirectory(jobsPath + job.ID) {
		makeDirectory(jobsPath + job.ID)
		group, _ := user.LookupGroup("odin")
		gid, _ := strconv.Atoi(group.Gid)
		GID = gid
		ChownR(newFilePath, 0, gid)
		ChownR(newConfigPath, 0, gid)
		ChownR(logsPath+job.ID, 0, gid)
		ioutil.WriteFile(newFilePath, []byte(""), 0774)
		ioutil.WriteFile(newConfigPath, []byte(""), 0774)
		ioutil.WriteFile(logsPath+job.ID, []byte(""), 0774)
	}

	fileInput := copyFile(originalFile)
	configInput := copyFile(originalConfig)

	ioutil.WriteFile(logsPath+job.ID, []byte(""), 0774)
	ioutil.WriteFile(newFilePath, fileInput, 0774)
	ioutil.WriteFile(newConfigPath, configInput, 0774)

	os.Chown(newFilePath, 0, GID)
	os.Chown(newConfigPath, 0, GID)
	os.Chown(logsPath+job.ID, 0, GID)
	return newFilePath
}
