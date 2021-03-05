package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"github.com/theycallmemac/odin/odin-libraries/go/odinlib"
	"github.com/yhat/scrape"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func send(bodyArray []string) {
	var body string
	for _, line := range bodyArray {
		body = body + line + "\n"
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
	odin, _ := odinlib.Setup("giveaway.yml")
	var links [3]string
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
		output := scrape.Attr(article, "href")
		if len(output) > 34 && output[0:5] != "https" && output[34:37] == "win" {
			out := fmt.Sprintf("%s%s", "https:", output)
			respOut, errOut := http.Get(out)
			errorCheck(errOut)
			if respOut.StatusCode == 200 {
				links[count] = out
				count ++
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
		odin.Watch("final link", compLink)
		comps = append(comps, compLink)
	}
	send(comps)
}
