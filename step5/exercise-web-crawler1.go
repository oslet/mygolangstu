package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func crawl(url string, depth int, fetch Fetcher, out chan string, end chan bool) {
	if depth <= 0 {
		end <- true
		return
	}

	if _, ok := crawled[url]; ok {
		end <- true
		return
	}
	crawledMutex.Lock()
	crawled[url] = true
	crawledMutex.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		out <- fmt.Sprintln(err)
		end <- true
		return
	}

	out <- fmt.Sprintf("found: %s %q\n", url, body)
	subEnd := make(chan bool)
	for _, u := range urls {
		go crawl(u, depth-1, fetcher, out, subEnd)
	}

	for i := 0; i < len(urls); i++ {
		<-subEnd
	}

	end <- true
}

var crawled = make(map[string]bool)
var crawledMutex sync.Mutex

func main() {
	out := make(chan string)
	end := make(chan bool)

	go crawl("http://golang.org/", 4, fetcher, out, end)
	for {
		select {
		case t := <-out:
			fmt.Print(t)
		case <-end:
			return
		}
	}
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
