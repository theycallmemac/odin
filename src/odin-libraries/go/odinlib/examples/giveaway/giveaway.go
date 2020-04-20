package main

import (
	"fmt"
	"net/http"
        "net/smtp"
        "log"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

        "gitlab.computing.dcu.ie/mcdermj7/2020-ca400-urbanam2-mcdermj7/src/odin-libraries/go/odinlib"
)


func errorCheck(err error)  {
    if err != nil {
        panic(err)
    }
}

func send(bodyArray []string) {
	var body string
	for _, line := range bodyArray {
		body = body + line  + "\n"
	}
	from := "your.email@gmail.com"
	pass := "Gmail app password"
	to := "your.email@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Enter these competitions" + "\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}


func main() {
        val, _ := odinlib.Setup("giveaway.yml")
        if val {
            var links[3] string
            resp, err := http.Get("https://competitions.ie")
            errorCheck(err)
            root, err := html.Parse(resp.Body)
            errorCheck(err)
            articles := scrape.FindAll(root, scrape.ByTag(atom.A))
            count := 0
            for _, article := range articles {
                    if count >= 3 {
                        break
                    }
                    output :=  scrape.Attr(article, "href")
                    if len(output) > 34 && output[0:5] != "https" && output[34:37] == "win" {
                        out := fmt.Sprintf("%s%s", "https:", output)
                        respOut, errOut := http.Get(out)
                        errorCheck(errOut)
                        if respOut.StatusCode == 200 {
                            links[count] = out
                            count += 1
                        }
                }
            }
            var comps []string
            for _, link := range links {
                resp, err := http.Get(link)
                errorCheck(err)
                root, err := html.Parse(resp.Body)
                errorCheck(err)
                linksMatched := func(n *html.Node) bool {
                    if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
                        return scrape.Attr(n.Parent.Parent, "class") == "bluebg"
                    }
                    return false
                }
                compLinks := scrape.FindAll(root, linksMatched)
                compLink := scrape.Attr(compLinks[0], "href")[2:]
                odinlib.Watch("final link", compLink)
                comps = append(comps, compLink)
            }
            send(comps)
        } 
}
