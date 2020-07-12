package main

// reading in files in go https://appdividend.com/2019/12/19/how-to-read-files-in-golang-go-file-read-example/
// handling JSON in go https://blog.golang.org/json

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	story "story"
	"strings"
)

func main() {
	var storyPath = flag.String("file", "storyTemplate.json", "the file path for the story.json")
	var port = flag.Int("port", 3000, "the port to run the application on")

	f, err := os.Open(*storyPath)
	if err != nil {
		panic(err)
	}

	chapters, err := story.Read(f)
	if err != nil {
		panic(err)
	}

	h := NewHandler(chapters)
	fmt.Printf("Starting server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}

func NewHandler(s story.Story) handler {
	return handler{s}
}

type handler struct {
	s story.Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]
	if arc, ok := h.s[path]; ok {
		err := tpl.Execute(w, arc)
		if err != nil {
			panic(err)
		}
	}
}

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

var tpl *template.Template

var defaultHandlerTemplate = `
<!doctype html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport"
              content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>Choose Your Own Adventure</title>
    </head>
        <body>
        <h1>{{.Title}}</h1>
        {{range .Paragraph}}
            <p>{{.}}</p>
        {{end}}
        <ul>
        {{range .Options}}
            <li><a href="/{{.Arc}}">{{.Text}}</a></li>
        {{end}}
        </ul>
    </body>
</html>`

