package main

import (
    "fmt"
    "os"
    "os/exec"
)

type Data struct {
    output []byte
    error  error
}

func runCommand(ch chan<- Data, language string, file string) {
    cmd := exec.Command(language, file)
    data, err := cmd.CombinedOutput()
    ch <- Data{
        error:  err,
        output: data,
    }
}

func main() {
    channel := make(chan Data)
    filename := os.Args[1]
    language, file := getYaml(filename)
    dir, _ := os.Getwd()

    go runCommand(channel, language, dir+"/"+file)

    res := <-channel
    if res.error != nil {
        fmt.Println(res.error)
    }
}
