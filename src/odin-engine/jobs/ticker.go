package jobs

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math/rand"
    "net/http"
    "sort"
    "strings"
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
    Schedule []int
}

// this function is used to make a post request to a given url
// parameters: link (a string of the link to make a request to), data (a buffer to pass to the post request) 
// returns: string (the result of a POST to the provided link with the given data)
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

// this function is used to sort an array of nodes using an implementation of recursive quicksort which acts as a goroutine
// parameters: items (an array of jobs in the queue), done (a channel to signify when the routine has finished)
// returns: nil
func sortQueue(items []Node, done chan int) {
    if len(items) < 2 {
        done <- 1
        return
    }
    left, right := 0, len(items) - 1
    pivot := rand.Int() % len(items)
    items[pivot], items[right] = items[right], items[pivot]
    for i, _ := range items {
        if items[i].Schedule[0] < items[right].Schedule[0] {
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

// this function is used to check if the head fo the queue is in an execution state
// parameters: items (a map of ints to arrays of jobs)
// returns: nil
func checkHead(items map[int][]Node) {
    if _, ok := items[0]; ok {
        items, _ := json.Marshal(items[0])
        go MakePostRequest("http://localhost:3939/execute", bytes.NewBuffer(items))
    }
}
// this function is used to group jobs by the number of seconds until execution
// parameters: items (an array of jobs)
// returns: map[int][]Node (a map of the seconds until each job execute to the jobs scheduled to execute then)
func groupItems(items []Node) map[int][]Node {
    output := make(map[int][]Node)
    for _, item := range items {
        for i := 0; i < len(item.Schedule); i++ {
            if len(string(item.Schedule[i])) != 0 {
                output[item.Schedule[i]] = append(output[item.Schedule[i]], item)
            }
        }
    }
    return output
}

// this function is used to convert the cron time string into seconds
// parameters: cronTime (a string of the cron time string format for a job's execution)
// returns: []int (an arry of times until a job executes in seconds)
func cronToSeconds(cronTime string) []int {
    var times []int
    expressions := strings.Split(cronTime, ",")
    for i := 0; i < len(expressions) - 1; i++ {
        times = append(times, int(cronexpr.MustParse(expressions[i]).Next(time.Now()).Sub(time.Now()).Seconds()))
    }
    sort.Ints(times)
    return times
}

// this function is used to populate the queue, calling sorting and grouping methods before checking the head
// parameters: t (the time interval betwen each execution of the fillQueue function)
// returns: nil
func fillQueue(t time.Time) {
    var queue Queue
    var node Node
    jobs := GetAll(SetupClient())
    for _, j := range jobs {
        node.ID, node.UID, node.GID, node.Lang, node.File = j.ID, j.UID, j.GID, j.Language, j.File
        if len(j.Schedule) > 0 {
            node.Schedule = cronToSeconds(j.Schedule)
            queue.Items = append(queue.Items, node)
            channel := make(chan int)
            go sortQueue(queue.Items, channel)
            <-channel
        }
    }
    go checkHead(groupItems(queue.Items))
}

// this function is used to execute the fillQueue function every second
// parameters: d (the duration between execution of fillQueue), f (the function to execute - in this case it's fillQueue)
// returns: nil
func doEvery(d time.Duration, f func(time.Time)) {
    for x := range time.Tick(d) {
        go f(x)
    }
}

// this function starts the countdown process, specifying the paramaters of execution for doEvery
// parameters: nil
// returns: nil
func StartTicker() {
   go doEvery(1000*time.Millisecond, fillQueue)
}
