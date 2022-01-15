package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// A map to keep track of visited URLs
type UrlMap struct {
	mu   sync.Mutex
	urls map[string]bool
}

func (u *UrlMap) Insert(key string) {
	u.mu.Lock()
	u.urls[key] = true
	u.mu.Unlock()
}

func (u *UrlMap) Contains(key string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	_, inMap := u.urls[key]
	return inMap
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, urlMap *UrlMap, wg *sync.WaitGroup) {
	defer wg.Done()
	// TODO: Fetch URLs in parallel.
	if depth <= 0 {
		return
	}

	// TODO: Don't fetch the same URL twice.
	visited := urlMap.Contains(url)
	if visited {
		return
	}
	urlMap.Insert(url)

	var currWG sync.WaitGroup
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		currWG.Add(1)
		go Crawl(u, depth-1, fetcher, urlMap, &currWG)
	}
	currWG.Wait()

}

func main() {
	urlMap := UrlMap{urls: make(map[string]bool)}
	var wg sync.WaitGroup
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, &urlMap, &wg)

	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
