package odinlib

import (
	"os"

	"gopkg.in/yaml.v2"
)

// JobConfig is a type used to structure information from job Yaml files
type JobConfig struct {
	Provider struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"provider"`
	Job struct {
		Name        string `yaml:"name"`
		ID          string `yaml:"id"`
		Description string `yaml:"description"`
		Language    string `yaml:"language"`
		File        string `yaml:"file"`
		Schedule    string `yaml:"schedule"`
	} `yaml:"job"`
}

// ReadFile is used to read a specific files contents
// parameters: filename (a string containing the name of the file to be read)
// returns: *os.File (a pointer to the read file)
func ReadFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	return file
}

// ParseYaml is used to structure information found in job Yaml Files
// parameters: cfg (a JobConfig type), file (a file which has been read)
// returns: bool (true if successfully decoded, false if otherwise)
func ParseYaml(cfg *JobConfig, file *os.File) bool {
	decoder := yaml.NewDecoder(file)
	err := decoder.Decode(cfg)
	if err != nil {
		return false
	}
	return true
}
