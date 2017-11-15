package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	const google = "https://google.com"

	response, _ := http.Get(google)
	rawhtml, _ := ioutil.ReadAll(response.Body)

	fmt.Printf("Raw HTML from %s\n\n\n%s", google, rawhtml)
}
