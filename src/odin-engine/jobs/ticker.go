package jobs

import (
    "fmt"
    "github.com/gorhill/cronexpr"
    "time"
    "sort"
)
func doEvery(d time.Duration, f func(time.Time)) {
    for x := range time.Tick(d) {
        f(x)
    }
}

func countdown(t time.Time) {
    // this is a temporary solution
    var queue []int
    jobs := GetAll(SetupClient())
    for _, j := range jobs {
        schedule := j.Schedule[:len(j.Schedule)-1]
        queue = append(queue, int(cronexpr.MustParse(schedule).Next(time.Now()).Sub(time.Now()).Seconds()))
    }
    sort.Ints(queue[:])
    if queue[0] < 1 {
        // here we make a call to the /execute endpoint
	fmt.Println("execute job")
    }
}

func StartTicker() {
   doEvery(1000*time.Millisecond, countdown)
}
