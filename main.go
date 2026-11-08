package main

import "fmt"

type Link struct {
	Code string
	URL  string
	Hits int
}

var links = map[string]Link{}

func addLink(code, url string) {
	links[code] = Link{Code: code, URL: url}
}

func listLinks() {
	for _, link := range links {
		fmt.Printf("%s -> %s (hits: %d)\n", link.Code, link.URL, link.Hits)
	}
}

func findLink(code string) (Link, error) {
	link, ok := links[code]
	if !ok {
		return Link{}, fmt.Errorf("link %q not found", code)
	}
	return link, nil
}

func main() {
	addLink("go", "https://go.dev")
	addLink("py", "https://python.org")
	listLinks()

	for _, code := range []string{"go", "nope"} {
		link, err := findLink(code)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println("found:", link.URL)
	}
}
