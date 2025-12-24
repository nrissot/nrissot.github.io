package main

import (
	"fmt"
	"time"
)

const SRC = "../articles"

const DST = ".."

// current year
var YEAR = time.Now().Year()

func main() {
	// TODO : sitemap.txt

	// Find markdown blog articles
	files := FindMarkdownFiles(SRC)

	if len(files) == 0 {
		fmt.Printf("no file found at %s\n", SRC)
		return
	} else {
		fmt.Printf("%d files found at %s\n", len(files), SRC)
	}

	// Create the main blog page
	blog := CreateListPage("BLOG", "/blog", "Tout les articles :", "Tout les articles")

	// create the tags pages
	var tags_pages map[string]*ListPage = map[string]*ListPage{}

	for _, file := range files {
		page := file.AsPage("/blog/")
		for _, t := range page.Tags {
			lp, ok := tags_pages[t]
			if !ok {
				nlp := CreateListPage(t, "/tags/"+t, "Articles avec le tag ["+t+"] :", "Articles avec le tag ["+t+"]")
				tags_pages[t] = &nlp
				lp = &nlp
			}
			lp.AddPage(&page)
		}
		err := page.WriteToFile()
		if err != nil {
			panic(err)
		}
		fmt.Println(page.Title)
		blog.AddPage(&page)
	}

	// write the main blog page to file
	err := blog.WriteToFile()
	if err != nil {
		panic(err)
	}

	// create the main tags page
	tags := CreateListPage("TAGS", "/tags", "Tout les Tags :", "Tout les Tags")

	// write the tags pages to file, and add them to the main tags page.
	for _, lp := range tags_pages {
		nbarticles := len(lp.Articles)
		s := ""
		if nbarticles > 1 {
			s = "s"
		}
		lp.Description += fmt.Sprintf(" - %d article%s.", nbarticles, s)
		tags.AddPage(&lp.Page)
		err = lp.WriteToFile()
		if err != nil {
			panic(err)
		}
	}

	// write the main tags page to file
	err = tags.WriteToFile()
	if err != nil {
		panic(err)
	}

}
