package main

import (
	"fmt"
	"os"
	"strings"

	"net/http"

	"github.com/pkg/browser"
	"golang.org/x/net/html"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "luckygo"
	app.Usage = "Open the top search results of any google search"

	const searchEngine = "https://google.com"
	const urlPrefix = "/url?q="
	searchPrefix := searchEngine + "/search?q="

	var limit int
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "limit",
			Value:       5,
			Usage:       "number of search results to open",
			Destination: &limit,
		},
	}

	app.Action = func(c *cli.Context) error {

		searchBody := ""
		if c.NArg() == 0 {
			fmt.Printf("Must specify search query as first argument")
		}
		searchBody = strings.Replace(c.Args()[0], " ", "+", -1)

		response, _ := http.Get(searchPrefix + searchBody)
		tokenizer := html.NewTokenizer(response.Body)

		results := 0

		for {
			tagtok := tokenizer.Next()

			switch {

			case tagtok == html.ErrorToken:
				// End of html, done
				return nil

			case results == limit:
				return nil

			case tagtok == html.StartTagToken:
				tok := tokenizer.Token()

				isAnchor := tok.Data == "a"

				if isAnchor {
					for _, attr := range tok.Attr {
						if attr.Key == "href" {

							if strings.Contains(attr.Val, urlPrefix) {
								// exclude '/url?q=' from final url
								url := strings.Replace(attr.Val, urlPrefix, "", 1)
								// exclude & and everything after it, unneeded for destination url
								if endIndex := strings.Index(url, "&"); endIndex > 0 {
									url = url[:endIndex]
								}
								// only use this url if it doesn't have 'google' in it
								if !strings.Contains(url, "google") {
									results++
									fmt.Printf("%d) %s\n", results, url)
									browser.OpenURL(url)
								}
							}

							break
						}
					}
				}
			}
		}
	}

	app.Run(os.Args)
}
