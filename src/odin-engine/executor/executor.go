package executor

import (
    "os"
    "os/exec"
)

type Data struct {
    output []byte
    error  error
}


func exists(name string) bool {
    _, err := os.Stat(name)
    if os.IsNotExist(err) {
        return false
    }
    return err == nil
}

func runCommand(ch chan<- Data, language string, file string) {
    cmd := exec.Command(language, file)
    data, err := cmd.CombinedOutput()
    ch <- Data{
        error:  err,
        output: data,
    }
}

func Execute(filename string) bool {
    if exists(filename) {
        channel := make(chan Data)
        language, file := getYaml(filename)
        dir, _ := os.Getwd()
        go runCommand(channel, language, dir+"/"+file)
        res := <-channel
        if res.error != nil {
            return false
        }
        return true
    } else {
        return false
    }
}
