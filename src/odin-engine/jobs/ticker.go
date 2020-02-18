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

func sortQueue(items []Node) []Node {
    if len(items) < 2 {
        return items
    }
    left, right := 0, len(items)-1
    pivot := rand.Int() % len(items)
    items[pivot], items[right] = items[right], items[pivot]
    for i, _ := range items {
        if items[i].Schedule < items[right].Schedule {
            items[left], items[i] = items[i], items[left]
            left++
        }
    }
    items[left], items[right] = items[right], items[left]
    sortQueue(items[:left])
    sortQueue(items[left+1:])
    return items
}

func checkHead(items []Node) {
    if len(items) != 0 && items[0].Schedule <= 1 {
        top := items[0]
        resp := MakePostRequest("http://localhost:3939/execute", bytes.NewBuffer([]byte(top.Lang + " " + top.File + " " + top.ID)))
	fmt.Println("executed job", resp)
    }
}

func fillQueue(t time.Time) {
    var queue Queue
    var node Node
    jobs := GetAll(SetupClient())
    for _, j := range jobs {
        node.ID, node.Lang, node.File = j.ID, j.Language, j.File
        if len(j.Schedule) > 0 {
            node.Schedule = int(cronexpr.MustParse(j.Schedule[:len(j.Schedule)-1]).Next(time.Now()).Sub(time.Now()).Seconds())
            queue.Items = append(queue.Items, node)
        }
    }
    queue.Items = sortQueue(queue.Items)
    checkHead(queue.Items)
}

func doEvery(d time.Duration, f func(time.Time)) {
    for x := range time.Tick(d) {
        f(x)
    }
}

func StartTicker() {
   doEvery(1000*time.Millisecond, fillQueue)
}
