package main

import (
	"fmt"
	"strings"

	"github.com/pkg/browser"
	"golang.org/x/net/html"
	"net/http"
)

func main() {
	const google = "https://google.com"
	const search = google + "/search?q="
	const tmp = "funny+cat"
	const urlPrefix = "/url?q="
	const maxResults = 5

	response, _ := http.Get(search + tmp)
	tokenizer := html.NewTokenizer(response.Body)

	results := 0

	for {
		tagtok := tokenizer.Next()

		switch {

		case tagtok == html.ErrorToken:
			// End of html, done
			return

		case results == maxResults:
			return

		case tagtok == html.StartTagToken:
			tok := tokenizer.Token()

			isAnchor := tok.Data == "a"

			if isAnchor {
				for _, attr := range tok.Attr {
					if attr.Key == "href" {

						if strings.Contains(attr.Val, urlPrefix) {
							results += 1
							url := strings.Replace(attr.Val, urlPrefix, "", 1)
							fmt.Printf("%d) %s\n", results, url)
							browser.OpenURL(url)
						}

						break
					}
				}
			}
		}
	}
}
