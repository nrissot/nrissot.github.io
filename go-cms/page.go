package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title       string
	URL         string
	Content     string
	Description string
	HTML        string
	Tags        []string
}

type ListPage struct {
	Page
	Articles []Page
}

func CreatePage(title, url, content, description string, tags []string) Page {
	return Page{title, url, content, description, "", tags}
}

func CreateListPage(title, url, content, description string) ListPage {
	return ListPage{CreatePage(title, url, content, description, nil), []Page{}}
}

func (lp *ListPage) AddPage(p *Page) {
	lp.Articles = append(lp.Articles, *p)
}

func (lp *ListPage) GenerateHTML() {
	articlesHTML := "<nav class=\"articles\">"
	for _, p := range lp.Articles {
		articlesHTML += GenerateArticleCardHTML(p) + "\n"
	}
	articlesHTML += "</nav>"
	lp.HTML = fmt.Sprintf("%s\n<body><pre>%s\n%s\n%s\n%s</pre></body></html>", GenerateHeadHTML(), GenerateHeaderHTML(), ConvertToHTMLFragment(lp.Content), articlesHTML, GenerateFooterHTML())
}

func (p *Page) GenerateHTML() {
	p.HTML = fmt.Sprintf("%s<body><pre>%s</pre>\n%s\n%s</body></html>", GenerateHeadHTML(), GenerateHeaderHTML(), ConvertToHTMLFragment(p.Content), GenerateFooterHTML())
}

func (p Page) WriteToFile() error {
	if p.HTML == "" {
		p.GenerateHTML()
	}
	return os.WriteFile(DST+p.URL, []byte(p.HTML), 0666)
}

func (lp ListPage) WriteToFile() error {
	if lp.HTML == "" {
		lp.GenerateHTML()
	}
	return os.WriteFile(DST+lp.URL, []byte(lp.HTML), 0666)
}
