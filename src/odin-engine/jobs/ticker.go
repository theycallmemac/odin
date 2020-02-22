package jobs

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "math/rand"
    "net/http"
    "time"

    "github.com/gorhill/cronexpr"
)

type Queue struct {
    Items []Node
}

type Node struct {
    ID string
    UID string
    GID string
    Lang string
    File string
    Schedule int
}

func MakePostRequest(link string, data *bytes.Buffer) string {
    client := &http.Client{}
    req, _ := http.NewRequest("POST", link, data)
    response, clientErr := client.Do(req)
    if clientErr != nil {
        fmt.Println(clientErr)
    }
    bodyBytes, _ := ioutil.ReadAll(response.Body)
    return string(bodyBytes)
}


func sortQueue(items []Node, done chan int) {
    if len(items) < 2 {
        done <- 1
        return
    }
    left, right := 0, len(items) - 1
    pivot := rand.Int() % len(items)
    items[pivot], items[right] = items[right], items[pivot]
    for i, _ := range items {
        if items[i].Schedule < items[right].Schedule {
            items[left], items[i] = items[i], items[left]
            left++
        }
    }
    items[left], items[right] = items[right], items[left]
    childChan := make(chan int)
    go sortQueue(items[:left], childChan)
    go sortQueue(items[left+1:], childChan)
    for i := 0; i < 2; i++ {
        <-childChan
    }
    done <- 1
    return
}

func checkHead(items map[int][]Node) {
    if value, ok := items[0]; ok {
        for _, job := range value {
            go MakePostRequest("http://localhost:3939/execute", bytes.NewBuffer([]byte(job.UID + " " + job.GID + " " + job.Lang + " " + job.File + " " + job.ID)))
            fmt.Println("executed job")
        }
    }
}

func groupItems(items []Node) map[int][]Node {
    output := make(map[int][]Node)
    for _, item := range items {
        output[item.Schedule] = append(output[item.Schedule], item)
    }
    return output
}

func fillQueue(t time.Time) {
    var queue Queue
    var node Node
    jobs := GetAll(SetupClient())
    for _, j := range jobs {
        node.ID, node.UID, node.GID, node.Lang, node.File = j.ID, j.UID, j.GID, j.Language, j.File
        if len(j.Schedule) > 0 {
            node.Schedule = int(cronexpr.MustParse(j.Schedule[:len(j.Schedule)-1]).Next(time.Now()).Sub(time.Now()).Seconds())
            queue.Items = append(queue.Items, node)
            channel := make(chan int)
            go sortQueue(queue.Items, channel)
	    <-channel
        }
    }
    go checkHead(groupItems(queue.Items))
}

func doEvery(d time.Duration, f func(time.Time)) {
    for x := range time.Tick(d) {
        go f(x)
    }
}

func StartTicker() {
   go doEvery(1000*time.Millisecond, fillQueue)
}
