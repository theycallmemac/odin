package jobs

import (
    "fmt"
    "log"
    "io/ioutil"
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
// parameters: err (an error to print out)
// returns: nil
func processError(err error) {
    fmt.Println(err)
    os.Exit(2)
}


// this function is used to check if a string is not empty
// parameters: s (a string to check)
// returns: boolean (true if not empty, false if otherwise)
func NotEmpty(s string) bool {
    if s == "" {
        return false
    }
    return true
}

// this function is used to read a file and return it's contents
// parameters: filename (a string of the path to the file)
// returns: *os.File (the file descriptor)
func readFile(filename string) *os.File {
    file, err := os.Open(filename)
    if err != nil {
        processError(err)
        var tmp *os.File
        return tmp
    }
    return file
}

// this function is used to read a file
// parameters: name (a string containing the path to a file)
// returns: []byte (an array of bytes containing the contents of the file)
func readConfigFile(name string) []byte {
    file, err := os.Open(name)
    if err != nil {
        log.Fatal(err)
    }
    bytes, err := ioutil.ReadAll(file)
    defer file.Close()
    return bytes
}

// this function is used to parse a given YAML config
// parameters: cfg (a *Config to be decoded into), file, (am *os.File to build the decoder on)
// returns: boolean (true if parseable, false if otherwise)
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
// parameters: filename (a string of the path to the file), job (a NewJob)
// returns: Config (either a successfully parse config or an empty one)
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

