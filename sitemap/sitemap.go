package sitemap

import (
	"encoding/xml"
	"gophercises/link"
	"net/http"
	"net/url"
	"strings"
)

type Sitemap struct {
	XMLName xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	UrlSet  []Url    `xml:"url"`
}

type Url struct {
	Loc string `xml:"loc"`
	// LastMod    string `xml:"lastmod"`
	// ChangeFreq string `xml:"changefreq"`
	// Priority   string `xml:"priority"`
}

func constructLink(domain string, href string) string {
	var s string
	// Handle "/path" and "http://domain.com"
	// If path, then add domain name to it
	if strings.HasPrefix(href, "/") {
		s = domain + href
	}
	// If a valid url, then return it
	if _, err := url.ParseRequestURI(href); err != nil {
		s = href
	}
	return s
}

func Get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	links, _ := link.Parse(resp.Body)

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	return Links(baseUrl.String(), links, WithPrefix(baseUrl.String()))
}

func Links(baseUrl string, links []link.Link, filterFn func(string) bool) []string {
	var ret []string
	for _, l := range links {
		s := constructLink(baseUrl, l.Href)
		// Filter url with specified domain name
		if filterFn(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func WithPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}

func BFS(urlStr string, maxDepth int) Sitemap {
	// Empty struct{} is cheaper than bool
	// Source: https://dave.cheney.net/2014/03/25/the-empty-struct
	// visited
	seen := make(map[string]struct{})
	// queue
	var q map[string]struct{}
	// next queue
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}
	for i := 0; i < maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		for url := range q {
			if _, ok := seen[url]; ok {
				// If the page is already seen, skip it
				continue
			}
			seen[url] = struct{}{}

			for _, link := range Get(url) {
				nq[link] = struct{}{}
			}
		}
	}
	var ret Sitemap
	for url := range seen {
		ret.UrlSet = append(ret.UrlSet, Url{url})
	}
	return ret
}

func XMLEncode(v interface{}) ([]byte, error) {
	xml, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, err
	}
	return xml, nil
}
