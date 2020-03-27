package resources

import (
    "io/ioutil"
    "os"

    "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-engine/internal/types"

    "gopkg.in/yaml.v2"
)


/*                                               STRING FUNCTIONS                                           */
//////////////////////////////////////////////////////////////////////////////////////////////////////////////

// this function is used to check if a string is not empty
// parameters: s (a string to check)
// returns: boolean (true if not empty, false if otherwise)
func NotEmpty(s string) bool {
    if s == "" {
        return false
    }
    return true
}

/*                                               TEXT FILE FUNCTIONS                                        */
//////////////////////////////////////////////////////////////////////////////////////////////////////////////

// this function is used to read a file and return it's contents
// parameters: filename (a string of the path to the file)
// returns: *os.File (the file descriptor)
func ReadFile(filename string) *os.File {
    file, err := os.Open(filename)
    if err != nil {
        return nil
    }
    return file
}

// this function is used to read a file
// parameters: name (a string containing the path to a file)
// returns: []byte (an array of bytes containing the contents of the file)
func ReadFileBytes(name string) []byte {
    file, err := os.Open(name)
    if err != nil {
        return nil
    }
    bytes, err := ioutil.ReadAll(file)
    defer file.Close()
    return bytes
}


/*                                                  YAML FUNCTIONS                                          */
//////////////////////////////////////////////////////////////////////////////////////////////////////////////

// this function is used to unmarshal YAML
// parameters: byteArray (an array of bytes representing the contents of a file)
// returns: Config (a struct form of the YAML)
func UnmarsharlYaml(byteArray []byte) types.EngineConfig {
    var cfg types.EngineConfig
    yaml.Unmarshal([]byte(byteArray), &cfg)
    return cfg
}

// this function is used to return the yaml attribute needed for the scheduler endpoint
// parameters: filename (a string containing a file to read)
// returns: string (the job schedule if successfully parsed, a failure message if otherwise)
func SchedulerYaml(filename string) string {
    var cfg types.JobConfig
    if ParseYaml(&cfg, ReadFile(filename)) {
        return cfg.Job.Schedule
    } else {
        return "Failed to read file."
    }
}

// this function is used to return the yaml attribute needed for the executor endpoint
// parameters: filename (a string containing a file to read)
// returns: string, string (the job language and file if successful, a failure message if otherwise)
func ExecutorYaml(filename string) (string, string) {
    var cfg types.JobConfig
    if ParseYaml(&cfg, ReadFile(filename)) {
        return cfg.Job.Language, cfg.Job.File
    } else {
        return "Failed to read file.", ""
    }
}

// this function is used to parse a given YAML config
// parameters: cfg (a *Config to be decoded into), file, (am *os.File to build the decoder on)
// returns: boolean (true if parseable, false if otherwise)
func ParseYaml(cfg *types.JobConfig, file *os.File)  bool {
    decoder := yaml.NewDecoder(file)
    err := decoder.Decode(cfg)
    if err != nil {
        return false
    }
    return true
}


