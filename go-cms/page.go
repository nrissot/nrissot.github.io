package main

import (
	"os"
	"slices"
	"time"
)

type Page struct {
	Title       string
	URL         string
	Content     string
	Description string
	HTML        string
	TXT         string
	Date        time.Time
	Tags        []string
}

type ListPage struct {
	Page
	Articles []*Page
}

func CreatePage(title, url, content, description string, date time.Time, tags []string) Page {
	return Page{title, url, content, description, "", "", date, tags}
}

func CreateListPage(title, url, content, description string) ListPage {
	return ListPage{CreatePage(title, url, content, description, time.Now(), nil), []*Page{}}
}

func (lp *ListPage) AddPage(p *Page) {
	lp.Articles = append(lp.Articles, p)
}

func (p Page) WriteToFile() error {
	if p.HTML == "" {
		p.GenerateHTML()
		p.GenerateTXT()
	}
	err := os.WriteFile(DST+p.URL+".html", []byte(p.HTML), 0666)
	if err != nil {
		return err
	}
	err = os.WriteFile(DST+p.URL+".txt", []byte(p.TXT), 0666)
	return err
}

func (p Page) WriteToFileTxtOnly() error {
	if p.TXT == "" {
		p.GenerateTXT()
	}
	return os.WriteFile(DST+p.URL+".txt", []byte(p.TXT), 0666)
}

func (lp ListPage) WriteToFile() error {
	if lp.HTML == "" {
		// Sort articles in LWFS order (Last Written, First Shown)
		slices.SortFunc(lp.Articles, func(a, b *Page) int { return b.Date.Compare(a.Date) })
		lp.GenerateHTML()
		lp.GenerateTXT()
	}
	err := os.WriteFile(DST+lp.URL+".html", []byte(lp.HTML), 0666)
	if err != nil {
		return err
	}
	err = os.WriteFile(DST+lp.URL+".txt", []byte(lp.TXT), 0666)
	return err
}
