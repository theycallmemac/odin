package odinlib

import (
    "os"

    "gopkg.in/yaml.v2"
)

type JobConfig struct {
    Provider struct {
        Name string `yaml:"name"`
        Version string `yaml:"version"`
    } `yaml:"provider"`
    Job struct {
        Name string `yaml:"name"`
        ID string `yaml:"id"`
        Description string `yaml:"description"`
        Language string `yaml:"language"`
        File string `yaml:"file"`
        Schedule string `yaml:"schedule"`
    } `yaml:"job"`
}

func ReadFile(filename string) *os.File {
    file, err := os.Open(filename)
    if err != nil {
        return nil
    }
    return file
}

func ParseYaml(cfg *JobConfig, file *os.File) bool {
    decoder := yaml.NewDecoder(file)
    err := decoder.Decode(cfg)
    if err != nil {
        return false
    }
    return true
}

