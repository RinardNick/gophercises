package main

import (
	"fmt"
	"net/http"

	link "github.com/RinardNick/gophercises/04_htmllinkparser/link"
)

func main() {

	response, err := http.Get("https://www.talagentfinancial.com/")
	if err != nil {
		panic(err)
	}

	links, err := link.Parse(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", links)
}
