package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/RinardNick/gophercises/04_htmllinkparser/link"
)

func main() {

	urlFlag := flag.String("url", "http://www.talagentfinancial.com", "the url to create sitemap for")
	flag.Parse()

	pages := get(*urlFlag)
	for _, p := range pages {
		fmt.Printf("%#v\n", p)
	}
}

func get(urlStr string) []string {
	resp, err := http.Get(*&urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Uncomment this line to see whole webpage
	// io.Copy(os.Stdout, resp.Body)

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()

	return filter(base, hrefs(resp.Body, base))
}

func hrefs(body io.Reader, base string) []string {
	links, _ := link.Parse(body)
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "#"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		default:
			continue
		}
	}
	return ret
}

func filter(base string, links []string) []string {
	var ret []string

	for _, l := range links {
		if strings.HasPrefix(l, base) {
			ret = append(ret, l)
		}
	}

	return ret
}

// 1. Get page
// 2. Get all links on first page
// 3. build proper urls so all pages have domain
// 	3a. filter out anly links with different domain
// 4. Store links in first level
// visit each link
//	if link matches any links in above levels -> ignore
//	if link has domain
//	store link under respective top link
//
