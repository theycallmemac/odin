package scheduler

import (
    "fmt"
    "os"
    "gopkg.in/yaml.v2"
)

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

func processError(err error) {
    fmt.Println(err)
    os.Exit(2)
}

func readFile(filename string) *os.File {
    file, err := os.Open(filename)
    if err != nil {
        processError(err)
        var tmp *os.File
        return tmp
    }
    return file
}

func parseYaml(cfg *Config, file *os.File)  bool {
    decoder:= yaml.NewDecoder(file)
    err := decoder.Decode(cfg)
    if err != nil {
        processError(err)
        return false
    }
    return true
}

func getYaml(filename string) string {
    var cfg Config
    yamlFile := readFile(filename)
    successfulParse := parseYaml(&cfg, yamlFile)
    if successfulParse {
        return cfg.Job.Schedule
    } else {
        return "Failed to read file."
    }
}

