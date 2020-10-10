package main

import (
	"fmt"

	link "github.com/RinardNick/gophercises/04_htmllinkparser/link"
)

func main() {
	htmlFilepath := "../solutions/04 link/ex1.html"

	link, err := link.Parse(htmlFilepath)
	if err != nil {
		panic(err)
	}

	for _, v := range link {
		fmt.Println(v)
	}

}
