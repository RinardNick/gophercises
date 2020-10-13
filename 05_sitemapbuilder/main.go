package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/RinardNick/gophercises/04_htmllinkparser/link"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlSet struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {

	urlFlag := flag.String("url", "http://www.a-p.com", "the url to create sitemap for")
	maxDepth := flag.Int("depth", 3, "max link search depth")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)
	toXml := urlSet{
		Xmlns: xmlns,
	}
	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}
	fmt.Print(xml.Header)

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
}

// bfs = breadth first search
func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}
	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 {
			break
		}
		for url, _ := range q {
			if _, ok := seen[url]; ok { // if we have seen the page, pass
				continue
			} else { // if we have not seen the page
				seen[url] = struct{}{}          // mark url as seen
				for _, link := range get(url) { // for every link in put in next queue
					if _, ok := seen[link]; !ok {
						nq[link] = struct{}{}
					}
				}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}

	return ret
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

	return filter(hrefs(resp.Body, base), withPrefix(base))
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

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string

	for _, l := range links {
		if keepFn(l) {
			ret = append(ret, l)
		}
	}

	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
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
