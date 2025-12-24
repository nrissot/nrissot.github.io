package main

import "fmt"

// │ ├ └ ─ ╴

func DrawTree(pages []Page) {
	for _, p := range pages {
		fmt.Println(p.URL)
	}
}
