package main

import (
	"fmt"
	"regexp"
	"strings"
)

var header string = `<header><h1><a href="/index.html">Nathan Rissot</a></h1>                                    <a href="https://github.com/nrissot">Github</a>  <nav><a href="/blog.html">Blog</a>  <a href="/index.html#contact">Contact</a></nav>
<span aria-hidden="true" class="gray">::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::</span></header>`

var footer string = `<footer>Copyright (c) %d Nathan Rissot · curl this page:
$ curl <a href="https://nrissot.github.io%s.txt">https://nrissot.github.io%s.txt</a>
sitemap: <a href="https://nrissot.github.io/map.txt">https://nrissot.github.io/map.txt</a></footer>`

var head string = `<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Nathan Rissot [%s]</title>
    <link rel="stylesheet" href="/static/style.css">
    <link rel="icon" type="image/x-icon" href="/static/assets/images/favicon.ico">
</head>
<script src="/script.js"></script>
`

var details string = `<span class="gray"><time datetime="%s">%s</time> ·%s</span>`

var article_card string = `<article>
	<h3><a href="%s.html"> -> %s</a></h3>
	%s
	<p>%s</p>
</article>`

func generateTags(tags []string) string {
	var out string = ""
	for _, tag := range tags {
		out += " <a href=\"/tags/" + tag + ".html\">[" + tag + "]</a>"
	}
	return out
}

func (p *Page) GenerateHTML() {
	// assumption: a page is always a blog article

	// HEAD
	var html string = fmt.Sprintf(head, "BLOG")

	html += "<body>\n"

	// HEADER
	html += header

	// MAIN
	html += "\n<main>\n"

	html += "<h1>" + p.Title + "</h1>\n"

	// tags & date
	html += fmt.Sprintf(details, p.Date.Format("2006-01-02"), p.Date.Format("02/01/2006"), generateTags(p.Tags))

	html += _normalize_quotes(_add_hashtag_to_titles(ConvertToHTMLFragment(p.Content)))

	html += "</main>\n"

	// FOOTER
	html += fmt.Sprintf(footer, YEAR, p.URL, p.URL)

	// CLOSURE
	html += "</body></html>"
	p.HTML = html
}

func (lp *ListPage) GenerateHTML() {
	// HEAD
	var html string = fmt.Sprintf(head, strings.ToUpper(lp.Title))
	html += "<body>\n"

	// HEADER
	html += header

	// MAIN
	html += "\n<main>\n"

	html += "<h1>" + lp.Content + "</h1>\n"

	for _, article := range lp.Articles {
		html += fmt.Sprintf(article_card, article.URL, article.Title, fmt.Sprintf(details, article.Date.Format("2006-01-02"), article.Date.Format("02/01/2006"), generateTags(article.Tags)), article.Description)
	}

	// FOOTER
	html += fmt.Sprintf(footer, YEAR, lp.URL, lp.URL)

	// CLOSURE
	html += "</body></html>"
	lp.HTML = html
}

func _add_hashtag_to_titles(s string) string {
	h2reg := regexp.MustCompile(`(^|[^_])(<\bh2\b.*">)([^_]|$)`)
	h3reg := regexp.MustCompile(`(^|[^_])(<\bh3\b.*">)([^_]|$)`)
	h4reg := regexp.MustCompile(`(^|[^_])(<\bh4\b.*">)([^_]|$)`)
	h5reg := regexp.MustCompile(`(^|[^_])(<\bh5\b.*">)([^_]|$)`)
	h6reg := regexp.MustCompile(`(^|[^_])(<\bh6\b.*">)([^_]|$)`)
	return h6reg.ReplaceAllString(h5reg.ReplaceAllString(h4reg.ReplaceAllString(h3reg.ReplaceAllString(h2reg.ReplaceAllString(s, "$1$2## $3"), "$1$2### $3"), "$1$2#### $3"), "$1$2##### $3"), "$1$2###### $3")
}

func _normalize_quotes(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "’", "'"), "”", "\""), "“", "\"")
}
