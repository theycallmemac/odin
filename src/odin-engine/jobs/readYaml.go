package jobs

import (
    "fmt"
    "os"

    "gopkg.in/yaml.v2"
)

// create Config type to tbe used for accessing config information
type Config struct {
    Provider struct {
        Name string `yaml:"name"`
        Version string `yaml:"version"`
    } `yaml:"provider"`
    Job struct {
        Name string `yaml:"name"`
        Description string `yaml:"description"`
        Language string `yaml:"language"`
        File string `yaml:"file"`
        Schedule string `yaml:"schedule"`
    } `yaml:"job"`
}

// this function is used to handle an error and exit upon doing so
func processError(err error) {
    fmt.Println(err)
    os.Exit(2)
}

// this function is used to read a file and return it's contents
func readFile(filename string) *os.File {
    file, err := os.Open(filename)
    if err != nil {
        processError(err)
        var tmp *os.File
        return tmp
    }
    return file
}

// this function is used to parse a given YAML config
func parseYaml(cfg *Config, file *os.File)  bool {
    decoder:= yaml.NewDecoder(file)
    err := decoder.Decode(cfg)
    if err != nil {
        processError(err)
        return false
    }
    return true
}

// this function is used to return the YAML of a config
func ToYaml(filename string, job NewJob) Config {
    var cfg Config
    yamlFile := readFile(filename)
    successfulParse := parseYaml(&cfg, yamlFile)
    if successfulParse {
        return cfg
    } else {
        var tmpCfg Config
        return tmpCfg
    }
}

