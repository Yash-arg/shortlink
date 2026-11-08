package main

import (
	"fmt"

	"shortlink/links"
)

func main() {
	store := links.NewStore()
	store.Add("go", "https://go.dev")
	store.Add("py", "https://python.org")

	for _, l := range store.List() {
		fmt.Printf("%s -> %s (hits: %d)\n", l.Code, l.URL, l.Hits)
	}

	for _, code := range []string{"go", "nope"} {
		link, err := store.Find(code)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println("found:", link.URL)
	}
}
